// Code generated by Kitex v0.2.1. DO NOT EDIT.

package chartservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/huolunl/nika-component-base/pkg/nika_api/nika_chartmuseum/v1/kitex_gen/nika_chartmuseum"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	ListChart(ctx context.Context, Req *nika_chartmuseum.ListChartRequest, callOptions ...callopt.Option) (r *nika_chartmuseum.ListChartResponse, err error)
	DeleteChart(ctx context.Context, Req *nika_chartmuseum.DeleteChartRequest, callOptions ...callopt.Option) (r *nika_chartmuseum.DeleteChartResponse, err error)
	UploadProvenanceFile(ctx context.Context, Req *nika_chartmuseum.UploadProvenanceFileRequest, callOptions ...callopt.Option) (r *nika_chartmuseum.UploadProvenanceFileResponse, err error)
	GetChartByName(ctx context.Context, Req *nika_chartmuseum.GetChartByNameRequest, callOptions ...callopt.Option) (r *nika_chartmuseum.GetChartByNameResponse, err error)
	PageChartList(ctx context.Context, Req *nika_chartmuseum.PageChartListRequest, callOptions ...callopt.Option) (r *nika_chartmuseum.PageChartListResponse, err error)
	UploadChart(ctx context.Context, Req *nika_chartmuseum.UploadChartRequest, callOptions ...callopt.Option) (r *nika_chartmuseum.UploadChartResponse, err error)
	DescribeChartVersion(ctx context.Context, Req *nika_chartmuseum.DescribeChartVersionRequest, callOptions ...callopt.Option) (r *nika_chartmuseum.DescribeChartVersionResponse, err error)
	CheckChartIsExist(ctx context.Context, Req *nika_chartmuseum.CheckChartIsExistRequest, callOptions ...callopt.Option) (r *nika_chartmuseum.CheckChartIsExistResponse, err error)
	CheckChartVersionIsExist(ctx context.Context, Req *nika_chartmuseum.CheckChartVersionIsExistRequest, callOptions ...callopt.Option) (r *nika_chartmuseum.CheckChartVersionIsExistResponse, err error)
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
	return &kChartServiceClient{
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

type kChartServiceClient struct {
	*kClient
}

func (p *kChartServiceClient) ListChart(ctx context.Context, Req *nika_chartmuseum.ListChartRequest, callOptions ...callopt.Option) (r *nika_chartmuseum.ListChartResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListChart(ctx, Req)
}

func (p *kChartServiceClient) DeleteChart(ctx context.Context, Req *nika_chartmuseum.DeleteChartRequest, callOptions ...callopt.Option) (r *nika_chartmuseum.DeleteChartResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteChart(ctx, Req)
}

func (p *kChartServiceClient) UploadProvenanceFile(ctx context.Context, Req *nika_chartmuseum.UploadProvenanceFileRequest, callOptions ...callopt.Option) (r *nika_chartmuseum.UploadProvenanceFileResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadProvenanceFile(ctx, Req)
}

func (p *kChartServiceClient) GetChartByName(ctx context.Context, Req *nika_chartmuseum.GetChartByNameRequest, callOptions ...callopt.Option) (r *nika_chartmuseum.GetChartByNameResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetChartByName(ctx, Req)
}

func (p *kChartServiceClient) PageChartList(ctx context.Context, Req *nika_chartmuseum.PageChartListRequest, callOptions ...callopt.Option) (r *nika_chartmuseum.PageChartListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PageChartList(ctx, Req)
}

func (p *kChartServiceClient) UploadChart(ctx context.Context, Req *nika_chartmuseum.UploadChartRequest, callOptions ...callopt.Option) (r *nika_chartmuseum.UploadChartResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UploadChart(ctx, Req)
}

func (p *kChartServiceClient) DescribeChartVersion(ctx context.Context, Req *nika_chartmuseum.DescribeChartVersionRequest, callOptions ...callopt.Option) (r *nika_chartmuseum.DescribeChartVersionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DescribeChartVersion(ctx, Req)
}

func (p *kChartServiceClient) CheckChartIsExist(ctx context.Context, Req *nika_chartmuseum.CheckChartIsExistRequest, callOptions ...callopt.Option) (r *nika_chartmuseum.CheckChartIsExistResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckChartIsExist(ctx, Req)
}

func (p *kChartServiceClient) CheckChartVersionIsExist(ctx context.Context, Req *nika_chartmuseum.CheckChartVersionIsExistRequest, callOptions ...callopt.Option) (r *nika_chartmuseum.CheckChartVersionIsExistResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckChartVersionIsExist(ctx, Req)
}
