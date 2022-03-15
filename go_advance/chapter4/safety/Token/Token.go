package Token

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// 创建Token验证对象
type Authentication struct {
	User     string
	Password string
}

// 实现grpc.PreRPCCredentials接口
func (a *Authentication) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}

// 使用安全链接
func (a *Authentication) RequireTransportSecurity() bool {
	return true
}

// 定义验证方法
func (a *Authentication) Auth(ctx context.Context) error {
	// 从上下文中读取信息
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("missing credentials")
	}

	// 获取值
	var appid string
	var appkey string

	if val, ok := md["user"]; ok {
		appid = val[0]
	}

	if val, ok := md["password"]; ok {
		appkey = val[0]
	}
	// 验证值
	if appid != a.User || appkey != a.Password {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}
	return nil
}
