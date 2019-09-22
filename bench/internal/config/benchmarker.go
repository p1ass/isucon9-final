package config

import "time"

const (
	InitializeTimeout = 20 * time.Second
	APITimeout        = 1 * time.Second
	// BenchmarkTimeout  = 180 * time.Second
	BenchmarkTimeout = 30 * time.Second
)

const (
	AttackSearchTrainTimeout    = 20 * time.Second
	AttackListTrainSeatsTimeout = 20 * time.Second
)

var Debug bool
