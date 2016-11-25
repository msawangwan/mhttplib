# mhttplib

* * *

## about

* * *

collection of http and web services by me

- mproxd // a single host reverse proxy with logging
- msrvd // a simple http web server
- mtngin // a simple html template engine
- mgeoloc // a small service that queries a db for geographic data based on ip

## usage

* * *

from the example dir, try running (go run):
- main_prx.go <-- reverse proxy
- main_srv.go <-- serves static files

## status

* * *

wip

- need to add options to pass in flags and opts on startup
- need to find a better service than freegeoip
