package runner

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"parseAOF/src/global"
	"parseAOF/src/writer"
	"time"
)

func HandleAOFFile(splitFilePath string, parsedFilePath string) error {
	var splitFileResource *os.File
	var parsedFileResource *os.File
	var lineNumber = 0
	var lineByte []byte
	var err error

	splitFileResource, err = os.Open(splitFilePath)
	if err != nil {
		return err
	}
	buf := bufio.NewReader(splitFileResource)
	parsedFileResource, err = os.OpenFile(parsedFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, global.DefaultFileMode)
	if err != nil {
		return err
	}
	defer splitFileResource.Close()
	defer parsedFileResource.Close()

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

		_, err = writer.AppendFile(parsedFileResource, plainText)
		if err != nil {
			return err
		}
		if lineNumber%100 == 0 {
			time.Sleep(10 * time.Millisecond)
		}
	}

	_ = os.Remove(splitFilePath)
	return nil
}
