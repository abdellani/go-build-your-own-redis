package rdb

import (
	"encoding/binary"
	"log"
	"os"

	"github.com/abdellani/go-build-your-own-redis/app/storage"
)

type Decoder struct {
	Index   int
	Data    []byte
	Content RdbContent
}

type RdbContent struct {
	MagicNumber         string
	VersionNumber       string
	Auxilaries          map[string]string
	ResizedbInformation ResizedbInformation
	Entries             map[string]storage.Data
}

type ResizedbInformation struct {
	DatabaseHashTableSize int
	ExpiryHashTableSize   int
}

type Storage interface {
	Set(key string, value storage.Data)
}

func New(data []byte) *Decoder {
	return &Decoder{
		Data:  data,
		Index: 0,
		Content: RdbContent{
			Auxilaries:          map[string]string{},
			ResizedbInformation: ResizedbInformation{},
			Entries:             map[string]storage.Data{},
		},
	}
}

func (d *Decoder) Decode() *RdbContent {
	d.ReadHeaderSection()
	d.ReadAuxilarySection()
	d.ReadDatabaseSection()
	d.ReadEndOfFileSection()
	return &d.Content
}

func (d *Decoder) WriteEntries(storage Storage) {
	for key, value := range d.Content.Entries {
		storage.Set(key, value)
	}
}

func ConvertBytesToInt(bytes []byte) int {
	switch len(bytes) {
	case 1:
		return int(bytes[0])
	case 2:
		return int(binary.BigEndian.Uint16(bytes))
	case 4:
		return int(binary.BigEndian.Uint32(bytes))
	default:
		log.Fatal("Unexpected case!")
		os.Exit(1)
		return 0
	}
}

func (d *Decoder) ValidateSection(sectionCode byte) bool {
	flag := d.ReadFlag()
	if flag == sectionCode {
		return true
	}
	d.ReturnCursorNStepsBack(1)
	return false
}
