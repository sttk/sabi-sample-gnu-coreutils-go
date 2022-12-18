package main

import (
	"os"
	"os/user"
	"strconv"
)

type OsDax struct {
}

func NewOsDax() OsDax {
	return OsDax{}
}

func (dax OsDax) GetEffectiveUserId() string {
	return strconv.Itoa(os.Geteuid())
}

func (dax OsDax) GetUserNameByUserId(uid string) string {
	user, err := user.LookupId(uid)
	if err != nil {
		return ""
	}
	return user.Username
}
