package main

type VaccineInterface interface {
	DiseaseProbability() float64 // вероятность заболеть, от 0 до 1
}
type Vaccine struct {
	probability float64 // вероятность заболеть, от 0 до 1
}

func (v *Vaccine) DiseaseProbability() float64 {
	return v.probability
}
func NewSputnikV() *Vaccine {
	return &Vaccine{probability: 0.03}
}
func NewPhizer() *Vaccine {
	return &Vaccine{probability: 0.04}
}

type Human struct {
	appliedVaccine VaccineInterface // какую вакцину поставили
}

func (h *Human) ApplyVaccine(v VaccineInterface) {
	h.appliedVaccine = v
}

func main() {
	sputnikV := NewSputnikV()
	phizer := NewPhizer()

	Ivan := Human{}
	John := Human{}

	Ivan.ApplyVaccine(sputnikV)
	John.ApplyVaccine(phizer)
}
