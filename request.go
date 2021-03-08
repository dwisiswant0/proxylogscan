package main

import (
	"net/http"
	"net/url"
	"path"

	"ktbs.dev/mubeng/pkg/mubeng"
)

func request(URL string, method string, proxyURL string) (*http.Request, error) {
	u, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "/owa/auth/x.js")
	URL = u.String()

	req, err := http.NewRequest(method, URL, nil)
	if err != nil {
		return req, err
	}
	req.Header.Set("Cookie", "X-AnonResource=true; X-AnonResource-Backend=localhost/ecp/default.flt?~3; X-BEResource=localhost/owa/auth/logon.aspx?~3;")

	if proxyURL != "" {
		client.Transport, err = mubeng.Transport(proxyURL)
		if err != nil {
			return req, err
		}
	}

	return req, nil
}
