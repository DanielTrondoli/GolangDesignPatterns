package dragonball

type Shenlong struct{}

func (s Shenlong) Resurect(p Person) (Person, error) {
	return p.Resurect()
}
