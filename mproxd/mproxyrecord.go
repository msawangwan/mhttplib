package mproxd

import (
	"errors"
	"github.com/msawangwan/mhttplib/mgeoloc"
)

var NilRecordError = errors.New("tried to read a nil record")

type ProxyRequestRecord struct {
	RequestedResource string
	mgeoloc.GeographicLocation
}

func NewProxyRequestRecord(ip, resource string) *ProxyRequestRecord {
	return &ProxyRequestRecord{
		RequestedResource:  resource,
		GeographicLocation: mgeoloc.FromAddr(ip),
	}
}
