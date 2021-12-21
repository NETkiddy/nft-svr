package main

import (
	"github.com/NETkiddy/common-go/svr_adapter/glue"
	"github.com/NETkiddy/nft-svr/common"
)

func InitDatabase() {
	params := []*common.ConnectParam{
		{SourceNameKey: "mysqldata.source_name", Driver: "mysql"},
		{SourceNameKey: "mysqldata.ro_source_name", Driver: "mysql"},
	}
	common.SetConnectParams(params)
	common.InitDb()
}

func CloseDatabase() {
	common.CloseDb()
}

func main() {

	serverOptions := make([]glue.ServerOption, 0)
	serverOptions = append(serverOptions, glue.UseGinFrameWork())
	serverOptions = append(serverOptions, glue.WithExceptionHandler(common.ExceptionHandler))
	serverOptions = append(serverOptions, glue.UseHttpHandler(InitRoute))
	serverOptions = append(serverOptions, glue.RegisterInitEnv(func() { InitDatabase() }))
	serverOptions = append(serverOptions, glue.RegisterDestroy(func() { CloseDatabase() }))

	glue.Run("bee_provider_service", "bee contact service", serverOptions...)
}
