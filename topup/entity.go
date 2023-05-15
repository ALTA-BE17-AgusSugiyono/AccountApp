package topup

import "time"

type Topup struct {
	ID      int
	User_ID int
	Amount  int
	Tanggal time.Time
}
