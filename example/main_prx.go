package main

import (
	"github.com/msawangwan/mhttplib/mproxd"
)

const (
	LISTEN_PORT       = ":8080"
	PROXY_PASS        = "http://127.0.0.1:9090"
	WHITELIST         = `^\/$|[\w|/]*.js|/tyranny`
	LISTEN_PORT_USAGE = "listen for requests on this port ... i.e, 80, 8080"
	PROXY_PASS_USAGE  = "redirect target address ... i.e, http://127.0.0.1:8080"
	WHITELIST_USAGE   = "match on whitelisted characters ... i.e, regex"
)

func main() {
	proxyGateway := mproxd.NewReverseProxyGateway(PROXY_PASS, WHITELIST)
	proxyGateway.ListenAndRedirect(LISTEN_PORT)
}
