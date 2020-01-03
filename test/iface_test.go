package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToAndFrom(t *testing.T) {
	abc := ABC{}

	assert.NoError(t, abc.SetInterface1(&A{X: 1}))
	iface1 := abc.GetInterface1()
	assert.NotNil(t, iface1)
	assert.Equal(t, "A.X:1", iface1.SomeMethod())

	assert.NoError(t, abc.SetInterface1(&B{Y: 2}))
	iface1 = abc.GetInterface1()
	assert.NotNil(t, iface1)
	assert.Equal(t, "B.Y:2", iface1.SomeMethod())

	assert.NoError(t, abc.SetInterface1(&C{Z: true}))
	iface1 = abc.GetInterface1()
	assert.NotNil(t, iface1)
	assert.Equal(t, "C.Z:true", iface1.SomeMethod())
}
