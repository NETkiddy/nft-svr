module github.com/NETkiddy/nft-svr

go 1.16

require (
	github.com/NETkiddy/common-go v0.0.0-20211217080833-2f1b79e72fc5
	github.com/gin-gonic/gin v1.7.7
	github.com/go-resty/resty/v2 v2.7.0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/jinzhu/gorm v1.9.16
	github.com/microcosm-cc/bluemonday v1.0.16
	github.com/stretchr/testify v1.7.0
	golang.org/x/net v0.0.0-20211215060638-4ddde0e984e9
)

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
