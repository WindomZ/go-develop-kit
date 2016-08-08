# go-develop-kit
一些Golang项目常用的开发组件

## Cache
* [cache](https://github.com/WindomZ/go-develop-kit/tree/master/cache)
    * [freecache](https://github.com/WindomZ/go-develop-kit/tree/master/cache/freecache)
    * [gocache](https://github.com/WindomZ/go-develop-kit/tree/master/cache/gocache)
* [numcache](https://github.com/WindomZ/go-develop-kit/tree/master/cache/numcache)

## GoogleAuthenticator
* [One-time Password](https://github.com/WindomZ/go-develop-kit/tree/master/googleauth/otp)
    * HMAC-Based One-time Password (HOTP) algorithm specified in RFC 4226
    * Time-based One-time Password (TOTP) algorithm specified in RFC 6238

## Math
* [float64](https://github.com/WindomZ/go-develop-kit/blob/master/math/float.go)
    * 加减乘除,四舍五入
    * 消除浮点计算精度误差

## Mutex
* [mutex](https://github.com/WindomZ/go-develop-kit/blob/master/mutex/mutex.go)
    * 仍在开发完善 :)

## Security
* salt
    * [hmac](https://github.com/WindomZ/go-develop-kit/blob/master/security/salt/hmac.go)
    * [md5](https://github.com/WindomZ/go-develop-kit/blob/master/security/salt/md5.go)

## Unit
* [currency](https://github.com/WindomZ/go-develop-kit/blob/master/unit/currency.go)
    * 货币单位, 多格式映射
* [float price](https://github.com/WindomZ/go-develop-kit/blob/master/unit/float_price.go)
    * 价格单位, 浮点计算
* [int price](https://github.com/WindomZ/go-develop-kit/blob/master/unit/int_price.go)
    * 价格单位, 整数计算