package main

import (
	"fmt"
	"testing"
)

// run tests with; ... % cd hello;  % go test . -v

func Test_hello(t *testing.T) {
	expect := "hello world"
	if hello(nil) != expect {
		t.Errorf("failed, expect: %s", expect)
	}

	empty := ""
	if hello(&empty) != expect {
		t.Errorf("failed, expect: %s", expect)
	}

	kitty := "kitty"
	expect = fmt.Sprintf("hello %s", kitty)
	if hello(&kitty) != expect {
		t.Errorf("failed, expect: %s", expect)
	}

}
