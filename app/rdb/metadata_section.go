package rdb

import "log"

func (d *Decoder) ReadAuxilarySection() {
	// Read auxilary
	flag := d.ReadBytes(1)[0]
	for {
		switch flag {
		case 0xFA:
			d.DecodeAuxilary()
			flag = d.ReadBytes(1)[0]
		default:
			// next read will reload the flag
			d.ReturnCursorNStepsBack(1)
			return
		}
	}

}
func (d *Decoder) ReturnCursorNStepsBack(n int) {
	d.Index -= n
}
func (d *Decoder) DecodeAuxilary() {
	key := d.ReadString()
	value := d.ReadString()
	d.Content.Auxilaries[key] = value
	log.Printf("Auxilary: key: %q value: %q", key, value)
}
