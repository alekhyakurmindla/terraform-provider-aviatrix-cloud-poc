package utils

import "time"

const (
	HttpRetryWaitMin = 1 * time.Second
	HttpRetryWaitMax = 30 * time.Second
	HttpRetryMax     = 3
	AviatrixHost     = "AVIATRIX_HOST"
	AviatrixUsername = "AVIATRIX_USERNAME"
	AviatrixPassword = "AVIATRIX_PASSWORD"
)
