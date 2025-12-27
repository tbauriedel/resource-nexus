package config

import (
	"time"

	"golang.org/x/time/rate"
)

// Listener represents the listener configuration.
type Listener struct {
	ListenAddr                 string        `json:"listenAddr"`
	ReadTimeout                time.Duration `json:"readTimeout"`
	IdleTimeout                time.Duration `json:"idleTimeout"`
	TLSEnabled                 bool          `json:"tlsEnabled"`
	TLSSkipVerify              bool          `json:"tlsSkipVerify"`
	TLSCertPath                string        `json:"tlsCertFile"`
	TLSKeyPath                 string        `json:"tlsKeyFile"`
	GlobalRateLimitGeneration  rate.Limit    `json:"globalRateLimitGeneration"`
	GlobalRateLimitBucketSize  int           `json:"globalRateLimitBucketSize"`
	IpBasedRateLimitGeneration rate.Limit    `json:"ipBasedRateLimitGeneration"`
	IpBasedRateLimitBucketSize int           `json:"ipBasedRateLimitBucketSize"`
}

// Logger represents the logging configuration.
type Logger struct {
	Type  string `json:"type"`
	File  string `json:"file"`
	Level string `json:"level"`
}

// Database represents the database configuration.
type Database struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
	TLSMode  string `json:"tlsmode"`
}
