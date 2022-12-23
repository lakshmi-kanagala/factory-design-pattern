package plan

type Plan interface {
	SetRate(rate float64)
	GetAmmount() float64
	GetRate() float64
	SetAmmount(ammount float64)
}
