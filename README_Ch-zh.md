# go-develop-kit

[![Build Status](https://travis-ci.org/WindomZ/go-develop-kit.svg?branch=master)](https://travis-ci.org/WindomZ/go-develop-kit)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)

> 我的Go项目常用的开发组件

[英文文档](https://github.com/WindomZ/go-develop-kit/blob/master/README.md#readme)

## 安装

```bash
go get github.com/WindomZ/go-develop-kit...
```

## 使用

### Bytes
- [bytes] (https://github.com/WindomZ/go-develop-kit/tree/master/bytes)
    - Format 将字节数转换成可读字符串
    - Parse 将可读字符串转换成字节数

### Cache
- [cache](https://github.com/WindomZ/go-develop-kit/tree/master/cache)
    - [freecache](https://github.com/WindomZ/go-develop-kit/tree/master/cache/freecache)
    - [gocache](https://github.com/WindomZ/go-develop-kit/tree/master/cache/gocache)
- [numcache](https://github.com/WindomZ/go-develop-kit/tree/master/cache/numcache)

### GoogleAuthenticator
- [One-time Password](https://github.com/WindomZ/go-develop-kit/tree/master/googleauth/otp)
    - HMAC-Based One-time Password (HOTP) algorithm specified in RFC 4226
    - Time-based One-time Password (TOTP) algorithm specified in RFC 6238

### IP
- [ip](https://github.com/WindomZ/go-develop-kit/blob/master/iputil/ip.go)
    - 获取当前的IP地址
    - 获取当前的IP地址尾部数字

### Math
- [float64](https://github.com/WindomZ/go-develop-kit/blob/master/math/float.go)
    - 加减乘除,四舍五入
    - 消除浮点计算精度误差

### Mutex
- [mutex](https://github.com/WindomZ/go-develop-kit/blob/master/mutex/mutex.go)
    - 仍在开发完善 :)

### Path
[README](https://github.com/WindomZ/go-develop-kit/blob/master/path#readme)

- 获取用户当前**`HOME`目录**路径
- **文件操作**功能集合

### Queue
- [queue](https://github.com/WindomZ/go-develop-kit/blob/master/queue/queue.go)
    - Push 推入队列
    - Pull 推出队列，FIFO
    - Exchange 安全地改变队列容量

### Security
- salt
    - [hmac](https://github.com/WindomZ/go-develop-kit/blob/master/security/salt/hmac.go)
        - hmac-md5
    - [md5](https://github.com/WindomZ/go-develop-kit/blob/master/security/salt/md5.go)
        - md5-prefix

### StringUtil
- [string](https://github.com/WindomZ/go-develop-kit/blob/master/stringutil/string.go)
    - 截取字符串
- [regexp](https://github.com/WindomZ/go-develop-kit/blob/master/stringutil/regexp.go)
    - 通过正则表达式来截取字符串

### Unit
- [currency](https://github.com/WindomZ/go-develop-kit/blob/master/unit/currency.go)
    - 货币单位, 多格式映射
- [float price](https://github.com/WindomZ/go-develop-kit/blob/master/unit/float_price.go)
    - 价格单位, 浮点计算
- [int price](https://github.com/WindomZ/go-develop-kit/blob/master/unit/int_price.go)
    - 价格单位, 整数计算

### UUID
- [uuid](https://github.com/WindomZ/go-develop-kit/blob/master/uuid/uuid.go)
    - 常用的UUID安全生成方法
- [validator](https://github.com/WindomZ/go-develop-kit/blob/master/uuid/validator.go)
    - 常用的UUID验证器
- [tool](https://github.com/WindomZ/go-develop-kit/blob/master/uuid/tool.go)
    - 常用的UUID转换工具

## 贡献

欢迎你提交PR、汇报Bugs、在[issues page](https://github.com/WindomZ/go-develop-kit/issues)提出新想法、新要求或者讨论问题，
我很乐意能一起参与。

如果你喜欢这个项目，可以点下 :star: 以支持！

## 许可

[MIT](https://github.com/WindomZ/go-develop-kit/blob/master/LICENSE)
