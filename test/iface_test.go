package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToAndFrom(t *testing.T) {
	abc := TestInterfaceMsg{}
	assert.NoError(t, abc.FromInterface(&A{}))
	assert.NotNil(t, abc.ToInterface())
	assert.NoError(t, abc.FromInterface(&B{}))
	assert.NotNil(t, abc.ToInterface())
	assert.NoError(t, abc.FromInterface(&C{}))
	assert.NotNil(t, abc.ToInterface())
}
