package calculator

type Calculator struct {
	Bilangan []float64
	Operator string
}

func (c *Calculator) Proses() (hasil float64) {
	if c.Operator == "kali" {
		hasil = c.Dikali()
	} else if c.Operator == "tambah" {
		hasil = c.Ditambah()
	}
	return
}
func (c *Calculator) Dikali() float64 {
	return c.Bilangan[0] * c.Bilangan[1]
}
func (c *Calculator) Ditambah() float64 {
	return c.Bilangan[0] + c.Bilangan[1]
}
