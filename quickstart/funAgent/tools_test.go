package main

import (
	"context"
	"testing"
)

func TestOpenBrowser(t *testing.T) {
	param := &BrowserParam{
		Url: "https://www.baidu.com",
	}
	msg, err := OpenBrowser(context.Background(), param)
	if err != nil {
		t.Errorf("OpenBrowser failed, err=%v", err)
	}
	t.Logf("OpenBrowser success, msg=%s", msg)
}
