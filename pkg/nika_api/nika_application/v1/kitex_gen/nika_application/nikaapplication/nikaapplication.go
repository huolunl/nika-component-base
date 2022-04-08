// Code generated by Kitex v0.2.1. DO NOT EDIT.

package nikaapplication

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/huolunl/nika-component-base/pkg/nika_api/nika_application/v1/kitex_gen/nika_application"
	"google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return nikaApplicationServiceInfo
}

var nikaApplicationServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "NikaApplication"
	handlerType := (*nika_application.NikaApplication)(nil)
	methods := map[string]kitex.MethodInfo{
		"CreateApplication": kitex.NewMethodInfo(createApplicationHandler, newCreateApplicationArgs, newCreateApplicationResult, false),
		"UpdateApplication": kitex.NewMethodInfo(updateApplicationHandler, newUpdateApplicationArgs, newUpdateApplicationResult, false),
		"GetApplication":    kitex.NewMethodInfo(getApplicationHandler, newGetApplicationArgs, newGetApplicationResult, false),
		"ListApplication":   kitex.NewMethodInfo(listApplicationHandler, newListApplicationArgs, newListApplicationResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.2.1",
		Extra:           extra,
	}
	return svcInfo
}

func createApplicationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(nika_application.CreateApplicationRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(nika_application.NikaApplication).CreateApplication(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *CreateApplicationArgs:
		success, err := handler.(nika_application.NikaApplication).CreateApplication(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*CreateApplicationResult)
		realResult.Success = success
	}
	return nil
}
func newCreateApplicationArgs() interface{} {
	return &CreateApplicationArgs{}
}

func newCreateApplicationResult() interface{} {
	return &CreateApplicationResult{}
}

type CreateApplicationArgs struct {
	Req *nika_application.CreateApplicationRequest
}

