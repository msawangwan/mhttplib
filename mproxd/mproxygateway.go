package mproxd

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
)

type proxyGateway struct {
	target        *url.URL
	reverseproxy  *httputil.ReverseProxy
	routepatterns []*regexp.Regexp
	*ProxyLogger
}

func (pg *proxyGateway) ProxyPassHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("x-mlogaccess", "mlogaccess")

	pg.LogRecord(
		*NewProxyRequestRecord(
			r.Header.Get("x-forwarded-for"),
			"/",
		),
	)

	if pg.routepatterns == nil || pg.parseWhitelist(r) {
		pg.reverseproxy.ServeHTTP(w, r)
	}
}

func (pg *proxyGateway) parseWhitelist(r *http.Request) bool {
	for _, regexp := range pg.routepatterns {
		if regexp.MatchString(r.URL.Path) {
			log.Printf("route accepted: %x", r.URL.Path)
			return true
		}
	}
	log.Printf("route rejected: %x", r.URL.Path)
	return false
}

func NewReverseProxyGateway(targetAddr string, whitelist string) *proxyGateway {
	var redirecthost *url.URL
	var rewhitelist *regexp.Regexp
	var err error

	if redirecthost, err = url.Parse(targetAddr); err != nil {
		log.Println("err, parsing proxy pass target addr: %s", err)
	}

	if rewhitelist, err = regexp.Compile(whitelist); err != nil {
		log.Println("err compiling regex: %s", err)
	}

	return &proxyGateway{
		target:       redirecthost,
		reverseproxy: httputil.NewSingleHostReverseProxy(redirecthost),
		routepatterns: []*regexp.Regexp{
			rewhitelist,
		},
		ProxyLogger: NewProxyLogger(),
	}
}
