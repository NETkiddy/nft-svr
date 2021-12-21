package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/NETkiddy/common-go/config"
	"github.com/NETkiddy/common-go/log"
	"github.com/NETkiddy/nft-svr/common"
	"github.com/NETkiddy/nft-svr/common/cycleImportModels"
	"github.com/NETkiddy/nft-svr/token/model"
	"github.com/NETkiddy/nft-svr/token/protocol/jsn"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

type TokenService struct {
	Ctx   context.Context
	Skey  string
	Sid   string
	Model *model.ProductModel
}

func NewTokenService(ctx context.Context) (ser *TokenService, err error) {
	sk, ok := ctx.Value("BEE-SESSION_KEY").(string)
	if !ok {
		log.LoggerFromContextWithCaller(ctx).Warnf("ctx.Value(BEE-SESSION_KEY) to string failed, %v", ctx.Value("BEE-SESSION_KEY"))
	}
	sid, ok := ctx.Value("BEE-SID").(string)
	if !ok {
		log.LoggerFromContextWithCaller(ctx).Warnf("ctx.Value(BEE-SID) to string failed, %v", ctx.Value("BEE-SID"))
	}
	mod, err := model.NewProductModel(ctx)
	if err != nil {
		err := fmt.Errorf("NewTokenService failed, model.NewProductModel error info: %v", err.Error())
		return nil, err
	}

	ser = &TokenService{
		Skey:  sk,
		Sid:   sid,
		Ctx:   ctx,
		Model: mod,
	}
	return
}

func (this *TokenService) CreateTokenClasses(ctx context.Context, classes cycleImportModels.TokenClasses) (resp jsn.CreateTokenClassesResponse, err error) {
	resp = jsn.CreateTokenClassesResponse{}

	appConfig := config.GetViper("app")
	endpoint := "token_classes"
	content, err := json.Marshal(classes)
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("marshal failed, err info: %s", err.Error())
		return
	}
	openapiURL := fmt.Sprintf("%s%s", appConfig.GetString("server.openapi_url"), endpoint)
	gmtDate := time.Now().Format(http.TimeFormat)
	signature := common.GetSignature(appConfig.GetString("server.secret"), "POST", endpoint, string(content), gmtDate, "")
	auth := fmt.Sprintf("%v %v:%v", "NFT", appConfig.GetString("openapi.key"), signature)
	restyReq := resty.New().SetRetryCount(3).SetTimeout(10 * time.Second).SetRetryWaitTime(100 * time.Millisecond).SetRetryMaxWaitTime(1 * time.Second).R()
	value, err := restyReq.
		SetHeader("Content-Type", "application/json").
		SetHeader("Date", gmtDate).
		SetHeader("Authorization", auth).
		SetBody(content).
		Post(openapiURL)
	log.LoggerWrapperWithCaller().Debugf("CreateTokenClasses bodyStr: %v", string(content))
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("CreateTokenClasses resty.Get[%s] failed, err info: %s", openapiURL, err.Error())
		return
	}

	err = json.Unmarshal(value.Body(), &resp)
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("unmarshal failed, err info: %s", err.Error())
		return
	}
	return
}

func (this *TokenService) GetTokenClasses(ctx context.Context, tokenUuid string) (resp jsn.GetTokenClassesResponse, err error) {
	resp = jsn.GetTokenClassesResponse{}

	appConfig := config.GetViper("app")
	endpoint := "token_classes/" + tokenUuid
	content := ""
	openapiURL := fmt.Sprintf("%s%s", appConfig.GetString("server.openapi_url"), endpoint)
	gmtDate := time.Now().Format(http.TimeFormat)
	signature := common.GetSignature(appConfig.GetString("server.secret"), "GET", endpoint, string(content), gmtDate, "")
	auth := fmt.Sprintf("%v %v:%v", "NFT", appConfig.GetString("openapi.key"), signature)
	restyReq := resty.New().SetRetryCount(3).SetTimeout(10 * time.Second).SetRetryWaitTime(100 * time.Millisecond).SetRetryMaxWaitTime(1 * time.Second).R()
	value, err := restyReq.
		SetHeader("Content-Type", "application/json").
		SetHeader("Date", gmtDate).
		SetHeader("Authorization", auth).
		SetBody(content).
		Get(openapiURL)
	log.LoggerWrapperWithCaller().Debugf("GetTokenClasses bodyStr: %v", string(content))
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("GetTokenClasses resty.Get[%s] failed, err info: %s", openapiURL, err.Error())
		return
	}

	err = json.Unmarshal(value.Body(), &resp)
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("unmarshal failed, err info: %s", err.Error())
		return
	}
	return
}
