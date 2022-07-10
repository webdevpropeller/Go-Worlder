package config

import "strings"

const domain = "http://118.27.23.183:8090"

type URL struct {
	Activate      string
	PasswordReset string
}

func NewURL() *URL {
	return &URL{
		Activate:      strings.Join([]string{domain, "activate"}, "/"),
		PasswordReset: strings.Join([]string{domain, "password", "reset"}, "/"),
	}
}
