package pojo

type Plan struct {
	Rate    float64
	Ammount float64
}

func (p *Plan) SetRate(rate float64) {
	p.Rate = rate
}

func (p *Plan) GetRate() float64 {
	return p.Rate
}

func (p *Plan) SetAmmount(ammount float64) {
	p.Ammount = ammount
}

func (p *Plan) GetAmmount() float64 {
	return p.Ammount
}
