<p align="center">
<img width="350" alt="img" src="https://user-images.githubusercontent.com/35942268/144242038-940e428f-5a99-4bcf-9d68-5d9e4f9b7a40.png">
</p>

<p align="center">
    <img src="https://img.shields.io/badge/Go-1.16+-blue.svg">
    <a href="https://app.travis-ci.com/github/WGrape/parseAOF"><img src="https://app.travis-ci.com/WGrape/parseAOF.svg?branch=main"><a>
    <img src="https://img.shields.io/badge/Document-中文/English-orange.svg">
    <img src="https://img.shields.io/badge/License-MIT-green.svg">
</p>

<div align="center">    
    <p>A simple and fast tool to parse the AOF file of redis</p>
    <p>Document ：<a href="/README.zh-CN.md">中文</a> / <a href="/README.md">English</a></p>
</div>


## Content
- [1、Features](#1)
- [2、Install](#2)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(1) Linux/Mac](#21)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(2) Windows](#22)
- [3、Usage](#3)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(1) The output file](#31)

## <span id="1">1、Features</span>
- Code is clean, simple and easy to customize
- Speed up parsing through multiple goroutines

## <span id="2">2、Install</span>

### <span id="21">(1) Linux/Mac</span>
```bash
cd ~

git clone https://github.com/WGrape/parseAOF

# Move your aof file to the data directory of parseAOF.
# In this way, the files generated during the whole process will be in the data directory, which is convenient for management
mv your_aof_file.aof ./parseAOF/data/appendonly.aof
```

### <span id="22">(2) Windows</span>
Windows is temporarily not supported

## <span id="3">3、Usage</span>
Run the ```start.sh``` script with the path of the aof file

```bash
bash ./start.sh ./data/appendonly.aof
```

### <span id="31">(1) The output file</span>
After the parsing is complete, the file ```aof.merged``` will be generated in the directory of ```data```, the content is as follows

```text
--------------------parseAOF | version=0.5.0--------------------
SELECT 0 
set key1 1 
set key2 2 
set key3 3 
sadd key4 1 2 3 4 
lpush key5 1 2 3 4 5 
zadd key6 1 2 3 4 5 6 
```
