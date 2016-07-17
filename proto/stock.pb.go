// Code generated by protoc-gen-go.
// source: github.com/chazsmi/stock-service/proto/stock.proto
// DO NOT EDIT!

/*
Package stock is a generated protocol buffer package.

It is generated from these files:
	github.com/chazsmi/stock-service/proto/stock.proto

It has these top-level messages:
	StockRequest
	StockResponse
*/
package stock

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StockRequest struct {
	Sku string `protobuf:"bytes,1,opt,name=sku" json:"sku,omitempty"`
}

func (m *StockRequest) Reset()                    { *m = StockRequest{} }
func (m *StockRequest) String() string            { return proto.CompactTextString(m) }
func (*StockRequest) ProtoMessage()               {}
func (*StockRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StockResponse struct {
	Sku    string `protobuf:"bytes,1,opt,name=sku" json:"sku,omitempty"`
	Amount int32  `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *StockResponse) Reset()                    { *m = StockResponse{} }
func (m *StockResponse) String() string            { return proto.CompactTextString(m) }
func (*StockResponse) ProtoMessage()               {}
func (*StockResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*StockRequest)(nil), "StockRequest")
	proto.RegisterType((*StockResponse)(nil), "StockResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Stock service

type StockClient interface {
	Check(ctx context.Context, in *StockRequest, opts ...client.CallOption) (*StockResponse, error)
}

type stockClient struct {
	c           client.Client
	serviceName string
}

func NewStockClient(serviceName string, c client.Client) StockClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "stock"
	}
	return &stockClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *stockClient) Check(ctx context.Context, in *StockRequest, opts ...client.CallOption) (*StockResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Stock.Check", in)
	out := new(StockResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Stock service

type StockHandler interface {
	Check(context.Context, *StockRequest, *StockResponse) error
}

func RegisterStockHandler(s server.Server, hdlr StockHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Stock{hdlr}, opts...))
}

type Stock struct {
	StockHandler
}

func (h *Stock) Check(ctx context.Context, in *StockRequest, out *StockResponse) error {
	return h.StockHandler.Check(ctx, in, out)
}

func init() { proto.RegisterFile("github.com/chazsmi/stock-service/proto/stock.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xce, 0x48, 0xac, 0x2a, 0xce, 0xcd, 0xd4, 0x2f,
	0x2e, 0xc9, 0x4f, 0xce, 0xd6, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2f, 0x28, 0xca,
	0x2f, 0xc9, 0x87, 0x88, 0xe9, 0x81, 0xd9, 0x4a, 0x0a, 0x5c, 0x3c, 0xc1, 0x20, 0x6e, 0x50, 0x6a,
	0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x00, 0x17, 0x73, 0x71, 0x76, 0xa9, 0x04, 0xa3, 0x02, 0xa3,
	0x06, 0x67, 0x10, 0x88, 0xa9, 0x64, 0xc9, 0xc5, 0x0b, 0x55, 0x51, 0x5c, 0x90, 0x9f, 0x57, 0x9c,
	0x8a, 0xa9, 0x44, 0x48, 0x8c, 0x8b, 0x2d, 0x31, 0x37, 0xbf, 0x34, 0xaf, 0x44, 0x82, 0x09, 0x28,
	0xc8, 0x1a, 0x04, 0xe5, 0x19, 0x19, 0x72, 0xb1, 0x82, 0xb5, 0x0a, 0x69, 0x70, 0xb1, 0x3a, 0x67,
	0xa4, 0x02, 0x19, 0xbc, 0x7a, 0xc8, 0xb6, 0x49, 0xf1, 0xe9, 0xa1, 0x18, 0xad, 0xc4, 0x90, 0xc4,
	0x06, 0x76, 0x96, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x3e, 0xab, 0x7d, 0xcc, 0x00, 0x00,
	0x00,
}
