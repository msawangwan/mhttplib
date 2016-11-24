package mproxd

import (
	"fmt"
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

func (pg *proxyGateway) ListenAndRedirect(port string) {
	pg.LogStatus("initialising proxy server ...")
	pg.LogStatus(fmt.Sprintf("listening on %s and redirecting to %v", port, pg.target))

	http.HandleFunc("/", pg.proxyPassHandler)
	http.ListenAndServe(port, nil)
}

func (pg *proxyGateway) proxyPassHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("x-mlogaccess", "mlogaccess")

	pg.LogRecord(
		*NewProxyRequestRecord(
			r.Header.Get("x-forwarded-for"), // request origin
			r.URL.Path,                      // requested resource
		),
	)

	if pg.routepatterns == nil || pg.parseWhitelist(r) {
		pg.LogStatus(fmt.Sprintf("route accepted: %s", r.URL.Path))
		pg.reverseproxy.ServeHTTP(w, r)
	} else {
		pg.LogStatus(fmt.Sprintf("route rejected: %s", r.URL.Path))
	}
}

func (pg *proxyGateway) parseWhitelist(r *http.Request) bool {
	for _, regexp := range pg.routepatterns {
		if regexp.MatchString(r.URL.Path) {
			return true
		}
	}
	return false
}

func NewReverseProxyGateway(targetAddr string, whitelist string) *proxyGateway {
	var redirecthost *url.URL
	var rewhitelist *regexp.Regexp
	var err error

	pl := NewProxyLogger()

	if redirecthost, err = url.Parse(targetAddr); err != nil {
		pl.LogStatus(fmt.Sprintf("err parsing proxy pass target: %v", err))
	}

	if rewhitelist, err = regexp.Compile(whitelist); err != nil {
		pl.LogStatus(fmt.Sprintf("err compiling regex: %v", err))
	}

	return &proxyGateway{
		target:       redirecthost,
		reverseproxy: httputil.NewSingleHostReverseProxy(redirecthost),
		routepatterns: []*regexp.Regexp{
			rewhitelist,
		},
		ProxyLogger: pl,
	}
}
