module wcrwg-iot-ingress

go 1.19

// But the original grpc-gateway v2.
replace github.com/grpc-ecosystem/grpc-gateway/v2 => github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3

replace gopkg.in/DATA-DOG/go-sqlmock.v1 => gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0

require (
	github.com/766b/chi-prometheus v0.0.0-20211217152057-87afa9aa2ca8
	github.com/go-chi/chi v1.5.4
	github.com/go-chi/render v1.0.2
	github.com/golang/protobuf v1.5.2
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/prometheus/client_golang v1.14.0
	github.com/streadway/amqp v1.0.0
	go.thethings.network/lorawan-stack/v3 v3.23.1
)

require (
	github.com/TheThingsIndustries/protoc-gen-go-flags v1.0.6 // indirect
	github.com/TheThingsIndustries/protoc-gen-go-json v1.4.0 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.6.3 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/gotnospirit/makeplural v0.0.0-20180622080156-a5f48d94d976 // indirect
	github.com/gotnospirit/messageformat v0.0.0-20190719172517-c1d0bdacdea2 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.0.0-00010101000000-000000000000 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/exp v0.0.0-20220706164943-b4a6d9510983 // indirect
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220825204002-c680a09ffe64 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
	google.golang.org/grpc v1.46.2 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
