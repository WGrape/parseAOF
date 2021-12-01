## 目录
- [1、特性](#1)
- [2、安装](#2)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(1) Linux/Mac](#21)
- &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[(2) Windows](#22)
- [3、使用](#3)

## <span id="1">1、特性</span>
- 代码简洁并且易于定制化
- 通过多协程加速解析AOF文件

## <span id="2">2、安装</span>

### <span id="21">(1) Linux / Mac</span>
```bash
git clone https://github.com/WGrape/parseAOF
```

### <span id="22">(2) Windows</span>
暂不支持

## <span id="3">3、Usage</span>
执行 ```start.sh``` 脚本，并传入待解析的AOF文件路径

```bash
bash ./start.sh ./data/appendonly.aof
```


