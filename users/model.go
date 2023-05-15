package users

import (
	"database/sql"
	"errors"
)

type UsersModel struct {
	DB *sql.DB
}

func (m *UsersModel) Register(phoneNumber, password, name, gender, tanggalLahir string) error {
	// cek apakah nomor telepon sudah digunakan
	var count int
	err := m.DB.QueryRow("SELECT COUNT(*) FROM users WHERE phone_number = ?", phoneNumber).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("Nomor telepon sudah digunakan")
	}

	// insert user baru ke database
	_, err = m.DB.Exec("INSERT INTO users(phone_number, password, name, gender, tanggal_lahir, balance) VALUES(?,?,?,?,?,?)", phoneNumber, password, name, gender, tanggalLahir, 0)
	if err != nil {
		return err
	}

	return nil
}
