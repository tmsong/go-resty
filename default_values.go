/**
 * @note
 * default.go
 *
 * @author	songtianming
 * @date 	2020-09-10
 */
package resty

import "time"

//Retry
var (
	DefaultMaxRetries  = 3
	DefaultWaitTime    = time.Duration(100) * time.Millisecond
	DefaultMaxWaitTime = time.Duration(2000) * time.Millisecond
)

func SetDefaultMaxRetries(retry int) {
	DefaultMaxRetries = retry
}

func SetDefaultRetryWaitTime(waitTime time.Duration) {
	DefaultWaitTime = waitTime
}

func SetDefaultMaxRetryWaitTime(maxWaitTime time.Duration) {
	DefaultMaxWaitTime = maxWaitTime
}

//Transport

var (
	DefaultTransportTimeout               = 30 * time.Second
	DefaultTransportKeepAlive             = 30 * time.Second
	DefaultTransportMaxIdleConns          = 100
	DefaultTransportIdleConnTimeout       = 90 * time.Second
	DefaultTransportTLSHandshakeTimeout   = 10 * time.Second
	DefaultTransportExpectContinueTimeout = 1 * time.Second
)
