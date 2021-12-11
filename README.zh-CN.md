## 目录
- [1、介绍](#1)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(1) 特性](#11)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(2) 架构](#12)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(3) 原理](#13)
- [2、安装](#2)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(1) Linux/Mac](#21)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(2) Windows](#22)
- [3、使用](#3)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(1) 输入文件](#31)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(2) 输出文件](#32)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(3) 使用示例](#33)
- [4、配置](#4)
- [5、性能](#5)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(1) 测试](#51)

## <span id="1">1、介绍</span>
parseAOF是一个简单快速的解析Redis AOF文件的工具

### <span id="11">(1) 特性</span>

- 代码简洁并且易于定制化
- 通过多协程加速解析AOF文件
- 解析后会生成命令列表，可用于日志查询等

### <span id="12">(2) 架构</span>
<img width="700" src="https://user-images.githubusercontent.com/35942268/145674949-1459562a-4555-493b-9aea-ed1d7d3f23a4.png">

### <span id="13">(3) 原理</span>
关于项目的相关原理可以参考[这篇文章](https://github.com/WGrape/Blog/issues/11)

## <span id="2">2、安装</span>

### <span id="21">(1) Linux / Mac</span>
```bash
git clone https://github.com/WGrape/parseAOF
cd parseAOF
go mod download
```

### <span id="22">(2) Windows</span>
暂不支持

## <span id="3">3、使用</span>
执行 ```start.sh``` 脚本，并传入待解析的AOF文件路径

```bash
bash ./start.sh ./data/appendonly.aof
```

### <span id="31">(1) 输入文件</span>
> 为了便于测试，可以使用 [./data/appendonly.aof](./data/appendonly.aof) 这个示例输入文件

开始执行前，传递AOF文件路径参数给 ```start.sh``` 脚本, AOF文件内容如下所示

```text
*2
$6
SELECT
$1
0
... ...
```


### <span id="32">(2) 输出文件</span>
> 为了便于测试，可以使用 [./data/aof.merged](./data/aof.merged) 这个示例输出文件

解析完成后，会在 ```data``` 目录下生成 [aof.merged](./data/aof.merged) 文件，其内容如下所示

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

### <span id="33">(3) 使用示例</span>

<img width="770" src="https://user-images.githubusercontent.com/35942268/144350765-6409d955-5f99-4218-81a5-c6ea840a749b.png" />

## <span id="4">4、配置</span>

[配置文件](./config/config.yml) ：```config/config.yml```

| Key | Value | Default | Detail |
| --- | :----:  | :---: | :---: |
| debug | ```false``` / ```true``` | ```false``` | 调试模式 |
| maxRoutines | int | 1024 | 允许最大协程数 |

## <span id="5">5、性能</span>

- 平均以 ```50000行/s``` 的速度解析AOF文件
- 最大支持解析1GB的AOF文件

### <span id="51">(1) 测试</span>

| Id | Lines | Size | Cost | CPU |
| --- | :----:  | :---: | :---: | :---: |
| 1 | 1,2301,117行 | 39MB | 3m50s | <=65% |
| 2 | 3,435,263行 | 13MB | 1m12s | <=65% |
| 3 | 1,043,700行 | 6.6MB | 38s | <=65% |
