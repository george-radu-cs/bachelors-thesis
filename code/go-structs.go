type student struct {
	name string
	note []float64
}

func (s *student) computeAverageNote() float64 {
	var sum float64
	for _, n := range s.note {
		sum += n
	}
	return sum / float64(len(s.note))
}