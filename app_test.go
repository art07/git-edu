package main

import (
	ast "github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	inI  int32
	inI2 int32
	outI int32
}

func TestMyFunc(t *testing.T) {
	tests := []Test{
		{inI: 0, inI2: 0, outI: 0},
		{inI: -1, inI2: 5, outI: 4},
		{inI: 1, inI2: -21, outI: -20},
	}

	assert := ast.New(t)
	for _, test := range tests {
		res := myFunc(test.inI, test.inI2)
		assert.EqualValues(test.outI, res)
	}
}