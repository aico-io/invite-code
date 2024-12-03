# invite-code
一个Golang实现的邀请码生成器，根据ID生成对应邀请码，连续的ID生成的邀请码不连续，支持通过邀请码反推ID

An invitation code generator implemented in Golang, which generates corresponding invitation codes based on IDs. The invitation codes generated from consecutive IDs are not consecutive, and it supports reverse inferring the IDs from the invitation codes.

## Install

```shell
go get github.com/aico-io/invite-code@v0.1.0
```

## Example usage

```go
package main

import (
	"fmt"
	invite "github.com/aico-io/invite-code"
)

func main() {
	// 指定ID的类型(支持各种整型,但实际不能使用负数)以生成对应的邀请码生成器
	// 默认CHARSET为移除易混淆的字母数字后的57个字符
	// 可自定义,但注意需保证字符不重复,且如果字符集有符号,如-,可能会生成------这样的不友好邀请码
	g := invite.NewGenerator[uint16](invite.CHARSET, 6)
	fmt.Println(g.MaxSupportID())
	
	// 通过一个现有的非负整数ID生成对应的邀请码 
	curUID := uint16(1001)
	code, err := g.Encode(curUID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("curUID: %d code:%s\n", curUID, code)
	
	// 通过邀请码反推ID, 需保证邀请码生成器的CHARSET和长度一致
	id := g.Decode(code)
	fmt.Printf("code:%s id:%d\n", code, id)
}
```

- CHARSET为移除易混淆的字母数字后的57个字符
```text
23456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz
```
- 生成邀请码长度为6时，ID支持范围是 `0-34296447248(len(CHARSET)^6-1)`
```text
ID:0 code:999999
ID:1 code:7pMEF7
ID:2 code:FqLQMF
ID:3 code:EycjQE
ID:4 code:MrGcLM
ID:5 code:pbryqp
```
- 生成邀请码长度为7时，ID支持范围是 `0-1954897493192(len(CHARSET)^7-1)`
```text
ID:0 code:9999999
ID:1 code:7FMQ7Ep
ID:2 code:FMLcFQq
ID:3 code:EQc5Ejy
ID:4 code:MLG4Mcr
ID:5 code:pqrkpyb
```
- 生成邀请码长度为8时，ID支持范围是 `0-111429157112000(len(CHARSET)^8-1)`
```text
ID:0 code:99999999
ID:1 code:7EQ7MdFp
ID:2 code:FQcFL3Mq
ID:3 code:Ej5EcHQy
ID:4 code:Mc4MGnLr
ID:5 code:pykprgqb
```