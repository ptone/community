package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func getProxy(service string) (proxy *httputil.ReverseProxy, err error) {
	serviceInfo, _ := url.Parse(service)
	host := serviceInfo.Host
	ts := TokenSource(service)
	proxy = &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			fmt.Println("director")
			req.URL.Scheme = "https"
			req.Host = host
			req.URL.Host = host
			// Inject the proxy-target auth token
			t, _ := ts.Token()
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", t.AccessToken))
			fmt.Printf("%v+\n", req.URL)
			fmt.Printf("%v+\n", req.Header.Get("Authorization"))
		},
	}
	return proxy, nil

}
