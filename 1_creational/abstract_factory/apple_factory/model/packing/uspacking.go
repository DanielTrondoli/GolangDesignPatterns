package packing

type USPacking struct {
	IPacking
}

func (p USPacking) Pack() string {
	return "\t- Packing in English"
}
