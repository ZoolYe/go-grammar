package main

type Base struct {
	Name string
}

type Foo struct {
	Base
}

func (base *Base) Fool() {

}

func (base *Base) Bar() {

}

func (foo *Foo) Bar() {

}
