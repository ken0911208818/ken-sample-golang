package ken_sample_golang

func Post(p *Product) (float64, error) {
	feeByWeight := 80 + p.Weight*10
	size := p.size()
	feeBySize := size * 0.00002 * 1100
	if feeByWeight < feeBySize {
		return feeByWeight, nil
	}
	return feeBySize, nil
}

func FedEx(p *Product) (float64, error) {
	size := p.size()
	if p.Length > 100 || p.Width > 100 || p.Height > 100 {
		return size*0.00002*1100 + 500, nil
	}
	return size * 0.00002 * 1200, nil
}

func UPS(p *Product) (float64, error) {
	if p.Weight > 20 {
		return 500, nil
	}
	return 100 + p.Weight*10, nil
}
