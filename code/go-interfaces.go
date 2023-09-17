type student struct {
	name string
	note []float64
}

type person interface {
	getName() (name string)
}

func test(p person) {
	println(p.getName())
}

func main() {
	s := student{name: "John Doe", note: []float64{10, 9.4, 8}}
    // nu compileaza fara implementarea de mai jos a metodei getName pentru structura student
    test(&s)
}

func (s *student) getName() (name string) {
	return s.name
}
