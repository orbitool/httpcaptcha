package httpcaptcha

import "github.com/dchest/captcha"

type Config struct {
	IdHeader       string
	SolutionHeader string
	IdQuery        string
	ImageHeight    int
	ImageWidth     int
}

func useDefaults(cfg *Config) *Config {
	if cfg == nil {
		cfg = &Config{}
	}

	if cfg.IdHeader == "" {
		cfg.IdHeader = "X-Captcha"
	}

	if cfg.SolutionHeader == "" {
		cfg.SolutionHeader = "X-Captcha-Solution"
	}

	if cfg.IdQuery == "" {
		cfg.IdQuery = "captcha-id"
	}

	if cfg.ImageHeight == 0 {
		cfg.ImageHeight = captcha.StdHeight
	}

	if cfg.ImageWidth == 0 {
		cfg.ImageWidth = captcha.StdWidth
	}

	return cfg
}
