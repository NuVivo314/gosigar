package sigar

import (
	"testing"
)

func TestNetworkList(t *testing.T) {
	netlst := NetworkList{}

	err := netlst.Get()
	if err != nil {
		t.Error(err)
	}

	net := Network{}
	err = net.Get("lo")
	if err != nil || net.Name != "lo" {
		t.Errorf("Error to find loopback pseudo-device('lo'):", err)
	}

	netempty := Network{}
	err = netempty.Get("12341234123qwertyui")
	if err == nil {
		t.Error("Fault error, name of interface can't exist", netempty)
	}

}
