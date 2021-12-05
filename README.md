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
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(1) The input file](#31)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(2) The output file](#32)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(3) Example](#33)
- [4、Configuration](#4)

## <span id="1">1、Features</span>
- Code is clean, simple and easy to customize
- Speed up parsing through multiple goroutines
- A list of commands will be generated after parsing for log querying

## <span id="2">2、Install</span>

### <span id="21">(1) Linux/Mac</span>
```bash
git clone https://github.com/WGrape/parseAOF
cd parseAOF
go mod download
```

### <span id="22">(2) Windows</span>
Windows is temporarily not supported

## <span id="3">3、Usage</span>
Run the ```start.sh``` script with the path of the aof file

```bash
bash ./start.sh /path/appendonly.aof
```

### <span id="31">(1) The input file</span>
> Here's an example input file [./data/appendonly.aof](./data/appendonly.aof) for you to test

Before running, pass the path of the aof file to the ```start.sh``` script, the content is as follows

```text
*2
$6
SELECT
$1
0
... ...
```

### <span id="32">(2) The output file</span>
> Here's an example output file [./data/aof.merged](./data/aof.merged) for you to test

After the parsing is complete, the file [aof.merged](./data/aof.merged) will be generated in the directory of ```data```, the content is as follows

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

### <span id="33">(3) Example</span>
<img width="770" src="https://user-images.githubusercontent.com/35942268/144350765-6409d955-5f99-4218-81a5-c6ea840a749b.png" />

## <span id="4">4、Configuration</span>

[Config file](./config/config.yml) ：```config/config.yml```

| Key | Value | Default | Detail |
| --- | :----:  | :---: | :---: |
| debug | ```false``` / ```true``` | ```false``` | debug mode |
| maxRoutines | int | 1024 | allow max number of goroutines |

## <span id="5">5、Performance</span>

- The average speed to parse is ```50000 lines/s```
- The maximum size of the aof is 1GB

### <span id="51">(1) 测试</span>

| Id | Lines | Size | Cost | CPU |
| --- | :----:  | :---: | :---: | :---: |
| 1 | 1,2301,117行 | 39MB | 3m50s | <=65% |
| 2 | 3,435,263行 | 13MB | 1m12s | <=65% |
| 3 | 1,043,700行 | 6.6MB | 38s | <=65% |
