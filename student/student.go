package student

type Address struct {
	street string
	city   string
}

type Student struct {
	matnumber int
	name      string
	home      Address
}
