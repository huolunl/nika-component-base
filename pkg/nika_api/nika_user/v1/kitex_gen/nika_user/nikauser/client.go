// Code generated by Kitex v0.2.0. DO NOT EDIT.

package nikauser

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/huolunl/nika-component-base/pkg/nika_api/nika_user/v1/kitex_gen/nika_user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Login(ctx context.Context, Req *nika_user.LoginRequest, callOptions ...callopt.Option) (r *nika_user.LoginResponse, err error)
	CreateTenant(ctx context.Context, Req *nika_user.CreateTenantRequest, callOptions ...callopt.Option) (r *nika_user.CreateTenantResponse, err error)
	CreateUser(ctx context.Context, Req *nika_user.CreateUserRequest, callOptions ...callopt.Option) (r *nika_user.CreateUserResponse, err error)
	Authorization(ctx context.Context, Req *nika_user.AuthorizationRequest, callOptions ...callopt.Option) (r *nika_user.AuthorizationResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kNikaUserClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kNikaUserClient struct {
	*kClient
}

func (p *kNikaUserClient) Login(ctx context.Context, Req *nika_user.LoginRequest, callOptions ...callopt.Option) (r *nika_user.LoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Login(ctx, Req)
}

func (p *kNikaUserClient) CreateTenant(ctx context.Context, Req *nika_user.CreateTenantRequest, callOptions ...callopt.Option) (r *nika_user.CreateTenantResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateTenant(ctx, Req)
}

func (p *kNikaUserClient) CreateUser(ctx context.Context, Req *nika_user.CreateUserRequest, callOptions ...callopt.Option) (r *nika_user.CreateUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUser(ctx, Req)
}

func (p *kNikaUserClient) Authorization(ctx context.Context, Req *nika_user.AuthorizationRequest, callOptions ...callopt.Option) (r *nika_user.AuthorizationResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Authorization(ctx, Req)
}
