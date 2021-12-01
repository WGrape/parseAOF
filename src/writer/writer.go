package writer

import (
	"fmt"
	"os"
	"parseAOF/src/global"
	"parseAOF/src/parser"
)

func TranslateToPlainText(lineNumber int, content string) (string, error) {
	var plainText string
	matchType := parser.MatchLine(content)
	if matchType == global.MatchTypeCmdStart && lineNumber > 1 {
		plainText = fmt.Sprintf("\n")
	} else if matchType == global.MatchTypeArgRaw {
		plainText = fmt.Sprintf("%s ", content)
	} else {
		plainText = global.EmptyString
	}
	return plainText, nil
}

func AppendFile(filePath string, content string) (string, error) {
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, global.DefaultFileMode)
	if err != nil {
		return global.EmptyString, err
	}

	_, err = f.WriteString(content)
	if err != nil {
		return global.EmptyString, err
	}
	return content, nil
}
