package middlewares

import (
	"bytes"
	"context"
	json2 "encoding/json"
	"fmt"
	"github.com/NETkiddy/common-go/log"
	"github.com/NETkiddy/nft-svr/common"
	"github.com/NETkiddy/nft-svr/common/cycleImportModels"
	"github.com/NETkiddy/nft-svr/provider/protocol/json"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"time"
)

type CheckSessionByToken struct {
	RequestId string
	Sid       string
}

func CheckUin() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		bodyBytes, err := c.GetRawData()
		if err != nil {
			log.LoggerWrapperWithCaller().Errorf("GetRawData failed: %s", err.Error())
			ret := map[string]interface{}{
				"Response": map[string]interface{}{
					"Error": map[string]interface{}{
						"Code":    common.ErrInvalidParameter,
						"Message": "获取请求体失败",
					},
				},
			}
			c.JSON(200, ret)
			c.Abort()
			return
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		bodyJsonMap := make(map[string]interface{})
		if err := json2.Unmarshal(bodyBytes, &bodyJsonMap); err != nil {
			log.LoggerWrapperWithCaller().Errorf("json.Unmarshal failed: %s", err.Error())
			ret := map[string]interface{}{
				"Response": map[string]interface{}{
					"Error": map[string]interface{}{
						"Code":    common.ErrInvalidParameter,
						"Message": "Json参数错误",
					},
				},
			}
			c.JSON(200, ret)
			c.Abort()
			return
		}

		uin := ""
		if uinInterface, found := bodyJsonMap["Uin"]; found {
			ok := false
			uin, ok = uinInterface.(string)
			if !ok || uin == "" {
				log.LoggerFromContextWithCaller(ctx).Errorf("Missing Param Uin")
				ret := map[string]interface{}{
					"Response": map[string]interface{}{
						"Error": map[string]interface{}{
							"Code":    common.ErrInvalidParameter,
							"Message": "Missing Param Uin",
						},
					},
				}
				c.JSON(200, ret)
				c.Abort()
				return
			}
			//panic(&glue.ExceptionForAPIV3{Code: common.ErrInternalError, Message: "Missing Param Uin"})
		}

		provider, err := GetProviderByUin(ctx, uin)
		if err != nil {
			log.LoggerFromContextWithCaller(ctx).Errorf("GetProviderByUin failed")
			ret := map[string]interface{}{
				"Response": map[string]interface{}{
					"Error": map[string]interface{}{
						"Code":    common.ErrInvalidParameter,
						"Message": "GetProviderByUin failed",
					},
				},
			}
			c.JSON(200, ret)
			c.Abort()
			return

		}

		c.Set("BEE-ACCOUNT_ID", provider.ID)
		/*
			newCtx := context.WithValue(ctx, "BEE-ACCOUNT_ID", provider.ID)
			log.LoggerFromContextWithCaller(ctx).Debugf("--------newCtx: %v", newCtx)
			newReq := c.Request.WithContext(newCtx)
			c.Request = newReq
		*/
		c.Next()

	}
}

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		sid := c.GetHeader("sid")
		var skey string
		var uid uint
		if sid != "" {
			sm, err := GetSessionBySid(ctx, sid)
			if err != nil {
				log.LoggerFromContextWithCaller(ctx).Errorf("GetSessionBySid failed")
				ret := map[string]interface{}{
					"Response": map[string]interface{}{
						"Error": map[string]interface{}{
							"Code":    common.ErrInvalidParameter,
							"Message": "GetProviderByUin failed",
						},
					},
				}
				c.JSON(200, ret)
				c.Abort()
				return
			}
			log.LoggerFromContextWithCaller(ctx).Debugf("sm:%v", sm)
			if d, fd := sm["sessionKey"]; fd {
				skey = d.(string)
			}
			if d, fd := sm["id"]; fd {
				uid = uint(d.(float64))
			}
		}
		c.Set("BEE-SID", sid)
		c.Set("BEE-SESSION_KEY", skey)
		c.Set("BEE-UID", uid)
		log.LoggerFromContextWithCaller(ctx).Debugf("Set BEE-SID: %v", sid)
		log.LoggerFromContextWithCaller(ctx).Debugf("Set BEE-SESSION_KEY: %v", skey)
		log.LoggerFromContextWithCaller(ctx).Debugf("Set BEE-UID: %v", uid)

		c.Next()

	}
}

