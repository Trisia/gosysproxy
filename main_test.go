// +build windows

package gosysproxy

import "testing"

func TestFlush(t *testing.T) {
	err := Flush()
	if err != nil {
		t.Fatal(err)
	}
}

func TestOff(t *testing.T) {
	err := Off()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetPAC(t *testing.T) {
	err := SetPAC("http://127.0.0.1:7777/pac")
	if err != nil {
		t.Fatal(err)
	}
}

func TestSetGlobalProxy(t *testing.T) {
	err := SetGlobalProxy("127.0.0.1:7890")
	if err != nil {
		t.Fatal(err)
	}

	err = Off()
	if err != nil {
		t.Fatal(err)
	}

	err = SetGlobalProxy("127.0.0.1:7890", "foo", "bar")
	if err != nil {
		t.Fatal(err)
	}

	err = Off()
	if err != nil {
		t.Fatal(err)
	}
}
