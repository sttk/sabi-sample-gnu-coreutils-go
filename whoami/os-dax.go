package main

import (
	"os"
	"os/user"
	"strconv"
)

type osDax struct {
}

func newOsDax() osDax {
	return osDax{}
}

func (dax osDax) getEffectiveUserId() string {
	return strconv.Itoa(os.Geteuid())
}

func (dax osDax) getUsernameByUserId(uid string) string {
	user, err := user.LookupId(uid)
	if err != nil {
		return ""
	}
	return user.Username
}
