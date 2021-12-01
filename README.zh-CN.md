## 目录
- [1、特性](#1)
- [2、安装](#2)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(1) Linux/Mac](#21)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(2) Windows](#22)
- [3、使用](#3)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(1) 输出文件](#31)

## <span id="1">1、特性</span>
- 代码简洁并且易于定制化
- 通过多协程加速解析AOF文件

## <span id="2">2、安装</span>

### <span id="21">(1) Linux / Mac</span>
```bash
cd ~

git clone https://github.com/WGrape/parseAOF

# 把需要解析的AOF文件移动到parseAOF的data目录下
# 这样整个过程中产生的文件都会在此data目录下，方便管理
mv your_aof_file.aof ./parseAOF/data/appendonly.aof
```

### <span id="22">(2) Windows</span>
暂不支持

## <span id="3">3、Usage</span>
执行 ```start.sh``` 脚本，并传入待解析的AOF文件路径

```bash
bash ./start.sh ./data/appendonly.aof
```

### <span id="31">(1) 输出文件</span>

解析完成后，会在```data```目录下生成```aof.merged```文件，内容如下所示

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
