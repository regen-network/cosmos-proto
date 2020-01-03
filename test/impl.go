package test

import "fmt"

func (m A) SomeMethod() string {
	return fmt.Sprintf("A.X:%d", m.X)
}

func (m B) SomeMethod() string {
	return fmt.Sprintf("B.Y:%d", m.Y)
}

func (m C) SomeMethod() string {
	return fmt.Sprintf("C.Z:%t", m.Z)
}
