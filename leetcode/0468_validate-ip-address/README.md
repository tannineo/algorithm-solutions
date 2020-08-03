# 0468_validate-ip-address

这一题是**垃圾**.

题目描述对于`IPv6`不全. 且testcase不符合Wiki上的描述: https://en.wikipedia.org/wiki/IPv6

起码`::1`应该是合法的IPv6. testcase只允许`::`出现在中间, 开头的`0`不能被省略.

用`.`a 或者`:`分割, 判断.

PS: go官方实现: https://golang.org/src/net/ip.go

```go
func parseIPv6(s string) (ip IP)
```
