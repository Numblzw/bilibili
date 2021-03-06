// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/RoomRecommend.proto

package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "go-common/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports

// =======================
// RoomRecommend Interface
// =======================

type RoomRecommendRPCClient interface {
	// * app关播播放器推荐
	//
	GetPlayerList(ctx context.Context, req *RoomRecommendGetPlayerListReq, opts ...liverpc.CallOption) (resp *RoomRecommendGetPlayerListResp, err error)

	// * 获取首页推荐主播数据,可强推可刷新
	//
	ClientRecStrong(ctx context.Context, req *RoomRecommendClientRecStrongReq, opts ...liverpc.CallOption) (resp *RoomRecommendClientRecStrongResp, err error)
}

// =============================
// RoomRecommend Live Rpc Client
// =============================

type roomRecommendRPCClient struct {
	client *liverpc.Client
}

// NewRoomRecommendRPCClient creates a client that implements the RoomRecommendRPCClient interface.
func NewRoomRecommendRPCClient(client *liverpc.Client) RoomRecommendRPCClient {
	return &roomRecommendRPCClient{
		client: client,
	}
}

func (c *roomRecommendRPCClient) GetPlayerList(ctx context.Context, in *RoomRecommendGetPlayerListReq, opts ...liverpc.CallOption) (*RoomRecommendGetPlayerListResp, error) {
	out := new(RoomRecommendGetPlayerListResp)
	err := doRPCRequest(ctx, c.client, 1, "RoomRecommend.getPlayerList", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomRecommendRPCClient) ClientRecStrong(ctx context.Context, in *RoomRecommendClientRecStrongReq, opts ...liverpc.CallOption) (*RoomRecommendClientRecStrongResp, error) {
	out := new(RoomRecommendClientRecStrongResp)
	err := doRPCRequest(ctx, c.client, 1, "RoomRecommend.clientRecStrong", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}
