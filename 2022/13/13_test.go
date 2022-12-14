package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompare(t *testing.T) {
	a := List{4, 4}
	b := List{4, 4}
	assert.Equal(t, 0, Compare(a, b))

	a = List{9}
	b = List{8, 7, 6}
	assert.Equal(t, -1, Compare(a, b))

	a = List{List{1}, List{2, 3, 4}}
	b = List{List{1}, 4}
	assert.Equal(t, 1, Compare(a, b))

	a = List{List{4, 4}, 4, 4}
	b = List{List{4, 4}, 4, 4, 4}
	assert.Equal(t, 1, Compare(a, b))
}
