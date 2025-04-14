package rdb

func (d *Decoder) ReadHeaderSection() {
	//Read Header
	d.Content.MagicNumber = string(d.ReadBytes(5))
	d.Content.VersionNumber = string(d.ReadBytes(4))

}
