package balance

import (
	"context"
	"fmt"

	"upper.io/db.v3/lib/sqlbuilder"
)

const (
	tableBalance   = "balance"
	tableTxBalance = "tx_balance"

	trading       = 0
	traded        = 1
	tradeCanceled = 2
)

func getUIDAmountByTxID(sess sqlbuilder.Tx, txid string, status int) (int64, int64, error) {
	sql := fmt.Sprintf("SELECT `uid`, `amount` FROM %s WHERE `tx_id` = ? AND `status` = ?", tableTxBalance)
	row, err := sess.QueryRow(sql, txid, status)
	if err != nil {
		return 0, 0, err
	}
	var (
		uid    int64
		amount int64
	)
	err = row.Scan(&uid, &amount)
	if err != nil {
		return 0, 0, err
	}
	if amount <= 0 {
		return 0, 0, nil
	}
	return uid, amount, nil
}

func updateTxStatus(sess sqlbuilder.Tx, txid string, from, to int) error {
	sql := fmt.Sprintf("UPDATE %s SET `status` = ? WHERE `tx_id` = ? AND `status` = ?", tableTxBalance)
	res, err := sess.Exec(sql, to, txid, from)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return nil
	}
	return nil
}

func reduceTry(ctx context.Context, txid string, uid, amount int64) error {
	return nil
	//return sess.Tx(ctx, func(sess sqlbuilder.Tx) error {
	//	sql := fmt.Sprintf("UPDATE %s SET `balance` = `balance` - ?, `freeze_balance` = `freeze_balance` + ? WHERE `uid` = ? AND `balance` > ?", tableBalance)
	//	res, err := sess.Exec(sql, amount, amount, uid, amount)
	//	if err != nil {
	//		return err
	//	}
	//	rows, err := res.RowsAffected()
	//	if err != nil {
	//		return err
	//	}
	//	if rows == 0 {
	//		return &rpc.GrpcErr{Code: 1, Msg: "balance insufficient"}
	//	}
	//	sql = fmt.Sprintf("INSERT INTO %s (`tx_id`, `uid`, `amount`, `status`) VALUES (?, ?, ?, ?)", tableTxBalance)
	//	_, err = sess.Exec(sql, txid, uid, amount, statusUnpaid)
	//	return err
	//})
}

func reduceConfirm(ctx context.Context, txid string) error {
	return nil
	//return sess.Tx(ctx, func(sess sqlbuilder.Tx) error {
	//	uid, amount, err := getUIDAmountByTxID(sess, txid, statusUnpaid)
	//	if err != nil {
	//		return err
	//	}
	//	err = updateTxStatus(sess, txid, statusUnpaid, statusPaid)
	//	sql := fmt.Sprintf("UPDATE %s SET `freeze_balance` = `freeze_balance` - ? WHERE `uid` = ? AND `freeze_balance` >= ?", tableBalance)
	//	res, err := sess.Exec(sql, amount, uid, amount)
	//	if err != nil {
	//		return err
	//	}
	//	rows, err := res.RowsAffected()
	//	if err != nil {
	//		return err
	//	}
	//	if rows == 0 {
	//		return &rpc.GrpcErr{Code: 3, Msg: "confirmed"}
	//	}
	//	return nil
	//})
}

func reduceCancel(ctx context.Context, txid string) error {
	return nil
	//return sess.Tx(ctx, func(sess sqlbuilder.Tx) error {
	//	uid, amount, err := getUIDAmountByTxID(sess, txid, statusUnpaid)
	//	if err == db.ErrNoMoreRows {
	//		return nil
	//	}
	//	err = updateTxStatus(sess, txid, statusUnpaid, statusCancel)
	//	if err != nil {
	//		return err
	//	}
	//	sql := fmt.Sprintf("UPDATE %s SET `balance` = `balance` + ?, `freeze_balance` = `freeze_balance` - ? WHERE `uid` = ? AND `freeze_balance` >= ?", tableBalance)
	//	res, err := sess.Exec(sql, amount, amount, uid, amount)
	//	if err != nil {
	//		return err
	//	}
	//	rows, err := res.RowsAffected()
	//	if err != nil {
	//		return err
	//	}
	//	if rows == 0 {
	//		return &rpc.GrpcErr{Code: 4, Msg: "canceled"}
	//	}
	//	return nil
	//})
}
