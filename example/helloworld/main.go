package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/orbitool/httpcaptcha"
	"log"
	"net/http"
)

func main() {
	captcha := httpcaptcha.New(nil)

	router := httprouter.New()
	router.HandlerFunc("GET", "/captcha/create", captcha.Create)
	router.HandlerFunc("GET", "/captcha/challenge/:media", captcha.Challenge)

	// Protected route that requires a captcha solution to view
	// - Requires the header `X-Captcha` to be set to the captcha id generated
	//   by the `captcha/create` method.
	// - Requires the header `X-Captcha-Solution` to be set the captcha solution.
	router.Handler("GET", "/hello-world", captcha.Middleware(helloWorld()))

	log.Panic(http.ListenAndServe(":8080", router))
}

func helloWorld() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})
}
