# go-develop-kit

[![Build Status](https://travis-ci.org/WindomZ/go-develop-kit.svg?branch=master)](https://travis-ci.org/WindomZ/go-develop-kit)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](https://opensource.org/licenses/MIT)

> My development kits for Golang.

[中文文档](https://github.com/WindomZ/go-develop-kit/blob/master/README_Ch-zh.md#readme)

## Install

```bash
go get github.com/WindomZ/go-develop-kit...
```

## Usage

### Bytes
- [bytes](https://github.com/WindomZ/go-develop-kit/tree/master/bytes/bytes.go)
    - Format formats bytes integer to human readable string
    - Parse parses human readable bytes string to bytes integer

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
    - Get the current IP address
    - Get the end of current IP address

### Math
- [float64](https://github.com/WindomZ/go-develop-kit/blob/master/math/float.go)
    - Add, subtract, multiply, divide and round
    - Calculation of big number, eliminate the error precision

### Mutex
- [mutex](https://github.com/WindomZ/go-develop-kit/blob/master/mutex/mutex.go)
    - Still in development :)

### Path
[README](https://github.com/WindomZ/go-develop-kit/blob/master/path#readme)

- Get the user's **`home` directory**
- **File operation** function collection

### Queue
- [queue](https://github.com/WindomZ/go-develop-kit/blob/master/queue/queue.go)
    - Push push the queue.
    - Pull pull the queue, FIFO.
    - Exchange safely change queue capacity.

### Security
- salt
    - [hmac](https://github.com/WindomZ/go-develop-kit/blob/master/security/salt/hmac.go)
        - hmac-md5
    - [md5](https://github.com/WindomZ/go-develop-kit/blob/master/security/salt/md5.go)
        - md5-prefix

### StringUtil
- [string](https://github.com/WindomZ/go-develop-kit/blob/master/stringutil/string.go)
    - Split the string
- [regexp](https://github.com/WindomZ/go-develop-kit/blob/master/stringutil/regexp.go)
    - Split the string with regular expression

### Unit
- [currency](https://github.com/WindomZ/go-develop-kit/blob/master/unit/currency.go)
    - Unit currency, the multiple formats mapping
- [float price](https://github.com/WindomZ/go-develop-kit/blob/master/unit/float_price.go)
    - Unit price, floating point calculation
- [int price](https://github.com/WindomZ/go-develop-kit/blob/master/unit/int_price.go)
    - Unit price, integral calculation

### UUID
- [uuid](https://github.com/WindomZ/go-develop-kit/blob/master/uuid/uuid.go)
    - Common UUID safe generation method
- [validator](https://github.com/WindomZ/go-develop-kit/blob/master/uuid/validator.go)
    - Commonly used UUID validator
- [tool](https://github.com/WindomZ/go-develop-kit/blob/master/uuid/tool.go)
    - Commonly used UUID conversion tool

## Contributing

Welcome to pull requests, report bugs, suggest ideas and discuss **go-develop-kit**, 
i would love to hear what you think about **go-develop-kit** on [issues page](https://github.com/WindomZ/go-develop-kit/issues).

If you like it then you can put a :star: on it.

## License

[MIT](https://github.com/WindomZ/go-develop-kit/blob/master/LICENSE)
