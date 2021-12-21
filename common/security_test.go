package common

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetSignature(t *testing.T) {
	key := "44CF9590006BF252F707"
	secret := "OtxrzxIsfpFjA7SwPzILwy8Bw21TLhquhboDYROV"
	method := "GET"
	endpoint := "/api/v1/token_classes"
	content := ""
	content_type := "application/json"
	gmt := "Tue, 06 Jul 2021 00:00:34 GMT"

	signature := GetSignature(secret, method, endpoint, content, gmt, content_type)

	authorization := fmt.Sprintf("%v %v:%v", "NFT", key, signature)

	assert.Equal(t, "NFT 44CF9590006BF252F707:SXc3VHXXbU08qzYdAm1RvwMWaUw=", authorization)

}
