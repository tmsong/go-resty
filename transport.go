// +build go1.13

// Copyright (c) 2015-2020 Jeevanandam M (jeeva@myjeeva.com), All rights reserved.
// resty source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package resty

import (
	"net"
	"net/http"
	"runtime"
)

func createTransport(localAddr net.Addr) *http.Transport {
	dialer := &net.Dialer{
		Timeout:   DefaultTransportTimeout,
		KeepAlive: DefaultTransportKeepAlive,
	}
	if localAddr != nil {
		dialer.LocalAddr = localAddr
	}
	return &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          DefaultTransportMaxIdleConns,
		IdleConnTimeout:       DefaultTransportIdleConnTimeout,
		TLSHandshakeTimeout:   DefaultTransportTLSHandshakeTimeout,
		ExpectContinueTimeout: DefaultTransportExpectContinueTimeout,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
	}
}
