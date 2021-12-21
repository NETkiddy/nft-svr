package handler

import (
	"context"
	"github.com/NETkiddy/common-go/log"
	"github.com/NETkiddy/common-go/svr_adapter/glue"
	"github.com/NETkiddy/nft-svr/common"
	"github.com/NETkiddy/nft-svr/common/cycleImportModels"
	"github.com/NETkiddy/nft-svr/token/protocol/jsn"
	"github.com/NETkiddy/nft-svr/token/service"
)

type TokenHandler struct {
}

func NewToken() *TokenHandler {
	return &TokenHandler{}
}

func (h *TokenHandler) CreateTokenClasses(ctx context.Context, req *jsn.CreateTokenClassesRequest) (resp jsn.CreateTokenClassesResponse) {
	resp = jsn.CreateTokenClassesResponse{}

	//参数处理
	if len(req.Name) == 0 || len(req.Description) == 0 {
		log.LoggerFromContextWithCaller(ctx).Errorf("empty Name")
		panic(&glue.ExceptionForAPIV3{Code: common.ErrInvalidParameterValue, Message: "empty Name"})
	}
	if len(req.Description) == 0 {
		log.LoggerFromContextWithCaller(ctx).Errorf("empty Description")
		panic(&glue.ExceptionForAPIV3{Code: common.ErrInvalidParameterValue, Message: "empty Description"})
	}

	classes := cycleImportModels.TokenClasses{
		Name:          req.Name,
		Description:   req.Description,
		Total:         req.CoverImageUrl,
		Renderer:      req.CoverImageUrl,
		CoverImageUrl: req.CoverImageUrl,
	}

	//
	ps, err := service.NewTokenService(ctx)
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("CreateTokenClasses failed, Service.NewArtistService() err info: %s", err.Error())
		panic(&glue.ExceptionForAPIV3{Code: common.ErrInternalError, Message: err.Error()})
	}
	resp, err = ps.CreateTokenClasses(ctx, classes)
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("CreateTokenClasses failed, Service.Create() err info: %s", err.Error())
		panic(&glue.ExceptionForAPIV3{Code: common.ErrInternalError, Message: err.Error()})
	}

	return
}

func (h *TokenHandler) GetTokenClasses(ctx context.Context, req *jsn.GetTokenClassesRequest) (resp jsn.GetTokenClassesResponse) {
	resp = jsn.GetTokenClassesResponse{}

	//参数处理
	if len(req.TokenUuid) == 0 {
		log.LoggerFromContextWithCaller(ctx).Errorf("empty TokenUuid")
		panic(&glue.ExceptionForAPIV3{Code: common.ErrInvalidParameterValue, Message: "empty Name"})
	}

	//
	ps, err := service.NewTokenService(ctx)
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("GetTokenClasses failed, Service.NewArtistService() err info: %s", err.Error())
		panic(&glue.ExceptionForAPIV3{Code: common.ErrInternalError, Message: err.Error()})
	}
	resp, err = ps.GetTokenClasses(ctx, req.TokenUuid)
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("GetTokenClasses failed, Service.Create() err info: %s", err.Error())
		panic(&glue.ExceptionForAPIV3{Code: common.ErrInternalError, Message: err.Error()})
	}

	return
}
