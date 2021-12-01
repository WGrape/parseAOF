package parser

import (
	"fmt"
	"parseAOF/src/global"
	"testing"
)

func TestMatchLine(t *testing.T) {
	realMatchType := MatchLine(global.EmptyString)
	if realMatchType != global.MatchTypeArgRaw {
		msg := fmt.Sprintf("[TestMatchLine 01] Test failed: %d != %d\n", realMatchType, global.MatchTypeArgRaw)
		t.Error(msg)
		return
	}
	fmt.Printf("[TestMatchLine] Test success\n")
}
