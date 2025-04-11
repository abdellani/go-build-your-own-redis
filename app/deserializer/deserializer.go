package deserializer

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type Deserializer struct {
	Payload string
	Index   int
}

type Command struct {
	Command   string
	Arguments []string
}

func New(payload string) *Deserializer {
	return &Deserializer{Payload: payload}
}

func (d *Deserializer) Deserialize() *Command {
	symbol, _ := d.ReadSymbol()
	if IsArray(symbol) == true {
		return d.ReadArray()
	} else {
		os.Exit(-1)
		return nil
	}
}

func (d *Deserializer) ReadArray() *Command {
	size, _ := d.ReadInt()
	d.ReadSeperator()
	command := &Command{}
	symbol, _ := d.ReadSymbol()

	if IsString(symbol) == true {
		command.Command = d.ReadString()
	}

	for i := 1; i < size; i++ {
		symbol, _ := d.ReadSymbol()
		if IsString(symbol) == true {
			command.Arguments = append(command.Arguments,
				d.ReadString(),
			)
		}
	}
	return command
}

func (d *Deserializer) ReadString() string {
	stringLength, _ := d.ReadInt()
	d.ReadSeperator()
	s := d.ReadStringOfSize(stringLength)
	d.ReadSeperator()
	return s
}

func (d *Deserializer) ReadSymbol() (string, error) {
	if d.Index >= len(d.Payload) {
		return "", errors.New("reached the end")
	}
	symbol := d.Payload[d.Index : d.Index+1]
	d.Index++
	return symbol, nil
}

func (d *Deserializer) ReadInt() (int, error) {
	i := d.Index
	for i < len(d.Payload) && d.Payload[i] != '\r' {
		i++
	}
	stringifiedInteger := d.Payload[d.Index:i]

	d.Index = i
	return strconv.Atoi(stringifiedInteger)
}

func (d *Deserializer) ReadSeperator() error {
	if strings.Compare(d.Payload[d.Index:d.Index+2], "\r\n") != 0 {
		return errors.New("not a separator")
	}
	d.Index += 2
	return nil
}

func (d *Deserializer) ReadStringOfSize(length int) string {
	str := d.Payload[d.Index : d.Index+length]
	d.Index += length
	return str
}
