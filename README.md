# Sologenic

This is a home assignment task I completed during one of recruitment processes.
Task description is available in [./doc/description.pdf](./doc/description.pdf).

## Task A.1

* proto spec is in [./exchange/exchange.proto](./exchange/exchange.proto)
* `go:generate` command is in [./generate.go](./generate.go)
* call `go generate` from command line to generate code
* generated code will be in [./exchange/generated/exchange.pb.go](./exchange/generated/exchange.pb.go)
* generated code is git-ignored
* I haven't added endpoint for placing single order because this is just a special case of sending a batch of orders 

## Task A.2

* source code is in [./cache](./cache)
* logic is in [./cache/cmd/main.go](./cache/cmd/main.go)
* two fake providers are implemented in [./cache/supply/providers/grouping/provider.go](./cache/supply/providers/grouping/provider.go) and [./cache/supply/providers/individual/provider.go](./cache/supply/providers/individual/provider.go)
* to run it call `go run ./cache/cmd`

## Task B
* please read [./freeze/doc.md](./freeze/doc.md)