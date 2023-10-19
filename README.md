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
- [Content](#content)
- [1、Introduction](#1introduction)
  - [(1) Features](#1-features)
  - [(2) Architecture](#2-architecture)
- [2、Build](#2build)
- [3、Usage](#3usage)
  - [(1) The input file](#1-the-input-file)
  - [(2) The output file](#2-the-output-file)
  - [(3) Example](#3-example)
- [4、Performance](#4performance)
  - [(1) Testing](#1-testing)

## 1、Introduction
A simple and fast tool to parse the AOF file of redis

### (1) Features
- Code is clean, simple and easy to customize
- Speed up parsing through multiple goroutines
- A list of commands will be generated after parsing for log querying

### (2) Architecture
<img width="700" src="https://user-images.githubusercontent.com/35942268/145674949-1459562a-4555-493b-9aea-ed1d7d3f23a4.png">

## 2、Build

```bash
git clone https://github.com/WGrape/parseAOF
cd parseAOF
go mod download
make build
```
## 3、Usage
Run the binary under `bin` dir `parseAOF_<os>_<arch>`  with the path of the aof file

```bash
./bin/parseAOF_macos_arm64 -i ~/Download/appendonly.aof -r 8
./bin/parseAOF_macos_arm64 -h
parse redis aof to readable

Usage:
  parseAOF [flags]

Flags:
  -h, --help            help for parseAOF
  -i, --input string    input AOF file path
  -o, --output string   output dir path
  -r, --routines int    max goroutines (default 8)
```

### (1) The input file
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

### (2) The output file
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

### (3) Example
![example](example.png)

## 4、Performance

- The average speed to parse is ```50000 lines/s```
- The maximum size of the aof is 1GB

### (1) Testing

| Id | Lines | Size | Cost | CPU |
| --- | :----:  | :---: | :---: | :---: |
| 1 | 1,2301,117 | 39MB | 3m50s | <=65% |
| 2 | 3,435,263 | 13MB | 1m12s | <=65% |
| 3 | 357,850 | 8.6MB | 3.47s | <=113% |
