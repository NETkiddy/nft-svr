package handler

import (
	"context"
	"github.com/NETkiddy/common-go/log"
	"github.com/NETkiddy/common-go/svr_adapter/glue"
	"github.com/NETkiddy/nft-svr/common"
	"github.com/NETkiddy/nft-svr/common/cycleImportModels"
	"github.com/NETkiddy/nft-svr/provider/model"
	"github.com/NETkiddy/nft-svr/provider/service"
	"strings"

	"github.com/NETkiddy/nft-svr/provider/protocol/json"
)

type ProviderHandler struct {
}

func NewProvider() *ProviderHandler {
	return &ProviderHandler{}
}

/*
func (this *ProviderHandler) SignUp(ctx context.Context, req *json.SignUpRequest) (resp json.SignUpResponse) {
	resp = json.SignUpResponse{}
	//
	as, err := service.NewProviderService(ctx, "")
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("SignUp failed, Service.NewProviderService() err info: %s", err.Error())
		panic(&glue.ExceptionForAPIV3{Code: common.ErrInternalError, Message: err.Error()})
	}
	if as.ProviderId != 0 { //already login
		log.LoggerFromContextWithCaller(ctx).Debugf("provider logged")
		resp.SignCode = common.ACCOUNT_LOGGED
		resp.Sid = as.Sid
		return
	}
	//检查是否已有账号
	body := model.QueryProviderBody{
		Identifier:     req.Auth.Identifier,
		IdentityType:   req.Auth.IdentityType,
		ProviderStates: []int{common.ACCOUNT_STATE_ACTIVE},
		Offset:         0,
		Limit:          1,
	}
	totalCount, providers, err := as.QueryProvider(body)
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("QueryProvider failed, err info: %s", err.Error())
		panic(&glue.ExceptionForAPIV3{Code: common.ErrInternalError, Message: err.Error()})
	}

	if totalCount != 0 {
		log.LoggerFromContextWithCaller(ctx).Debugf("provider exists, got: %v", providers)
		resp.SignCode = common.ACCOUNT_REGGED
		return
	}

	//创建用户
	u, err := uuid.NewRandom()
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("SignUp failed, Service.NewRandom() err info: %s", err.Error())
		panic(&glue.ExceptionForAPIV3{Code: common.ErrInternalError, Message: err.Error()})
	}
	su := model.SignUpBody{
		Uin:      u.String(),
		Nickname: req.Nickname,
		Username: req.Username,
		Gender:   req.Gender,
		Avatar:   req.Avatar,
		Auth:     req.Auth,
	}

	sid, _, err := as.SignUp(su)
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("SignUp failed: %s", err.Error())
		panic(&glue.ExceptionForAPIV3{Code: common.ErrInternalError, Message: err.Error()})
	}

	resp.Sid = sid
	return
}

func (this *ProviderHandler) SignIn(ctx context.Context, req *json.SignInRequest) (resp json.SignInResponse) {
	resp = json.SignInResponse{}

	as, err := service.NewProviderService(ctx, "")
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("SignIn failed, err info: %s", err.Error())
		panic(&glue.ExceptionForAPIV3{Code: common.ErrInternalError, Message: err.Error()})
	}

	if as.ProviderId != 0 { //already login
		log.LoggerFromContextWithCaller(ctx).Debugf("provider logged")
		resp.SignCode = common.ACCOUNT_LOGGED
		return
	}

	si := service.SignInBody{
		Identifier:   req.Auth.Identifier,
		IdentityType: req.Auth.IdentityType,
	}

	sid, err := as.SignIn(si)
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("SignIn failed, err info: %s", err.Error())
		panic(&glue.ExceptionForAPIV3{Code: common.ErrInternalError, Message: err.Error()})
	}

	resp.Sid = sid
	return
}

func (this *ProviderHandler) SignOut(ctx context.Context, req *json.SignOutRequest) (resp json.SignOutResponse) {
	resp = json.SignOutResponse{}

	as, err := service.NewProviderService(ctx, "")
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("SignOut failed, err info: %s", err.Error())
		panic(&glue.ExceptionForAPIV3{Code: common.ErrInternalError, Message: err.Error()})
	}
	if as.Sid == "" || as.ProviderId == 0 {
		log.LoggerFromContextWithCaller(ctx).Warnf("SignOut warning: sid[%v], providerId[%v]", as.Sid, as.ProviderId)
		return
	}

	si := service.SignOutBody{
		Sid:        as.Sid,
		ProviderId: as.ProviderId,
	}

	err = as.SignOut(si)
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("SignOut failed, err info: %s", err.Error())
		panic(&glue.ExceptionForAPIV3{Code: common.ErrInternalError, Message: err.Error()})
	}

	return
}
*/

func (this *ProviderHandler) CreateProvider(ctx context.Context, req *json.CreateProviderRequest) (resp json.CreateProviderResponse) {

	return
}

func (this *ProviderHandler) DeleteProvider(ctx context.Context, req *json.DeleteProviderRequest) {

}

func (this *ProviderHandler) QueryProvider(ctx context.Context, req *json.QueryProviderRequest) (resp json.QueryProviderResponse) {
	log.LoggerFromContextWithCaller(ctx).Debugf("QueryProvider req: %v", req)

	resp = json.QueryProviderResponse{
		ProviderList: make([]cycleImportModels.Provider, 0),
		TotalCount:   0,
	}

	// 参数处理
	offset := req.Offset
	if offset <= 0 {
		offset = common.DEFAULT_OFFSET
	}
	limit := req.Limit
	if limit <= 0 {
		limit = common.DEFAULT_LIMIT
	}
	queryStr := req.QueryStr
	if queryStr != "" {
		queryStr = strings.Replace(queryStr, "_", "\\_", -1)
		queryStr = strings.Replace(queryStr, "%", "\\%", -1)
	}
	if req.ProviderUuids == nil {
		req.ProviderUuids = make([]string, 0)
	}

	//
	cs, err := service.NewProviderService(ctx, req.Uin)
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("QueryProvider failed, err info: %s", err.Error())
		panic(&glue.ExceptionForAPIV3{Code: common.ErrInternalError, Message: err.Error()})
	}

	body := model.QueryProviderBody{
		QueryStr:       queryStr,
		ProviderStates: req.ProviderStates,
		Offset:         offset,
		Limit:          limit,
	}
	totalCount, contacts, err := cs.QueryProvider(body)
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("QueryProvider failed, err info: %s", err.Error())
		panic(&glue.ExceptionForAPIV3{Code: common.ErrInternalError, Message: err.Error()})
	}

	resp.TotalCount = totalCount
	resp.ProviderList = contacts
	return

}

func (this *ProviderHandler) UpdateProvider(ctx context.Context, req *json.UpdateProviderRequest) {

}
