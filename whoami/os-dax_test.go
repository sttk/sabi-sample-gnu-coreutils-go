package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOsDax_GetEffectiveUserId(t *testing.T) {
	dax := NewOsDax()
	assert.NotEqual(t, dax.GetEffectiveUserId(), "")
}

func TestOsDax_GetUserNameByUserId(t *testing.T) {
	dax := NewOsDax()
	euid := dax.GetEffectiveUserId()
	assert.NotEqual(t, dax.GetUserNameByUserId(euid), "")
	assert.Equal(t, dax.GetUserNameByUserId("xxx"), "")
}
