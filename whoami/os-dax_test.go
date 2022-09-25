package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOsDax_getEffectiveUserId(t *testing.T) {
	dax := newOsDax()
	assert.NotEqual(t, dax.getEffectiveUserId(), "")
}

func TestOsDax_getUsernameByUserId(t *testing.T) {
	dax := newOsDax()
	euid := dax.getEffectiveUserId()
	assert.NotEqual(t, dax.getUsernameByUserId(euid), "")
	assert.Equal(t, dax.getUsernameByUserId("xxx"), "")
}
