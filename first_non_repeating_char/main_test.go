package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstNonRepeating(t *testing.T) {
	for c, r := range map[string]string{
		"stress": "t",
		"rsssry": "y",
		"sTreSS": "T",
	} {
		assert.Equal(t, r, FirstNonRepeating(c))
	}
}
