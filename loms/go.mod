module route256/loms

go 1.22.3

require (
	github.com/envoyproxy/protoc-gen-validate v1.0.4
	github.com/gojuno/minimock/v3 v3.3.11
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.1
	github.com/stretchr/testify v1.9.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240610135401-a8a62080eff3
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.34.2
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240610135401-a8a62080eff3 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// replace google.golang.org/genproto => google.golang.org/genproto v0.0.0-20220407144326-9054f6ed7bac

replace google.golang.org/genproto => google.golang.org/genproto v0.0.0-20240610135401-a8a62080eff3

replace github.com/gojuno/minimock/v3 => github.com/gojuno/minimock/v3 v3.3.11
