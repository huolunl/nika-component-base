// Code generated by Kitex v0.2.1. DO NOT EDIT.

package nikauser

import (
	"github.com/cloudwego/kitex/server"
	"github.com/huolunl/nika-component-base/pkg/nika_api/nika_user/v1/kitex_gen/nika_user"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler nika_user.NikaUser, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}