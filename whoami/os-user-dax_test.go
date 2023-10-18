package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhoami_OsUserDax_GetEffectiveUserId(t *testing.T) {
	dax := NewOsUserDax()
	assert.NotEqual(t, dax.GetEffectiveUserId(), "")
}

func TestWhoami_OsUserDax_GetUserNameByUserId(t *testing.T) {
	unm := os.Getenv("USER")

	dax := NewOsUserDax()
	euid := dax.GetEffectiveUserId()
	u, err := dax.GetUserNameByUserId(euid)
	assert.True(t, err.IsOk())
	assert.Equal(t, u, unm)
}

func TestWhoami_OsUserDax_GetUserNameByUserId_error(t *testing.T) {
	euid := "xxx"

	dax := NewOsUserDax()
	_, err := dax.GetUserNameByUserId(euid)
	assert.True(t, err.IsNotOk())
	switch r := err.Reason().(type) {
	case FailToGetUserName:
		assert.Equal(t, r.Uid, euid)
		assert.Equal(t, err.Cause().Error(), `strconv.Atoi: parsing "xxx": invalid syntax`)
	default:
		assert.Fail(t, err.Error())
	}
}
