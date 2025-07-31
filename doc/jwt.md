# 什么是 JWT Token

**Token（令牌）** 是一种用于身份验证和授权的机制，它是在用户登录后生成的一段字符串，用于识别用户的身份信息，
以便在用户进行后续操作时进行身份验证和授权
**Token**的生成和解析通常采用**JWT（JSON Web Token）** 协议，它是一种轻量级的身份验证和授权机制。JWT由三部分组成：
**头部（Header）、载荷（Payload）和签名（Signature）**，其中头部和载荷都是使用Base64 编码的JSON 对象，签名是由头部、载荷和密钥生成的哈希值

# JWT 的生成过程

- 创建 JWT 头部（Header）对象，包含算法和类型信息
- 创建 JWT 载荷（Payload）对象，包含用户信息和过期时间等信息
- 使用密钥对 JWT 头部和载荷进行签名生成 JWT 签名（Signature）
- 将 JWT 头部、载荷和签名拼接成一个字符串形式的 JWT

## JWT 的几种签名算法
最常见的 JWT 签名算法有 **HS256(HMAC-SHA256)、RS256(RSA-SHA256)、ES256(ECDSA-SHA256)**
这三种算法都是一种消息签名算法，得到的都只是一段无法还原的签名。区别在于消息签名与签名验证需要的 「key」不同。 
1. HS256 使用同一个「secret_key」进行签名与验证（对称加密）。一旦 secret_key 泄漏，就毫无安全性可言了。
   - 因此 HS256 只适合集中式认证，签名和验证都必须由可信方进行。
   - 传统的单体应用广泛使用这种算法，但是请不要在任何分布式的架构中使用它！
2. RS256 是使用 RSA 私钥进行签名，使用 RSA 公钥进行验证。公钥即使泄漏也毫无影响，只要确保私钥安全就行。
   - RS256 可以将验证委托给其他应用，只要将公钥给他们就行。
3. ES256 和 RS256 一样，都使用私钥签名，公钥验证。算法速度上差距也不大，但是它的签名长度相对短很多（省流量），并且算法强度和 RS256 差不多。

ES256 使用 ECDSA 进行签名，它的安全性和运算速度目前和 RS256 差距不大，但是拥有更短的签名长度。
对于需要频繁发送的 JWT 而言，更短的长度长期下来可以节约大量流量。

因此更推荐使用 ES256 算法。
## 使用 OPpenSSL 生成 RSA/ECC 公私钥
```shell
# 1. 生成 2048 位（不是 256 位）的 RSA 密钥
openssl genrsa -out rsa-private-key.pem 2048

# 2. 通过密钥生成公钥
openssl rsa -in rsa-private-key.pem -pubout -out rsa-public-key.pem
```
ES256 使用 ECDSA 算法进行签名，该算法使用 ECC 密钥，生成命令如下：
```shell
# 1. 生成 ec 算法的私钥，使用 prime256v1 算法，密钥长度 256 位。（强度大于 2048 位的 RSA 密钥）
openssl ecparam -genkey -name prime256v1 -out ecc-private-key.pem
# 2. 通过密钥生成公钥
openssl ec -in ecc-private-key.pem -pubout -out ecc-public-key.pem
```
## 调试
https://www.jwt.io/



# Gin JWT
## 安装
```bash
export GO111MODULE=on
go get github.com/appleboy/gin-jwt/v2
```
```go
package jwt
import "github.com/appleboy/gin-jwt/v2"
```
## 使用案例
https://github.com/appleboy/gin-jwt/blob/master/_example/basic/server.go

