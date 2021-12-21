package json

import (
	"github.com/NETkiddy/nft-svr/common/cycleImportModels"
)

//=============================================================================
type SignUpRequest struct {
	RequestId string
	Nickname  string
	Username  string
	Gender    int
	Avatar    string
	Auth      AuthData
}
type SignUpResponse struct {
	Sid      string
	SignCode int //0,未注册，1.已注册(未登陆)，2.已登陆
}

type AuthData struct {
	ProviderId   int
	IdentityType int
	Identifier   string
	Credential   string
	State        int
}

//=============================================================================
type SignInRequest struct {
	RequestId string
	Auth      AuthData
}
type SignInResponse struct {
	Sid      string
	SignCode int //0,未注册，1.已注册(未登陆)，2.已登陆
}

//=============================================================================
type SignOutRequest struct {
	RequestId string
}
type SignOutResponse struct {
	Id uint
}

//=============================================================================
type CheckProviderByTokenRequest struct {
	RequestId string
	Sid       string
}
type CheckProviderByTokenResponse struct {
	Provider cycleImportModels.Provider
}

//=============================================================================
type CreateProviderRequest struct {
	RequestId string
	Nickname  string
	Username  string
	Password  string
}
type CreateProviderResponse struct {
	Id uint
}

//=============================================================================
type DeleteProviderRequest struct {
	RequestId     string
	Uin           string
	ProviderUuids []string
}
type DeleteProviderResponse struct {
}

//=============================================================================
type QueryProviderRequest struct {
	RequestId      string
	Uin            string
	QueryStr       string
	ProviderUuids  []string
	ProviderStates []int
	Offset         int
	Limit          int
}
type QueryProviderResponse struct {
	ProviderList []cycleImportModels.Provider `json:"List"`
	TotalCount   int                          `json:"Count"`
}

//=============================================================================
type UpdateProviderRequest struct {
	RequestId  string
	Uin        string
	ArtistUuid string
	Name       string
}
type UpdateProviderResponse struct {
}
