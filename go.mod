module github.com/owncloud/ocis-graph

go 1.13

require (
	contrib.go.opencensus.io/exporter/jaeger v0.2.0
	contrib.go.opencensus.io/exporter/ocagent v0.6.0
	contrib.go.opencensus.io/exporter/zipkin v0.1.1
	github.com/ascarter/requestid v0.0.0-20170313220838-5b76ab3d4aee
	github.com/cespare/reflex v0.2.0 // indirect
	github.com/go-chi/chi v4.0.2+incompatible
	github.com/go-chi/render v1.0.1
	github.com/golang/protobuf v1.3.2
	github.com/grpc-ecosystem/grpc-gateway v1.12.1
	github.com/justinas/alice v1.2.0
	github.com/kballard/go-shellquote v0.0.0-20180428030007-95032a82bc51 // indirect
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.17.1
	github.com/micro/go-plugins v1.5.1
	github.com/ogier/pflag v0.0.1 // indirect
	github.com/oklog/run v1.0.0
	github.com/openzipkin/zipkin-go v0.2.2
	github.com/owncloud/ocis-hello v0.0.0-20191205174151-7e9b6068aec7
	github.com/owncloud/ocis-pkg v1.0.0
	github.com/prometheus/client_golang v1.2.1
	github.com/spf13/viper v1.5.0
	github.com/tomasen/realip v0.0.0-20180522021738-f0c99a92ddce
	github.com/yaegashi/msgraph.go v0.0.0-20191104022859-3f9096c750b2
	go.opencensus.io v0.22.2
	golang.org/x/net v0.0.0-20191204025024-5ee1b9f4859a
	google.golang.org/genproto v0.0.0-20191203220235-3fa9dbf08042
	gopkg.in/ldap.v3 v3.1.0
)
