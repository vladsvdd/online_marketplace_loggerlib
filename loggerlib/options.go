// Package online_marketplace_libs logger/options.go
package online_marketplace_loggerlib

// Option Вариативные параметры через ...Option (Go-идиоматично)
// Это гибкий и расширяемый способ, который часто используется в библиотеках.
type Option func(*Config)

func WithFilePath(path string) Option {
	return func(cfg *Config) {
		cfg.FilePath = path
	}
}

func WithDebugMode(debug bool) Option {
	return func(cfg *Config) {
		cfg.IsDebug = debug
	}
}

func WithFormat(fmt LogFormat) Option {
	return func(cfg *Config) {
		cfg.Format = fmt
	}
}
