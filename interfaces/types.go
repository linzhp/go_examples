package core

type Info struct {
	name string
}

type Methods interface {
	getInfo() Info
}
