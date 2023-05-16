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

func (m *UsersModel) Login(phoneNumber, password string) (int, error) {
	var id int
	var hashedPassword string

	row := m.DB.QueryRow("SELECT id, password FROM users WHERE phone_number = ?", phoneNumber)
	err := row.Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errors.New("phone number not found")
		}
		return 0, err
	}

	if password != hashedPassword {
		return 0, errors.New("incorrect password")
	}

	return id, nil
}

func (m *UsersModel) GetUserByPhoneNumber(phoneNumber string) (*Users, error) {
	user := &Users{}

	row := m.DB.QueryRow("SELECT id, phone_number, name, gender, tanggal_lahir, balance FROM users WHERE phone_number = ?", phoneNumber)
	err := row.Scan(&user.ID, &user.PhoneNumber, &user.Name, &user.Gender, &user.TanggalLahir, &user.Balance)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("phone number not found")
		}
		return nil, err
	}

	return user, nil
}

func (m *UsersModel) UpdateAccount(userID int, name, gender, tanggalLahir string) error {
	_, err := m.DB.Exec("UPDATE users SET name=?, gender=?, tanggal_lahir=? WHERE id=?", name, gender, tanggalLahir, userID)
	if err != nil {
		return err
	}

	return nil
}
