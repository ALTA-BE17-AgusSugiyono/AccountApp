package transfer

import "time"

type Transfer struct {
	ID          int
	Sender_ID   int
	Receiver_ID int
	Amount      int
	Tanggal     time.Time
}
