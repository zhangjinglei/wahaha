package warden

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/tidwall/gjson"
	nmd "github.com/zhangjinglei/wahaha/pkg/net/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

func GrpcAuthMiddleWare() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, args *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		incomingContext, ok := metadata.FromIncomingContext(ctx)
		if ok {
			authjwt, exist := incomingContext["authorization"]
			if exist {
				parts := strings.Split(authjwt[0], ".")
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

			}
		}

		resp, err = handler(ctx, req)
		return
	}
}
