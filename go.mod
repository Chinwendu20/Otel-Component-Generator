module github.com/Chinwendu20/otel_components_generator

go 1.19

replace (
	github.com/Chinwendu20/otel_components_generator/exporters => ./exporters
	github.com/Chinwendu20/otel_components_generator/internal => ./internal
)

require (
	github.com/Chinwendu20/otel_components_generator/exporters v0.0.0-00010101000000-000000000000
	github.com/Chinwendu20/otel_components_generator/internal v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.6.1
	go.uber.org/zap v1.24.0
)

require (
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)
