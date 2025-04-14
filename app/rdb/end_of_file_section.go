package rdb

import "log"

func (d *Decoder) ReadEndOfFileSection() {
	if d.ReadFlag() == 0xff {
		log.Println("EOF read successfully")
	}
}
