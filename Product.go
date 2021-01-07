package ken_sample_golang

type Product struct {
	Length float64
	Width  float64
	Height float64
	Weight float64
}

func NewProduct(length, width, height, weight float64) *Product {
	return &Product{
		Length: length,
		Width:  width,
		Height: height,
		Weight: weight,
	}
}

func (p *Product) Size() float64 {
	return p.Height * p.Width * p.Length
}
