package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"parseAOF/src/global"
	"parseAOF/src/runner"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

var (
	input         = ""
	output        = ""
	maxGoRountine = 8
	debug         = false
	rootCmd       = &cobra.Command{
		Use:   "parseAOF",
		Short: "parse redis aof to readable",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			global.Debug = debug
			fmt.Println(debug)
			Execute(input, output)
		},
	}
)

func init() {
	cobra.OnInitialize()
	rootCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "input AOF file path")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "output dir path")
	rootCmd.PersistentFlags().IntVarP(&maxGoRountine, "routines", "r", 8, "max goroutines")
	rootCmd.PersistentFlags().BoolVar(&debug, "v", false, "verbose mode for debug")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Execute(input string, output string) {
	if input == "" {
		return
	}
	if output != "" {
		global.DataDir = output
	} else {
		// 获取当前可执行文件路径
		execPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			global.LogError(fmt.Sprintf("get current path error: %s\n", err.Error()))
			return
		} else {
			global.DataDir = filepath.Join(execPath, "parseAOF_data")
		}
	}
	splitAOF(input)

	taskChan := make(chan string)
	files, _ := ioutil.ReadDir(global.DataDir)
	num := maxGoRountine
	if len(files) < num {
		num = len(files)
	}
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for fileName := range taskChan {
				splitFilePath := global.GetSplitFilePath(fileName)
				parsedFilePath := global.GetParsedFilePath(fileName)
				err := runner.HandleAOFFile(splitFilePath, parsedFilePath)
				if err != nil {
					global.LogError(fmt.Sprintf("HandleAOFFile failed: %s | error: %s\n", splitFilePath, err.Error()))
				} else {
					global.LogInfo(fmt.Sprintf("HandleAOFFile success: %s => %s => %s\n", fileName, splitFilePath, parsedFilePath))
				}
			}
		}()
	}
	for _, item := range files {
		if item.IsDir() {
			continue
		}
		if global.IsAOFFile(item.Name()) {
			taskChan <- item.Name()
		}
	}

	close(taskChan)
	wg.Wait()
	mergeParsedFile()
}

func splitAOF(name string) {
	global.LogInfo(fmt.Sprintf("splitAOF Start split file ...\n"))
	outputFilePrefix := filepath.Join(global.DataDir, "aof.split_")
	startNum := 10000
	// 读取文件
	fileContent, err := ioutil.ReadFile(name)
	if err != nil {
		global.LogError(fmt.Sprintf("splitAOF read src file: %s error: %s\n", name, err.Error()))
		return
	}
	_, err = os.Stat(global.DataDir)
	if os.IsNotExist(err) {
		// create data directory
		err = os.MkdirAll(global.DataDir, 0755)
		if err != nil {
			global.LogError(fmt.Sprintf("mergeParsedFile create merge data directory error: %s\n", err.Error()))
			return
		}
	}
	// 计算文件大小
	fileSize := len(fileContent)
	// 计算切分后的文件数量
	splitFileNum := (fileSize + 1024*100 - 1) / (1024 * 100)
	// 切分文件并写入输出文件
	for i := 0; i < splitFileNum; i++ {
		start := i * 1024 * 100
		end := (i + 1) * 1024 * 100
		if i == splitFileNum-1 {
			end = fileSize
		}
		outputFile := fmt.Sprintf("%s%d", outputFilePrefix, startNum+i)
		outputFileContent := fileContent[start:end]
		err = ioutil.WriteFile(outputFile, outputFileContent, 0644)
		if err != nil {
			global.LogError(fmt.Sprintf("splitAOF save split file: %s error: %s\n", outputFile, err.Error()))
			return
		}
		global.LogInfo(fmt.Sprintf("Split file %d: %s\n", i, outputFile))
	}
}

func mergeParsedFile() {
	global.LogInfo(fmt.Sprintf("mergeParsedFile Start merging ...\n"))
	// merge files
	parsedFiles, _ := ioutil.ReadDir(global.DataDir)
	// sort by name a -> z
	sort.SliceStable(parsedFiles, func(i, j int) bool {
		return parsedFiles[i].Name() < parsedFiles[j].Name()
	})
	outputFileName := filepath.Join(global.DataDir, "aof.merged")
	merged, err := os.OpenFile(outputFileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		global.LogError(fmt.Sprintf("mergeParsedFile create merged file error: %s\n", err.Error()))
		return
	}
	defer merged.Close()
	for _, item := range parsedFiles {
		name := item.Name()
		if !strings.HasPrefix(name, "aof.split_") || !strings.HasSuffix(name, ".parsed") || item.IsDir() {
			continue
		}
		content, err := ioutil.ReadFile(filepath.Join(global.DataDir, name))
		if err != nil {
			global.LogError(fmt.Sprintf("mergeParsedFile read file: %s | error: %s\n", name, err.Error()))
			continue
		}
		_, err = merged.Write(content)
		if err != nil {
			global.LogError(fmt.Sprintf("mergeParsedFile append content into merged file error: %s\n", err.Error()))
			continue
		}
		os.Remove(filepath.Join(global.DataDir, name))
	}
	global.LogInfo(fmt.Sprintf("Parse AOF success, output file: %s\n", outputFileName))
}
