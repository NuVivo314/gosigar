// +build linux

package sigar


import (
  "os"
  "testing"
)


func TestDiskIO(t *testing.T) {
  diskio := DiskIO{}

  err := diskio.Get("sda")
  if err != nil {
    t.Error("Fail to exec DiskIO test:", err.Error())
    t.Fail()
  }
  if diskio.Name != "sda" {
    t.Error("Invalid DiskIO: sda")
  }

  err = diskio.Get("FooBarITon")
  if err == nil {
    t.Error("Invalid DiskIO.Get()")
  }
}

func TestDiskIOList(t *testing.T) {
  diskiolst := DiskIOList{}

  err := diskiolst.Get()
  if err != nil {
    t.Error("Fail to get list DiskIO test:", err.Error())
  }

  for d := range diskiolst.List {
    if diskiolst.List[d].Name == "" {
      t.Error("Fail to making list")
    }
  }
}

func TestProcIO(t *testing.T) {
  procio := ProcIO{}

  err := procio.Get(os.Getpid())
  if err != nil {
    t.Error("Fail to get own pid")
  }

  if os.Getuid() > 0 {
    err = procio.Get(1)
    if err == nil {
      t.Error("Bad detecting of reading fail")
    }
  }
}
