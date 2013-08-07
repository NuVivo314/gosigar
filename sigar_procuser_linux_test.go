package sigar

import (
  "errors"
  "fmt"
  "os"
  "testing"
)

func TestProcUser(t *testing.T) {
  userInfo := ProcUser{}
  err := userInfo.Get(os.Getpid())
  if err != nil {
    t.Error(err)
  }

  if userInfo.UidReal != fmt.Sprint(os.Getuid()) {
    t.Error(errors.New("Real UID != os.Getuid{}"))
  }

  if userInfo.GidReal != fmt.Sprint(os.Getgid()) {
    t.Error(errors.New("Real GID != os.Getgid{}"))
  }

  err = userInfo.Get(invalidPid)
  if err == nil {
    t.Errorf("Invalid ProcUser.Get('%d')", invalidPid)
  }
}
