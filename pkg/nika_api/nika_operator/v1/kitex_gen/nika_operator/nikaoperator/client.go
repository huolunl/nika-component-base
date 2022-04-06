// Code generated by Kitex v0.2.1. DO NOT EDIT.

package nikaoperator

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/huolunl/nika-component-base/pkg/nika_api/nika_operator/v1/kitex_gen/nika_operator"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Apply(ctx context.Context, Req *nika_operator.ApplyRequest, callOptions ...callopt.Option) (r *nika_operator.ApplyResponse, err error)
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
	return &kNikaOperatorClient{
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

type kNikaOperatorClient struct {
	*kClient
}

func (p *kNikaOperatorClient) Apply(ctx context.Context, Req *nika_operator.ApplyRequest, callOptions ...callopt.Option) (r *nika_operator.ApplyResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Apply(ctx, Req)
}
