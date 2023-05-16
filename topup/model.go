package topup

import (
	"database/sql"
)

type TopupModel struct {
	DB *sql.DB
}

func (tm *TopupModel) TopUpAccount(userID int, amount int) error {
	_, err := tm.DB.Exec("INSERT INTO topup (user_id, amount, tanggal) VALUES (?, ?, NOW())",
		userID, amount)
	if err != nil {
		return err
	}

	_, err = tm.DB.Exec("UPDATE users SET balance = balance + ? WHERE id = ?", amount, userID)
	if err != nil {
		return err
	}

	return nil
}
