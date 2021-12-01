package main

import (
	"fmt"
	"io/ioutil"
	"parseAOF/src/global"
	"parseAOF/src/runner"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	files, _ := ioutil.ReadDir(global.DataDir)
	for _, file := range files {
		fileName := file.Name()
		if global.IsAOFFile(fileName) {
			wg.Add(global.Delta)
			go func() {
				defer wg.Done()
				splitFilePath := global.GetSplitFilePath(fileName)
				parsedFilePath := global.GetParsedFilePath(fileName)
				err := runner.HandleAOFFile(splitFilePath, parsedFilePath)
				if err != nil {
					global.LogError(fmt.Sprintf("HandleAOFFile failed: %s | message: %s\n", splitFilePath, err.Error()))
				} else {
					global.LogInfo(fmt.Sprintf("HandleAOFFile success: %s => %s\n", fileName, splitFilePath))
				}
			}()
		}
	}
	wg.Wait()
}
