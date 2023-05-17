package topup

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type TopUpModel struct {
	DB *sqlx.DB
}
type TopupModel struct {
	DB *sql.DB
}

func (tm *TopupModel) TopUpAccount(phoneNumber string, amount int) error {
	// Get the user ID based on the phone number
	var userID int
	err := tm.DB.QueryRow("SELECT id FROM users WHERE phone_number = ?", phoneNumber).Scan(&userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("user not found")
		}
		return err
	}

	// Insert top-up record
	_, err = tm.DB.Exec("INSERT INTO topup (user_id, amount, tanggal) VALUES (?, ?, NOW())", userID, amount)
	if err != nil {
		return err
	}

	// Update user balance
	_, err = tm.DB.Exec("UPDATE users SET balance = balance + ? WHERE id = ?", amount, userID)
	if err != nil {
		return err
	}

	return nil
}

func (tm *TopupModel) GetTopUpHistory(userID int) ([]Topup, error) {
	rows, err := tm.DB.Query("SELECT id, user_id, amount, tanggal FROM topup WHERE user_id = ?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	history := []Topup{}
	for rows.Next() {
		var t Topup
		var tanggalString string
		err := rows.Scan(&t.ID, &t.UserID, &t.Amount, &tanggalString)
		if err != nil {
			return nil, err
		}
		t.Tanggal, err = time.Parse("2006-01-02 15:04:05", tanggalString)
		if err != nil {
			return nil, err
		}
		history = append(history, t)
	}

	return history, nil
}
