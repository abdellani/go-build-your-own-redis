package rdb

import (
	"encoding/binary"
	"log"
	"os"
	"strconv"
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
	Entries             map[string]string
}

type ResizedbInformation struct {
	DatabaseHashTableSize int
	ExpiryHashTableSize   int
}

func New(data []byte) *Decoder {
	return &Decoder{
		Data:  data,
		Index: 0,
		Content: RdbContent{
			Auxilaries:          map[string]string{},
			ResizedbInformation: ResizedbInformation{},
			Entries:             map[string]string{},
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

func (d *Decoder) ReadBytes(n int) []byte {
	buffer := d.Data[d.Index : d.Index+n]
	d.Index += n
	return buffer
}

func (d *Decoder) ReachedEOF() bool {
	return d.Index >= len(d.Data)
}

func (d *Decoder) ReadFlag() byte {
	flag := d.Data[d.Index]
	d.Index++
	return flag
}

func (d *Decoder) ReadString() string {
	b := d.ReadBytes(1)[0]
	switch b & 0b11000000 {
	case 0b0000_0000:
		// the next 6 bits represent the length
		length := b & 0b0011_1111
		return string(d.ReadBytes(int(length)))
	case 0b0100_0000:
		log.Fatal("Decoder->ReadString 0b01 not processed")
	case 0b1000_0000:
		log.Fatal("Decoder->ReadString 0b10 not processed")
	case 0b1100_0000:
		length := int(b & 0b0011_1111)
		/*
			0 -> 8 bits follows
			1 -> 16 bits follows
			2 -> 32 bits follows
		*/
		integer := ConvertBytesToInt(d.ReadBytes(length + 1))
		return strconv.Itoa(integer)
	default:
		log.Fatal("Unexpected case")
	}
	os.Exit(-1)
	return ""
}

func (d *Decoder) ReadLength() int {
	b := d.ReadBytes(1)[0]
	switch b & 0b11000000 {
	case 0b0000_0000:
		// the next 6 bits represent the length
		length := b & 0b0011_1111
		return int(length)
	case 0b0100_0000:
		log.Fatal("Decoder->ReadLength 0b01 not processed")
	case 0b1000_0000:
		log.Fatal("Decoder->ReadLength 0b10 not processed")
	case 0b1100_0000:
		log.Fatal("Decoder->ReadLength 0b11  not processed ")
	default:
		log.Fatal("Unexpected case")
	}
	os.Exit(-1)
	return 0
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
