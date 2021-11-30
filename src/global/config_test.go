package global

import (
	"fmt"
	"testing"
)

func TestGetDebug(t *testing.T) {
	var config AppConfig
	config.SetDebug(true)
	realDebug := config.GetDebug()
	if realDebug != true {
		msg := fmt.Sprintf("[TestSetDebug 01] Test failed: %v != %v\n", realDebug, true)
		t.Error(msg)
		return
	}
	fmt.Printf("[TestSetDebug] Test success\n")
}
