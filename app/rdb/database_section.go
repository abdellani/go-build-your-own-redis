package rdb

import (
	"encoding/binary"
	"log"
	"os"
	"time"

	"github.com/abdellani/go-build-your-own-redis/app/storage"
)

func (d *Decoder) ReadDatabaseSection() {
	d.DecodeDatabaseSelector()
	d.DecodeHashTableSizeInformation()
	d.DecodeKeysValues()
}

func (d *Decoder) DecodeDatabaseSelector() {
	//database selector
	if !d.ValidateSection(0xFE) {
		return
	}
	//database number
	d.ReadLength()
}

func (d *Decoder) DecodeHashTableSizeInformation() {
	if !d.ValidateSection(0xFB) {
		return
	}
	d.Content.ResizedbInformation.DatabaseHashTableSize = d.ReadLength()
	d.Content.ResizedbInformation.ExpiryHashTableSize = d.ReadLength()
	log.Printf(
		"Resize Information: \n\t DatabaseHashTableSize: %q \n\t ExpiryHashTableSize: %q\n",
		d.Content.ResizedbInformation.DatabaseHashTableSize,
		d.Content.ResizedbInformation.ExpiryHashTableSize,
	)

}

func (d *Decoder) DecodeKeysValues() {
	for range d.Content.ResizedbInformation.DatabaseHashTableSize {
		switch d.ReadFlag() {
		case 0x00: // no expiration time
			d.ReturnCursorNStepsBack(1) // to reload the flag in the next function call
			d.ReadKeyValue(time.Time{})
		case 0xFC:
			d.ReadKeyValueWithExpiryInMilliseconds()
		case 0xFD:
			d.ReadKeyValueWithExpiryInSeconds()
		default:
			log.Fatal("Unexpected flag DecodeKeysValues ")
			os.Exit(-1)
		}
	}
}

func (d *Decoder) ReadKeyValue(t time.Time) {
	switch d.ReadFlag() {
	case 0x00:
		key := d.ReadString()
		value := d.ReadString()
		d.Content.Entries[key] = storage.Data{
			Value:          value,
			ExpirationTime: t,
		}
	default:
		log.Fatal("Unexpected flag")
		os.Exit(-1)
	}
}

func (d *Decoder) ReadKeyValueWithExpiryInMilliseconds() {
	milliseconds := binary.LittleEndian.Uint64(d.ReadBytes(8))
	t := time.UnixMilli(int64(milliseconds)).UTC()
	d.ReadKeyValue(t)
}

func (d *Decoder) ReadKeyValueWithExpiryInSeconds() {
	seconds := binary.LittleEndian.Uint32(d.ReadBytes(4)) // time in seconds
	t := time.Unix(int64(seconds), 0).UTC()
	d.ReadKeyValue(t)
}
