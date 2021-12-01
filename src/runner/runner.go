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
	var lineByte []byte

	f, err := os.Open(splitFilePath)
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewReader(f)

	for {
		lineNumber++
		lineByte, _, err = buf.ReadLine()
		if err == io.EOF {
			global.LogDebug(fmt.Sprintf("Line:%d |content=%s|\n", lineNumber, "<<EOF<<"))
			break
		}

		content := string(lineByte)
		plainText, _ := writer.TranslateToPlainText(lineNumber, content)
		global.LogDebug(fmt.Sprintf("Line:%d |content=%s|plainText=%s|\n", lineNumber, content, plainText))
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
