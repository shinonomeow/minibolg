package options

import (
	"errors"
	"fmt"
	"time"

	"github.com/shinonomeow/miniblog/internal/apiserver"
	"github.com/spf13/pflag"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/sets"
)

var availableServerModes = sets.New(
	"grpc",
	"grpc-gateway",
	"gin",
)

type ServerOptions struct {
	// ServerMode 定义服务器模式: grpc、Gin HTTP、HTTP Reverse Proxy
	ServerMode string `json:"server-mode" mapstructure:"server-mode"`
	// JWTKey 定义用于 JWT 认证的密钥
	JWTKey string `json:"jwt-key"     mapstructure:"jwt-key"`
	// Expiration 定义 JWT 令牌的过期时间
	Expiration time.Duration `json:"expiration"  mapstructure:"expiration"`
}

func (o *ServerOptions) Config() (*apiserver.Config, error) {
	return &apiserver.Config{
		ServerMode: o.ServerMode,
		JWTKey:     o.JWTKey,
		Expiration: o.Expiration,
	}, nil
}

func NewServerOptions() *ServerOptions {
	return &ServerOptions{
		ServerMode: "grpc-gateway",
		JWTKey:     "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5",
		Expiration: 2 * time.Hour,
	}
}

func (o *ServerOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(
		&o.ServerMode,
		"server-mode",
		o.ServerMode,
		fmt.Sprintf("Server mode, available options:%v", availableServerModes.UnsortedList()),
	)
	fs.StringVar(&o.JWTKey, "jwt-key", o.JWTKey, "JWT signing key. Must be at least 6 characters long.")
	fs.DurationVar(&o.Expiration, "expiration", o.Expiration, "The expiration duration of JWT tokens.")
}

// Validate 校验 ServerOptions 中的选项是否合法.
func (o *ServerOptions) Validate() error {
	errs := []error{}

	// 校验 ServerMode 是否有效
	if !availableServerModes.Has(o.ServerMode) {
		errs = append(errs, fmt.Errorf("invalid server mode: must be one of %v", availableServerModes.UnsortedList()))
	}

	// 校验 JWTKey 长度
	if len(o.JWTKey) < 6 {
		errs = append(errs, errors.New("JWTKey must be at least 6 characters long"))
	}

	// 合并所有错误并返回
	return utilerrors.NewAggregate(errs)
}

