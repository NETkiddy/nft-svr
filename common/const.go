package common

const (
	PRODUCT_INVALID = 0 //商品未生效
	PRODUCT_ACTIVE  = 1 //商品有效
	PRODUCT_UPDATE  = 2 //商品更新中
	PRODUCT_DELETED = 3 //商品无效

	BRAND_INVALID = 0 //品牌未生效
	BRAND_ACTIVE  = 1 //品牌有效
	BRAND_UPDATE  = 2 //品牌更新中
	BRAND_DELETED = 3 //品牌无效

	DEFAULT_OFFSET = 0
	DEFAULT_LIMIT  = 10
	SEP_KEYWORD    = "|"

	ACCOUNT_TYPE_PHONE  = 1 //电话登陆
	ACCOUNT_TYPE_EMAIL  = 2 //email登陆
	ACCOUNT_TYPE_WECHAT = 3 //微信账号登陆

	ACCOUNT_UNREG  = 0
	ACCOUNT_REGGED = 1
	ACCOUNT_LOGGED = 2

	ACCOUNT_STATE_INVAILD = 0
	ACCOUNT_STATE_ACTIVE  = 1
	ACCOUNT_STATE_UPDATE  = 2
	ACCOUNT_STATE_DELETED = 3

	AUTH_STATE_INVALID = 0 //auth未生效
	AUTH_STATE_ACTIVE  = 1 //auth生效

	INVOICE_INVALID = 0 //商品未生效
	INVOICE_ACTIVE  = 1 //商品有效
	INVOICE_UPDATE  = 2 //商品更新中
	INVOICE_DELETED = 3 //商品无效

	ARTIST_INVALID = 0 //商品未生效
	ARTIST_ACTIVE  = 1 //商品有效
	ARTIST_UPDATE  = 2 //商品更新中
	ARTIST_DELETED = 3 //商品无效

	ARTWORK_NOT_ACTIVE = 0 //商品未生效
	ARTWORK_ACTIVE     = 1 //商品有效
	ARTWORK_UPDATE     = 2 //商品更新中
	ARTWORK_DELETED    = 3 //商品无效

	ARTWORKGROUP_NOT_ACTIVE = 0 //商品未生效
	ARTWORKGROUP_ACTIVE     = 1 //商品有效
	ARTWORKGROUP_UPDATE     = 2 //商品更新中
	ARTWORKGROUP_DELETED    = 3 //商品无效

	CONTACT_NOT_ACTIVE = 0 //商品未生效
	CONTACT_ACTIVE     = 1 //商品有效
	CONTACT_UPDATE     = 2 //商品更新中
	CONTACT_INVALID    = 3 //商品无效

	CONTACTGROUP_NOT_ACTIVE = 0 //商品未生效
	CONTACTGROUP_ACTIVE     = 1 //商品有效
	CONTACTGROUP_UPDATE     = 2 //商品更新中
	CONTACTGROUP_INVALID    = 3 //商品无效

	EXHIBITION_NOT_ACTIVE = 0 //商品未生效
	EXHIBITION_ACTIVE     = 1 //商品有效
	EXHIBITION_UPDATE     = 2 //商品更新中
	EXHIBITION_DELETED    = 3 //商品无效

	//dev 'http://gz.mcproxy.tencentyun.com/interface.php'
	//pro  http://mcproxy.tencentyun.com/interface.php

	McProxyUrl = "http://mcproxy.tencentyun.com/interface.php"
	//dev http://account.tencentyun.com
	//pro http://account.tencentyun.com:50001
	CgwAccountPath   = "http://account.tencentyun.com:50001"
	RpcSetName       = "/set/zhuji"
	GOODS_NOT_ACTIVE = 0 //商品未生效
	GOODS_ACTIVE     = 1 //商品有效
	GOODS_UPDATE     = 2 //商品更新中
	GOODS_INVALID    = 3 //商品无效

	GOODS_NOT_ACTIVE_DESC = "新导入，待生效"
	GOODS_ACTIVE_DESC     = "已生效"
	GOODS_UPDATE_DESC     = "有更新,待生效"
	GOODS_INVALID_DESC    = "失效"

	AdminFakeLoginTokenSalt = "zhuji@sh||team***"

	//APIV3 common error code
	ErrInvalidParameter                   = "InvalidParameter"
	ErrInvalidParameterValue              = "InvalidParameterValue"
	ErrMissingParameter                   = "MissingParameter"
	ErrUnknownParameter                   = "UnknownParameter"
	ErrAuthFailure                        = "AuthFailure"
	ErrAuthFailure_AccountNotFound        = "AuthFailure.AccountNotFound"
	ErrAuthFailure_TokenFailure           = "AuthFailure.TokenFailure"
	ErrAuthFailure_CGWFailure             = "AuthFailure.CGWFailure"
	ErrAuthFailure_MCFailure              = "AuthFailure.MCFailure"
	ErrAuthFailure_MockFailure            = "AuthFailure.MockFailure"
	ErrAuthFailure_GeneralFailure         = "AuthFailure.GeneralFailure"
	ErrAuthFailure_SkeyFailure            = "AuthFailure.SkeyFailure"
	ErrAuthFailure_CSRFEmpty              = "AuthFailure.CSRFEmpty"
	ErrAuthFailure_CSRFFailure            = "AuthFailure.CSRFFailure"
	ErrInternalError                      = "InternalError"
	ErrInternalError_DBFailure            = "InternalError.DBFailure"
	ErrInvalidAction                      = "InvalidAction"
	ErrUnauthorizedOperation              = "UnauthorizedOperation"
	ErrRequestLimitExceeded               = "RequestLimitExceeded"
	ErrNoSuchVersion                      = "NoSuchVersion"
	ErrUnsupportedRegion                  = "UnsupportedRegion"
	ErrUnsupportedOperation               = "UnsupportedOperation"
	ErrResourceNotFound                   = "ResourceNotFound"
	ErrLimitExceeded                      = "LimitExceeded"
	ErrResourceUnavailable                = "ResourceUnavailable"
	ErrResourceUnavailable_RequestIdEmpty = "ResourceUnavailable.RequestIdEmpty"
	ErrResourceUnavailable_ResponseEmpty  = "ResourceUnavailable.ResponseEmpty"
	ErrResourceInsufficient               = "ResourceInsufficient"
	ErrFailedOperation                    = "FailedOperation"
	ErrResourceInUse                      = "ResourceInUse"
	ErrDryRunOperation                    = "DryRunOperation"

	// Description, APIV3 common error code
	ErrDescDescInvalidParameter               = "参数格式/类型错误"
	ErrDescInvalidParameterValue              = "参数取值错误"
	ErrDescMissingParameter                   = "缺少参数错误"
	ErrDescUnknownParameter                   = "未知参数错误"
	ErrDescAuthFailure                        = "鉴权错误"
	ErrDescAuthFailure_AccountNotFound        = "鉴权错误，账号未找到"
	ErrDescAuthFailure_TokenFailure           = "鉴权错误，Token错误"
	ErrDescAuthFailure_CGWFailure             = "鉴权错误.CGW错误"
	ErrDescAuthFailure_MCFailure              = "鉴权错误.MC错误"
	ErrDescAuthFailure_MockFailure            = "鉴权错误.Mock错误"
	ErrDescAuthFailure_GeneralFailure         = "鉴权错误.一般性错误"
	ErrDescAuthFailure_SkeyFailure            = "鉴权错误.Skey错误"
	ErrDescAuthFailure_CSRFEmpty              = "鉴权错误.CSRF为空"
	ErrDescAuthFailure_CSRFFailure            = "鉴权错误.CSRF错误"
	ErrDescInternalError                      = "内部错误"
	ErrDescInternalError_DBFailure            = "内部错误，数据库错误"
	ErrDescInvalidAction                      = "接口不存在"
	ErrDescUnauthorizedOperation              = "未授权操作"
	ErrDescRequestLimitExceeded               = "请求频次超限"
	ErrDescNoSuchVersion                      = "接口版本不存在"
	ErrDescUnsupportedRegion                  = "区域不支持"
	ErrDescUnsupportedOperation               = "操作不支持"
	ErrDescResourceNotFound                   = "资源未找到"
	ErrDescLimitExceeded                      = "超过配额限制"
	ErrDescResourceUnavailable                = "资源不可用"
	ErrDescResourceUnavailable_RequestIdEmpty = "资源不可用，无RequestId"
	ErrDescResourceUnavailable_ResponseEmpty  = "资源不可用，无可用数据"
	ErrDescResourceInsufficient               = "资源不足"
	ErrDescFailedOperation                    = "操作失败"
	ErrDescResourceInUse                      = "资源被占用"
	ErrDescDryRunOperation                    = "DryRunOperation"
)

