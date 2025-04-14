package rdb

import (
	"log"
	"os"
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
		case 0x00:
			d.ReturnCursorNStepsBack(1) // to reload the flag in the next function call
			d.ReadKeyValue()
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

func (d *Decoder) ReadKeyValue() {
	switch d.ReadFlag() {
	case 0x00:
		key := d.ReadString()
		value := d.ReadString()
		d.Content.Entries[key] = value
	default:
		log.Fatal("Unexpected flag")
		os.Exit(-1)
	}
}

func (d *Decoder) ReadKeyValueWithExpiryInMilliseconds() {
	d.ReadBytes(8) // time in millieseconds
	d.ReadKeyValue()
}

func (d *Decoder) ReadKeyValueWithExpiryInSeconds() {
	d.ReadBytes(4) // time in millieseconds
	d.ReadKeyValue()
}
