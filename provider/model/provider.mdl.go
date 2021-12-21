package model

import (
	"context"
	"github.com/NETkiddy/nft-svr/common"
	"github.com/NETkiddy/nft-svr/common/cycleImportModels"
	"github.com/NETkiddy/nft-svr/provider/protocol/json"
	"strings"
)

type ProviderModel struct {
	Ctx context.Context
	Uin string
}

func NewProviderModel(ctx context.Context, uin string) (*ProviderModel, error) {
	return &ProviderModel{
		Ctx: ctx,
		Uin: uin,
	}, nil
}

type SignUpBody struct {
	Uin      string
	Username string
	Nickname string
	Gender   int
	Avatar   string
	Auth     json.AuthData
}

func (this *ProviderModel) SignUp(body SignUpBody) (providerData cycleImportModels.Provider, err error) {
	db := common.GetMysqlDb()
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 先创建auth item
	item := cycleImportModels.Auth{
		ProviderId:   body.Auth.ProviderId,
		IdentityType: body.Auth.IdentityType,
		Identifier:   body.Auth.Identifier,
		Credential:   body.Auth.Credential,
		State:        common.AUTH_STATE_ACTIVE,
	}
	providerData.Uin = body.Uin
	providerData.Username = body.Username
	providerData.Nickname = body.Nickname
	providerData.Avatar = body.Avatar
	providerData.Gender = body.Gender
	providerData.State = common.ACCOUNT_STATE_ACTIVE
	providerData.Auths = []cycleImportModels.Auth{item}

	if err = tx.Create(&providerData).Error; err != nil {
		return
	}

	tx.Commit()

	return
}

type QueryProviderBody struct {
	Uin            string
	QueryStr       string
	ProviderStates []int
	Identifier     string
	IdentityType   int
	Offset         int
	Limit          int
}

func (this *ProviderModel) GetQueryCount(body QueryProviderBody) (totalCount int, err error) {
	db := common.GetMysqlDb()
	query := db.Table("provider")

	if this.Uin != "" {
		query = query.Where("uin = ?", this.Uin)
	}
	if len(body.ProviderStates) != 0 {
		query = query.Where("provider.state IN (?)", body.ProviderStates)
	}
	if body.QueryStr != "" {
		body.QueryStr = "%" + body.QueryStr + "%"
		query = query.Where(
			"name LIKE ? OR category LIKE ?",
			body.QueryStr, body.QueryStr)
	}

	if err = query.
		Count(&totalCount).Error; err != nil {
		return
	}

	return
}

func (this *ProviderModel) QueryWithLimitation(body QueryProviderBody) (providers []cycleImportModels.Provider, err error) {
	providers = make([]cycleImportModels.Provider, 0)
	if body.Limit < 0 {
		body.Limit = common.DEFAULT_LIMIT
	}

	db := common.GetMysqlDb()
	query := db.Table("provider")

	if this.Uin != "" {
		query = query.Where("uin = ?", this.Uin)
	}
	if len(body.ProviderStates) != 0 {
		query = query.Where("provider.state IN (?)", body.ProviderStates)
	}
	if body.QueryStr != "" {
		queryStr := strings.Replace(body.QueryStr, "_", "\\_", -1)
		queryStr = strings.Replace(body.QueryStr, "%", "\\%", -1)
		queryStr = "%" + queryStr + "%"
		query = query.Where(
			"name LIKE ? OR category LIKE ?",
			queryStr, queryStr)
	}

	if err = query.
		Order("id DESC", true).
		Offset(body.Offset).
		Limit(body.Limit).
		Find(&providers).Error; err != nil {
		return
	}

	return
}
