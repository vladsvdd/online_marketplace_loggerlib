// Package online_marketplace_libs logger/builder.go
package online_marketplace_loggerlib

// Builder pattern
type Builder struct {
	cfg Config
}

func NewLoggerBuilder() *Builder {
	return &Builder{
		cfg: Config{
			FilePath: DefaultLogFilePath,
			IsDebug:  false,
			Format:   FormatJSON,
		},
	}
}

func (b *Builder) WithFilePath(path string) *Builder {
	b.cfg.FilePath = path
	return b
}

func (b *Builder) WithDebugMode(debug bool) *Builder {
	b.cfg.IsDebug = debug
	return b
}

func (b *Builder) WithFormat(fmt LogFormat) *Builder {
	b.cfg.Format = fmt
	return b
}

func (b *Builder) Build() (*Logger, error) {
	return buildLogger(b.cfg)
}
