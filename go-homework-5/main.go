package main

type FactorialsCalculatorInterface interface {
	Calc(max int) map[int]int
}

type FactorialsCalculator struct{}

func NewCalculator() FactorialsCalculatorInterface {
	return &FactorialsCalculator{}
}

func (f *FactorialsCalculator) Calc(max int) map[int]int {
	factorials := map[int]int{}
	for i := 1; i <= max; i++ {
		// TODO: реализовать параллельный подсчет факториала для кажого i в разных горутинах
	}
	return factorials
}

func main() {
}
