package httpcaptcha

import (
	"fmt"
	"github.com/dchest/captcha"
	"net/http"
)

type HttpCaptcha struct {
	Config           *Config
	challengeHandler http.Handler
}

func New(cfg *Config) *HttpCaptcha {
	cfg = useDefaults(cfg)

	return &HttpCaptcha{
		Config:           cfg,
		challengeHandler: captcha.Server(cfg.ImageWidth, cfg.ImageHeight),
	}
}

func (c *HttpCaptcha) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, captcha.New())
}

func (c *HttpCaptcha) Challenge(w http.ResponseWriter, r *http.Request) {
	c.challengeHandler.ServeHTTP(w, r)
}

func (c *HttpCaptcha) Reload(w http.ResponseWriter, r *http.Request) {
	if !captcha.Reload(r.URL.Query().Get(c.Config.IdQuery)) {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "reloaded")
}

func (c *HttpCaptcha) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get(c.Config.IdHeader)
		solution := r.Header.Get(c.Config.SolutionHeader)

		if id == "" {
			errMissingHeader(w, c.Config.IdHeader)
			return
		}
		if solution == "" {
			errMissingHeader(w, c.Config.SolutionHeader)
			return
		}

		if !captcha.VerifyString(id, solution) {
			errInvalidCaptcha(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