func (p *CreateApplicationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in CreateApplicationArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *CreateApplicationArgs) Unmarshal(in []byte) error {
	msg := new(nika_application.CreateApplicationRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var CreateApplicationArgs_Req_DEFAULT *nika_application.CreateApplicationRequest

func (p *CreateApplicationArgs) GetReq() *nika_application.CreateApplicationRequest {
	if !p.IsSetReq() {
		return CreateApplicationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *CreateApplicationArgs) IsSetReq() bool {
	return p.Req != nil
}

type CreateApplicationResult struct {
	Success *nika_application.CreateApplicationResponse
}

var CreateApplicationResult_Success_DEFAULT *nika_application.CreateApplicationResponse

func (p *CreateApplicationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in CreateApplicationResult")
	}
	return proto.Marshal(p.Success)
}

func (p *CreateApplicationResult) Unmarshal(in []byte) error {
	msg := new(nika_application.CreateApplicationResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *CreateApplicationResult) GetSuccess() *nika_application.CreateApplicationResponse {
	if !p.IsSetSuccess() {
		return CreateApplicationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *CreateApplicationResult) SetSuccess(x interface{}) {
	p.Success = x.(*nika_application.CreateApplicationResponse)
}

func (p *CreateApplicationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func updateApplicationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(nika_application.UpdateApplicationRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(nika_application.NikaApplication).UpdateApplication(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UpdateApplicationArgs:
		success, err := handler.(nika_application.NikaApplication).UpdateApplication(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UpdateApplicationResult)
		realResult.Success = success
	}
	return nil
}
func newUpdateApplicationArgs() interface{} {
	return &UpdateApplicationArgs{}
}

func newUpdateApplicationResult() interface{} {
	return &UpdateApplicationResult{}
}

type UpdateApplicationArgs struct {
	Req *nika_application.UpdateApplicationRequest
}

func (p *UpdateApplicationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in UpdateApplicationArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *UpdateApplicationArgs) Unmarshal(in []byte) error {
	msg := new(nika_application.UpdateApplicationRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UpdateApplicationArgs_Req_DEFAULT *nika_application.UpdateApplicationRequest

func (p *UpdateApplicationArgs) GetReq() *nika_application.UpdateApplicationRequest {
	if !p.IsSetReq() {
		return UpdateApplicationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UpdateApplicationArgs) IsSetReq() bool {
	return p.Req != nil
}

type UpdateApplicationResult struct {
	Success *nika_application.UpdateApplicationResponse
}

var UpdateApplicationResult_Success_DEFAULT *nika_application.UpdateApplicationResponse

func (p *UpdateApplicationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in UpdateApplicationResult")
	}
	return proto.Marshal(p.Success)
}

func (p *UpdateApplicationResult) Unmarshal(in []byte) error {
	msg := new(nika_application.UpdateApplicationResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UpdateApplicationResult) GetSuccess() *nika_application.UpdateApplicationResponse {
	if !p.IsSetSuccess() {
		return UpdateApplicationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UpdateApplicationResult) SetSuccess(x interface{}) {
	p.Success = x.(*nika_application.UpdateApplicationResponse)
}

func (p *UpdateApplicationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getApplicationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(nika_application.GetApplicationRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(nika_application.NikaApplication).GetApplication(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetApplicationArgs:
		success, err := handler.(nika_application.NikaApplication).GetApplication(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetApplicationResult)
		realResult.Success = success
	}
	return nil
}
func newGetApplicationArgs() interface{} {
	return &GetApplicationArgs{}
}

func newGetApplicationResult() interface{} {
	return &GetApplicationResult{}
}

type GetApplicationArgs struct {
	Req *nika_application.GetApplicationRequest
}

func (p *GetApplicationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetApplicationArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetApplicationArgs) Unmarshal(in []byte) error {
	msg := new(nika_application.GetApplicationRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetApplicationArgs_Req_DEFAULT *nika_application.GetApplicationRequest

func (p *GetApplicationArgs) GetReq() *nika_application.GetApplicationRequest {
	if !p.IsSetReq() {
		return GetApplicationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetApplicationArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetApplicationResult struct {
	Success *nika_application.GetApplicationResponse
}

var GetApplicationResult_Success_DEFAULT *nika_application.GetApplicationResponse

func (p *GetApplicationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetApplicationResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetApplicationResult) Unmarshal(in []byte) error {
	msg := new(nika_application.GetApplicationResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetApplicationResult) GetSuccess() *nika_application.GetApplicationResponse {
	if !p.IsSetSuccess() {
		return GetApplicationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetApplicationResult) SetSuccess(x interface{}) {
	p.Success = x.(*nika_application.GetApplicationResponse)
}

func (p *GetApplicationResult) IsSetSuccess() bool {
	return p.Success != nil
}

func listApplicationHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(nika_application.ListApplicationRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(nika_application.NikaApplication).ListApplication(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *ListApplicationArgs:
		success, err := handler.(nika_application.NikaApplication).ListApplication(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListApplicationResult)
		realResult.Success = success
	}
	return nil
}
func newListApplicationArgs() interface{} {
	return &ListApplicationArgs{}
}

func newListApplicationResult() interface{} {
	return &ListApplicationResult{}
}

type ListApplicationArgs struct {
	Req *nika_application.ListApplicationRequest
}

func (p *ListApplicationArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in ListApplicationArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *ListApplicationArgs) Unmarshal(in []byte) error {
	msg := new(nika_application.ListApplicationRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListApplicationArgs_Req_DEFAULT *nika_application.ListApplicationRequest

func (p *ListApplicationArgs) GetReq() *nika_application.ListApplicationRequest {
	if !p.IsSetReq() {
		return ListApplicationArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListApplicationArgs) IsSetReq() bool {
	return p.Req != nil
}

type ListApplicationResult struct {
	Success *nika_application.ListApplicationResponse
}

var ListApplicationResult_Success_DEFAULT *nika_application.ListApplicationResponse

func (p *ListApplicationResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in ListApplicationResult")
	}
	return proto.Marshal(p.Success)
}

func (p *ListApplicationResult) Unmarshal(in []byte) error {
	msg := new(nika_application.ListApplicationResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListApplicationResult) GetSuccess() *nika_application.ListApplicationResponse {
	if !p.IsSetSuccess() {
		return ListApplicationResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListApplicationResult) SetSuccess(x interface{}) {
	p.Success = x.(*nika_application.ListApplicationResponse)
}

func (p *ListApplicationResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CreateApplication(ctx context.Context, Req *nika_application.CreateApplicationRequest) (r *nika_application.CreateApplicationResponse, err error) {
	var _args CreateApplicationArgs
	_args.Req = Req
	var _result CreateApplicationResult
	if err = p.c.Call(ctx, "CreateApplication", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UpdateApplication(ctx context.Context, Req *nika_application.UpdateApplicationRequest) (r *nika_application.UpdateApplicationResponse, err error) {
	var _args UpdateApplicationArgs
	_args.Req = Req
	var _result UpdateApplicationResult
	if err = p.c.Call(ctx, "UpdateApplication", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetApplication(ctx context.Context, Req *nika_application.GetApplicationRequest) (r *nika_application.GetApplicationResponse, err error) {
	var _args GetApplicationArgs
	_args.Req = Req
	var _result GetApplicationResult
	if err = p.c.Call(ctx, "GetApplication", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListApplication(ctx context.Context, Req *nika_application.ListApplicationRequest) (r *nika_application.ListApplicationResponse, err error) {
	var _args ListApplicationArgs
	_args.Req = Req
	var _result ListApplicationResult
	if err = p.c.Call(ctx, "ListApplication", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
