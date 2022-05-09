package packing

type BrazilianPacking struct {
	IPacking
}

func (p BrazilianPacking) Pack() string {
	return "\t- Empacotando em Portugues"
}