func GetProviderByUin(ctx context.Context, uin string) (provider cycleImportModels.Provider, err error) {

	proxy := ""
	providerAddr := "http://currykitty.com/bee/provider/CheckProviderByToken"

	var restyReq *resty.Request
	if len(proxy) > 0 {
		restyReq = resty.New().SetRetryCount(3).SetTimeout(10 * time.Second).SetRetryWaitTime(100 * time.Millisecond).SetRetryMaxWaitTime(1 * time.Second).SetProxy(proxy).R()
	} else {
		restyReq = resty.New().SetRetryCount(3).SetTimeout(10 * time.Second).SetRetryWaitTime(100 * time.Millisecond).SetRetryMaxWaitTime(1 * time.Second).R()
	}

	providerReq := json.QueryProviderRequest{
		RequestId: "TODO", //TODO, ctx.Value()
		Uin:       uin,
	}
	providerReqBuff, err := json2.Marshal(providerReq)
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("GetProviderByUin json.Marshal failed, err info: %s", err.Error())
		return
	}

	value, err := restyReq.
		SetHeader("Content-Type", "application/json").
		SetBody(providerReqBuff).
		Post(providerAddr)
	log.LoggerWrapperWithCaller().Debugf("GetProviderByUin bodyStr: %v", string(providerReqBuff))
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("GetProviderByUin resty.Get[%s] failed, err info: %s", providerAddr, err.Error())
		return
	}
	log.LoggerWrapperWithCaller().Debugf("GetProviderByUin respStr: %v", string(value.Body()))

	respData := make(map[string]map[string]json.QueryProviderResponse, 0)
	err = json2.Unmarshal(value.Body(), &respData)
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("GetProviderByUin json.Unmarshal failed, err info: %s", err.Error())
		return
	}
	resp := respData["Response"]["Data"]

	if resp.ProviderList == nil && len(resp.ProviderList) == 0 {
		err = fmt.Errorf("GetProviderByUin resp failed")
		log.LoggerFromContextWithCaller(ctx).Errorf("GetProviderByUin resp failed: %v", resp)
		return
	}

	provider = resp.ProviderList[0]
	return
}

func GetSessionBySid(ctx context.Context, sid string) (sm map[string]interface{}, err error) {
	sm = make(map[string]interface{})
	checkAddr := "https://currykitty.com/bee/account/CheckSessionByToken"

	var restyReq *resty.Request
	restyReq = resty.New().SetRetryCount(3).SetTimeout(10 * time.Second).SetRetryWaitTime(100 * time.Millisecond).SetRetryMaxWaitTime(1 * time.Second).R()

	checkReq := CheckSessionByToken{
		RequestId: "TODO", //TODO, ctx.Value()
		Sid:       sid,
	}
	checkReqBuff, err := json2.Marshal(checkReq)
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("CheckSessionByToken json.Marshal failed, err info: %s", err.Error())
		return
	}

	value, err := restyReq.
		SetHeader("Content-Type", "application/json").
		SetBody(checkReqBuff).
		Post(checkAddr)
	log.LoggerWrapperWithCaller().Debugf("CheckSessionByToken bodyStr: %v", string(checkReqBuff))
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("CheckSessionByToken resty.Get[%s] failed, err info: %s", checkAddr, err.Error())
		return
	}
	log.LoggerWrapperWithCaller().Debugf("CheckSessionByToken respStr: %v", string(value.Body()))

	respData := make(map[string]map[string]map[string]map[string]interface{}, 0)
	err = json2.Unmarshal(value.Body(), &respData)
	if err != nil {
		log.LoggerFromContextWithCaller(ctx).Errorf("CheckSessionByToken json.Unmarshal failed, err info: %s", err.Error())
		return
	}
	sm = respData["Response"]["Data"]["session_map"]

	return
}
