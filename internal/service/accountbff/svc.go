package accountbff

import (
	"context"

	"github.com/zxq97/account/api"
	"google.golang.org/grpc"
)

type AccountBFF struct {
}

func InitAccountBFF(conf *AccountBffConfig, conn *grpc.ClientConn) (*AccountBFF, error) {
	return &AccountBFF{}, nil
}

func (bff *AccountBFF) Black(ctx context.Context, req *api.BlackRequest) (*api.EmptyResponse, error) {
	return &api.EmptyResponse{}, nil
}

func (bff *AccountBFF) CancelBlack(ctx context.Context, req *api.BlackRequest) (*api.EmptyResponse, error) {
	return &api.EmptyResponse{}, nil
}

func (bff *AccountBFF) CheckBlacked(ctx context.Context, req *api.CheckBlackedRequest) (*api.CheckBlackedResponse, error) {
	return &api.CheckBlackedResponse{}, nil
}

func (bff *AccountBFF) GetUserWallet(ctx context.Context, req *api.SingleRequest) (*api.GetUserWalletResponse, error) {
	return &api.GetUserWalletResponse{}, nil
}

func (bff *AccountBFF) MultiGetUsersWallet(ctx context.Context, req *api.MultiRequest) (*api.MultiGetUsersWalletResponse, error) {
	return &api.MultiGetUsersWalletResponse{}, nil
}

func (bff *AccountBFF) Withdrawal(ctx context.Context, req *api.TradeRequest) (*api.EmptyResponse, error) {
	return &api.EmptyResponse{}, nil
}

func (bff *AccountBFF) Store(ctx context.Context, req *api.TradeRequest) (*api.EmptyResponse, error) {
	return &api.EmptyResponse{}, nil
}

func (bff *AccountBFF) Transfer(ctx context.Context, req *api.MultiTradeRequest) (*api.EmptyResponse, error) {
	return &api.EmptyResponse{}, nil
}

func (bff *AccountBFF) GetUser(ctx context.Context, req *api.SingleRequest) (*api.GetUserResponse, error) {
	return &api.GetUserResponse{}, nil
}

func (bff *AccountBFF) MultiGetUsers(ctx context.Context, req *api.MultiRequest) (*api.MultiGetUsersResponse, error) {
	return &api.MultiGetUsersResponse{}, nil
}
