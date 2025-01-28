package config

import "time"

const (
	HttpRetryWaitMin = 1 * time.Second
	HttpRetryWaitMax = 30 * time.Second
	HttpRetryMax     = 3
)
