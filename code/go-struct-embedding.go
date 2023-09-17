type base struct {
	name string
}

func (b *base) baseFunc() string {
	return b.name
}

type container struct {
	base
	otherField string
}

func main() {
	c := container{
		base: base{
			name: "John Doe",
		},
		otherField: "other field",
	}
	fmt.Printf("c base name: %s\n", c.base.name)
	fmt.Printf("c call to baseFunc: %s\n", c.baseFunc())
}
