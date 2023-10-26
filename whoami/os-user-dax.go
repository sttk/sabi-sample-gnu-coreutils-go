package main

import (
	"os"
	"os/user"
	"strconv"

	"github.com/sttk/sabi/errs"
)

type OsUserDax struct {
}

func NewOsUserDax() OsUserDax {
	return OsUserDax{}
}

func (dax OsUserDax) GetEffectiveUserId() string {
	return strconv.Itoa(os.Geteuid())
}

func (dax OsUserDax) GetUserNameByUserId(uid string) (string, errs.Err) {
	u, e := user.LookupId(uid)
	if e != nil {
		return "", errs.New(FailToGetUserName{Uid: uid}, e)
	}
	return u.Username, errs.Ok()
}
