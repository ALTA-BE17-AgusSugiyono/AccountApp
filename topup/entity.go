package topup

import "time"

type Topup struct {
	ID      int
	UserID  int
	Amount  int
	Tanggal time.Time
}
