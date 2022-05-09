package personanimation

type PersonSample struct {
	Sample Person
}

func (p *PersonSample) LoadStickMan() {
	p.Sample.L1 = "                   "
	p.Sample.L2 = "        o          "
	p.Sample.L3 = "       /|\\         "
	p.Sample.L4 = "_______/_\\_________"
}

func (p *PersonSample) LoadFatMan() {
	p.Sample.L1 = "                      "
	p.Sample.L2 = "       (^ ^)          "
	p.Sample.L3 = "       <)  )\\         "
	p.Sample.L4 = "_______/__\\___________"
}
