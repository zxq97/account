package api

import (
	"context"

	"google.golang.org/grpc"
)

type AccountClientImpl struct {
	client AccountClient
}

func NewAccountClientImpl(conn *grpc.ClientConn) *AccountClientImpl {
	return &AccountClientImpl{
		client: NewAccountClient(conn),
	}
}

func (aci *AccountClientImpl) Black(ctx context.Context, uid, touid int64) error {
	_, err := aci.client.Black(ctx, &BlackRequest{Uid: uid, ToUid: touid})
	return err
}

func (aci *AccountClientImpl) CancelBlack(ctx context.Context, uid, touid int64) error {
	_, err := aci.client.CancelBlack(ctx, &BlackRequest{Uid: uid, ToUid: touid})
	return err
}

func (aci *AccountClientImpl) CheckBlacked(ctx context.Context, uid int64, uids []int64) (map[int64]struct{}, error) {
	res, err := aci.client.CheckBlacked(ctx, &CheckBlackedRequest{Uid: uid, Uids: uids})
	if err != nil {
		return nil, err
	}
	m := make(map[int64]struct{}, len(res.BlackMap))
	for k, v := range res.BlackMap {
		if v {
			m[k] = struct{}{}
		}
	}
	return m, nil
}

func (aci *AccountClientImpl) GetUserWallet(ctx context.Context, uid int64) (*Wallet, error) {
	res, err := aci.client.GetUserWallet(ctx, &SingleRequest{Uid: uid})
	if err != nil {
		return nil, err
	}
	return walletDTO2VO(res.Wallet), nil
}

func (aci *AccountClientImpl) MultiGetUsersWallet(ctx context.Context, uids []int64) (map[int64]*Wallet, error) {
	res, err := aci.client.MultiGetUsersWallet(ctx, &MultiRequest{Uids: uids})
	if err != nil {
		return nil, err
	}
	return walletMapDTO2VO(res.WalletMap), nil
}

func (aci *AccountClientImpl) Withdrawal(ctx context.Context, txid string, uid, amount int64, from, to Plat) error {
	_, err := aci.client.Withdrawal(ctx, &TradeRequest{TxId: txid, Uid: uid, Amount: amount, From: from, To: to})
	return err
}

func (aci *AccountClientImpl) Store(ctx context.Context, txid string, uid, amount int64, from, to Plat) error {
	_, err := aci.client.Store(ctx, &TradeRequest{TxId: txid, Uid: uid, Amount: amount, From: from, To: to})
	return err
}

func (aci *AccountClientImpl) Transfer(ctx context.Context, txid string, uid, touid, amount int64, plat Plat) error {
	_, err := aci.client.Transfer(ctx, &MultiTradeRequest{TxId: txid, Uid: uid, ToUid: touid, Amount: amount, Plat: plat})
	return err
}

func (aci *AccountClientImpl) GetUser(ctx context.Context, uid int64) (*User, error) {
	res, err := aci.client.GetUser(ctx, &SingleRequest{Uid: uid})
	if err != nil {
		return nil, err
	}
	return userDTO2VO(res.User), nil
}

func (aci *AccountClientImpl) MultiGetUsers(ctx context.Context, uids []int64) (map[int64]*User, error) {
	res, err := aci.client.MultiGetUsers(ctx, &MultiRequest{Uids: uids})
	if err != nil {
		return nil, err
	}
	return userMapDTO2VO(res.UserMap), nil
}