var ErrDescMap map[string]string

func init() {
	ErrDescMap = make(map[string]string, 0)
	ErrDescMap[ErrInvalidParameter] = ErrDescDescInvalidParameter
	ErrDescMap[ErrInvalidParameterValue] = ErrDescInvalidParameterValue
	ErrDescMap[ErrMissingParameter] = ErrDescMissingParameter
	ErrDescMap[ErrUnknownParameter] = ErrDescUnknownParameter
	ErrDescMap[ErrAuthFailure] = ErrDescAuthFailure
	ErrDescMap[ErrAuthFailure_AccountNotFound] = ErrDescAuthFailure_AccountNotFound
	ErrDescMap[ErrAuthFailure_TokenFailure] = ErrDescAuthFailure_TokenFailure
	ErrDescMap[ErrAuthFailure_CGWFailure] = ErrDescAuthFailure_CGWFailure
	ErrDescMap[ErrAuthFailure_MCFailure] = ErrDescAuthFailure_MCFailure
	ErrDescMap[ErrAuthFailure_GeneralFailure] = ErrDescAuthFailure_GeneralFailure
	ErrDescMap[ErrAuthFailure_SkeyFailure] = ErrDescAuthFailure_SkeyFailure
	ErrDescMap[ErrAuthFailure_CSRFFailure] = ErrDescAuthFailure_CSRFFailure
	ErrDescMap[ErrAuthFailure_CSRFEmpty] = ErrDescAuthFailure_CSRFEmpty
	ErrDescMap[ErrInternalError] = ErrDescInternalError
	ErrDescMap[ErrInternalError_DBFailure] = ErrDescInternalError_DBFailure
	ErrDescMap[ErrInvalidAction] = ErrDescInvalidAction
	ErrDescMap[ErrUnauthorizedOperation] = ErrDescUnauthorizedOperation
	ErrDescMap[ErrRequestLimitExceeded] = ErrDescRequestLimitExceeded
	ErrDescMap[ErrNoSuchVersion] = ErrDescNoSuchVersion
	ErrDescMap[ErrUnsupportedRegion] = ErrDescUnsupportedRegion
	ErrDescMap[ErrUnsupportedOperation] = ErrDescUnsupportedOperation
	ErrDescMap[ErrResourceNotFound] = ErrDescResourceNotFound
	ErrDescMap[ErrLimitExceeded] = ErrDescLimitExceeded
	ErrDescMap[ErrResourceUnavailable] = ErrDescResourceUnavailable
	ErrDescMap[ErrResourceUnavailable_RequestIdEmpty] = ErrDescResourceUnavailable_RequestIdEmpty
	ErrDescMap[ErrResourceUnavailable_ResponseEmpty] = ErrDescResourceUnavailable_ResponseEmpty
	ErrDescMap[ErrResourceInsufficient] = ErrDescResourceInsufficient
	ErrDescMap[ErrFailedOperation] = ErrDescFailedOperation
	ErrDescMap[ErrResourceInUse] = ErrDescResourceInUse
	ErrDescMap[ErrDryRunOperation] = ErrDescDryRunOperation
}

type ResourceName string // 项目各个资源定义，用于区分模拟登录时的权限

const (
	// 项目各项资源代码，主要用于区分模拟登录时的权限
	ResourceNameZhuJi         ResourceName = "zhuji"          // 珠玑项目资源，所有珠玑项目的接口（不含珠玑分析）均可视作该资源
	ResourceNameZhuJiAnalysis ResourceName = "zhuji_analysis" // 所有珠玑分析项目的接口均可视作该资源
	ResourceNameAll           ResourceName = "*"              // 所有项目的所有资源，主要用于超级管理员
	ResourceNameAny           ResourceName = "-"              // 任意资源（不校验资源权限，任何运营人员均可访问），主要用于通用接口如获取用户信息
)
