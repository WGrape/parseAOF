package writer

import (
	"fmt"
	"parseAOF/src/global"
	"testing"
)

func TestTranslateToPlainText(t *testing.T) {
	var lineNumber int
	var expectContent string
	var realContent string

	realContent, _ = TranslateToPlainText(lineNumber, global.EmptyString)
	if realContent != global.WhitespaceString {
		msg := fmt.Sprintf("[TestWriteFile 01] Test failed: %s != %s\n", realContent, expectContent)
		t.Error(msg)
		return
	}

	realContent, _ = TranslateToPlainText(lineNumber, "$3")
	if realContent != global.EmptyString {
		msg := fmt.Sprintf("[TestWriteFile 02] Test failed: %s != %s\n", realContent, expectContent)
		t.Error(msg)
		return
	}

	fmt.Printf("[TestWriteFile] Test success\n")
}
