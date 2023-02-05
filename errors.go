package httpcaptcha

import (
	"fmt"
	"net/http"
)

func errMissingHeader(w http.ResponseWriter, header string) {
	http.Error(w, fmt.Sprintf("missing required header: '%s'", header), http.StatusBadRequest)
}

func errInvalidCaptcha(w http.ResponseWriter) {
	http.Error(w, "invalid httpcaptcha solution", http.StatusBadRequest)
}
