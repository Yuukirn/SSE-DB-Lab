package util

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
)

func NewErrResp(c *app.RequestContext, err error) {
	NewResp(c, -1, err.Error(), nil)
}

func NewOkResp(c *app.RequestContext, msg string, data any) {
	NewResp(c, 0, msg, data)
}

func NewResp(c *app.RequestContext, statusCode int, msg string, data any) {
	c.JSON(http.StatusOK, utils.H{
		"code": statusCode,
		"msg":  msg,
		"data": data,
	})
}

func Hash(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}
