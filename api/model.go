package api

type User struct {
	UID     int64  `json:"uid"`
	Name    string `json:"name"`
	Profile string `json:"profile"`
}

type Wallet struct {
	UID           int64 `json:"uid"`
	Balance       int64 `json:"balance"`
	FreezeBalance int64 `json:"freeze_balance"`
	Pay           int64 `json:"pay"`
	FreezePay     int64 `json:"freeze_pay"`
}

func userDTO2VO(u *Userinfo) *User {
	return &User{
		UID:     u.Uid,
		Name:    u.Name,
		Profile: u.Profile,
	}
}

func userMapDTO2VO(m map[int64]*Userinfo) map[int64]*User {
	userMap := make(map[int64]*User, len(m))
	for k, v := range m {
		userMap[k] = userDTO2VO(v)
	}
	return userMap
}

func walletDTO2VO(w *WalletInfo) *Wallet {
	return &Wallet{
		UID:           w.Uid,
		Balance:       w.Balance,
		FreezeBalance: w.FreezeBalance,
		Pay:           w.Pay,
		FreezePay:     w.FreezePay,
	}
}

func walletMapDTO2VO(m map[int64]*WalletInfo) map[int64]*Wallet {
	walletMap := make(map[int64]*Wallet, len(m))
	for k, v := range m {
		walletMap[k] = walletDTO2VO(v)
	}
	return walletMap
}
