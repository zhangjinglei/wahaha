package blademaster

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/tidwall/gjson"
	nmd "github.com/zhangjinglei/wahaha/pkg/net/metadata"
	"strings"
)

func HttpAuthMiddleWare() HandlerFunc {
	return func(ctx *Context) {

		authjwt := ctx.Request.Header.Get("authorization")
		parts := strings.Split(authjwt, ".")
		if len(parts) > 1 {
			segment, err := jwt.DecodeSegment(parts[1])
			if err == nil {
				substr := gjson.Get(string(segment), "sub").String()
				if substr != "" {
					fromContext, ok := nmd.FromContext(ctx)
					if ok {
						fromContext["bdw_user"] = substr
					}
				}
			}
		}
		ctx.Next()
	}

}
