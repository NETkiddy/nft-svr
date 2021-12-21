package service

import (
	"context"
	"fmt"
	"github.com/NETkiddy/common-go/log"
	"github.com/NETkiddy/nft-svr/common/cycleImportModels"
	"github.com/NETkiddy/nft-svr/provider/model"
	"github.com/NETkiddy/nft-svr/provider/protocol/json"
)

type ProviderService struct {
	Ctx   context.Context
	Skey  string
	Sid   string
	Model *model.ProviderModel
	json.QueryProviderRequest
}

func NewProviderService(ctx context.Context, uin string) (ser *ProviderService, err error) {
	sk, ok := ctx.Value("BEE-SESSION_KEY").(string)
	if !ok {
		log.LoggerFromContextWithCaller(ctx).Warnf("ctx.Value(BEE-SESSION_KEY) to string failed, %v", ctx.Value("BEE-SESSION_KEY"))
	}
	sid, ok := ctx.Value("BEE-SID").(string)
	if !ok {
		log.LoggerFromContextWithCaller(ctx).Warnf("ctx.Value(BEE-SID) to string failed, %v", ctx.Value("BEE-SID"))
	}

	mod, err := model.NewProviderModel(ctx, uin)
	if err != nil {
		err := fmt.Errorf("NewProviderService failed, model.NewProviderModel error info: %v", err.Error())
		return nil, err
	}

	ser = &ProviderService{
		Skey:  sk,
		Sid:   sid,
		Ctx:   ctx,
		Model: mod,
	}
	return
}

func (this *ProviderService) QueryProvider(body model.QueryProviderBody) (totalCount int, contacts []cycleImportModels.Provider, err error) {
	//从db获取总数
	totalCount, err = this.Model.GetQueryCount(body)
	if err != nil {
		log.LoggerFromContextWithCaller(this.Ctx).Errorf("Query failed, Model.GetQueryCount err info: %s", err.Error())
		return
	}
	if totalCount == 0 { //==0，直接返回，减少sql查询
		log.LoggerFromContextWithCaller(this.Ctx).Infof("Query Model.GetQueryCount return 0 value")
		return
	}

	//从db分页搜索
	contacts, err = this.Model.QueryWithLimitation(body)
	if err != nil {
		log.LoggerFromContextWithCaller(this.Ctx).Errorf("Query failed, Model.GetQueryCount err info: %s", err.Error())
		return
	}

	return
}
