module github.com/outofforest/sologenic

go 1.16

replace github.com/ridge/parallel => github.com/outofforest/parallel v0.1.2

require (
	github.com/outofforest/build v1.4.0
	github.com/outofforest/buildgo v0.2.1
	github.com/outofforest/ioc/v2 v2.5.0
	github.com/outofforest/libexec v0.2.1
	github.com/outofforest/logger v0.2.0
	github.com/outofforest/run v0.2.2
	github.com/ridge/must v0.6.0
	github.com/ridge/parallel v0.1.1
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.21.0
	google.golang.org/protobuf v1.27.1
)
