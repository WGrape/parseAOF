package runner

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"parseAOF/src/global"
	"parseAOF/src/writer"
)

func HandleAOFFile(splitFilePath string, parsedFilePath string) error {
	var lineNumber = 0

	f, err := os.Open(splitFilePath)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewReader(f)

	for {
		lineNumber++
		if err == io.EOF {
			global.LogDebug(fmt.Sprintf("The line:%d content=%s", lineNumber, "<<EOF<<\n"))
			break
		}

		line, _, _ := buf.ReadLine()
		content := string(line)
		plainText, _ := writer.TranslateToPlainText(lineNumber, content)
		global.LogDebug(fmt.Sprintf("The line:%d content=%s, plainText=%s", lineNumber, content, plainText))
		if plainText == global.EmptyString {
			continue
		}

		_, err = writer.AppendFile(parsedFilePath, plainText)
		if err != nil {
			return err
		}
	}

	_ = os.Remove(splitFilePath)
	return nil
}
