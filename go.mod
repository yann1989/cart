module github.com/yann1989/cart

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.0
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/yann1989/common v0.0.0-20210921130204-4162c2594487
	google.golang.org/protobuf v1.27.1
)
