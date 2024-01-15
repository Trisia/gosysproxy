//go:build windows
// +build windows

package gosysproxy

import (
	"fmt"
	"testing"
)

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

func TestStatus(t *testing.T) {
	err := SetGlobalProxy("127.0.0.1:7890")
	if err != nil {
		t.Fatal(err)
	}

	err = Off()
	if err != nil {
		t.Fatal(err)
	}

	err = SetGlobalProxy("127.0.0.1:7890", "foo", "bar", "<local>")
	if err != nil {
		t.Fatal(err)
	}
	res, err := Status()
	if err != nil {
		t.Fatal(err)
	}
	switch res.Type {
	case INTERNET_OPEN_TYPE_PRECONFIG:
		fmt.Println(">> 代理方式：", "use registry configuration")
	case INTERNET_OPEN_TYPE_DIRECT:
		fmt.Println(">> 代理方式：", "不代理 direct to net")
	case INTERNET_OPEN_TYPE_PROXY:
		fmt.Println(">> 代理方式：", "使用代理服务器 via named proxy")
	}
	fmt.Println(">> 代理地址：", res.Proxy)
	fmt.Println(">> 请勿对以下列条目开头的地址使用代理服务器：", res.Bypass)
	fmt.Println(">> 请勿将代理服务器用于本地(Intranet)地址：", res.DisableProxyIntranet)

	err = Off()
	if err != nil {
		t.Fatal(err)
	}

}
