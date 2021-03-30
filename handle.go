package main

import "time"

type (
	brutaStats struct {
		RequestsMade      int64
		RequestsFailed    int64
		RequestsSuccessed int64
		StartTime         time.Time
	}

	handle struct {
		UserAgents RainbowTable
		Logins     RainbowTable
		Passwords  RainbowTable
		Stats      brutaStats
		IsVerbose  bool
	}
)
