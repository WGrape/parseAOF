package main

import (
	"fmt"
	"io/ioutil"
	"parseAOF/src/global"
	"parseAOF/src/runner"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(global.Config.MaxRoutines)

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
					global.LogError(fmt.Sprintf("HandleAOFFile failed: %s | error: %s\n", splitFilePath, err.Error()))
				} else {
					global.LogInfo(fmt.Sprintf("HandleAOFFile success: %s => %s => %s\n", fileName, splitFilePath, parsedFilePath))
				}
			}()
			time.Sleep(500 * time.Millisecond)
		}
	}
	wg.Wait()
}
