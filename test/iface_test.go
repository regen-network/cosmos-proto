package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type D struct{}

func (m D) SomeMethod() string { return "D" }

func TestToAndFrom(t *testing.T) {
	abc := ABC{}

	assert.NoError(t, abc.SetMsg(&A{X: 1}))
	iface1 := abc.GetMsg()
	assert.NotNil(t, iface1)
	assert.Equal(t, "A.X:1", iface1.SomeMethod())

	assert.NoError(t, abc.SetMsg(&B{Y: 2}))
	iface1 = abc.GetMsg()
	assert.NotNil(t, iface1)
	assert.Equal(t, "B.Y:2", iface1.SomeMethod())

	assert.NoError(t, abc.SetMsg(&C{Z: true}))
	iface1 = abc.GetMsg()
	assert.NotNil(t, iface1)
	assert.Equal(t, "C.Z:true", iface1.SomeMethod())

	// Non-pointer cases
	assert.NoError(t, abc.SetMsg(A{X: 1}))
	iface1 = abc.GetMsg()
	assert.NotNil(t, iface1)
	assert.Equal(t, "A.X:1", iface1.SomeMethod())

	assert.NoError(t, abc.SetMsg(B{Y: 2}))
	iface1 = abc.GetMsg()
	assert.NotNil(t, iface1)
	assert.Equal(t, "B.Y:2", iface1.SomeMethod())

	assert.NoError(t, abc.SetMsg(C{Z: true}))
	iface1 = abc.GetMsg()
	assert.NotNil(t, iface1)
	assert.Equal(t, "C.Z:true", iface1.SomeMethod())

	// test nil
	assert.NoError(t, abc.SetMsg(nil))
	iface1 = abc.GetMsg()
	assert.Nil(t, iface1)

	// D implements Msg but isn't in the ABC oneof and so this should fail
	assert.Error(t, abc.SetMsg(&D{}))
}

func TestToAndFromNonPointer(t *testing.T) {
	abc := ABCNonPointer{}

	assert.NoError(t, abc.SetMsg(&A{X: 1}))
	iface1 := abc.GetMsg()
	assert.NotNil(t, iface1)
	assert.Equal(t, "A.X:1", iface1.SomeMethod())

	assert.NoError(t, abc.SetMsg(&B{Y: 2}))
	iface1 = abc.GetMsg()
	assert.NotNil(t, iface1)
	assert.Equal(t, "B.Y:2", iface1.SomeMethod())

	assert.NoError(t, abc.SetMsg(&C{Z: true}))
	iface1 = abc.GetMsg()
	assert.NotNil(t, iface1)
	assert.Equal(t, "C.Z:true", iface1.SomeMethod())

	// Non-pointer cases should error here
	assert.Error(t, abc.SetMsg(A{X: 1}))
	assert.Error(t, abc.SetMsg(B{Y: 2}))
	assert.Error(t, abc.SetMsg(C{Z: true}))

	// test nil
	assert.NoError(t, abc.SetMsg(nil))
	iface1 = abc.GetMsg()
	assert.Nil(t, iface1)

	// D implements Msg but isn't in the ABC oneof and so this should fail
	assert.Error(t, abc.SetMsg(&D{}))
}
