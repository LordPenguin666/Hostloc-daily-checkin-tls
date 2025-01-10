package request

import (
	"github.com/imroc/req/v3"
	"net/http"
)

type Request struct {
	client  *req.Client
	cookies []*http.Cookie
}

func New() *Request {
	return &Request{
		client: req.C().
			SetUserAgent("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36").
			SetTLSFingerprintChrome(),
	}
}

func (r *Request) UpdateCookies(cookies []*http.Cookie) {
	r.cookies = cookies
}
