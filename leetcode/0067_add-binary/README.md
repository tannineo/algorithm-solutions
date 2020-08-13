# 0067_add-binary

水. 注意处理进位.

string还是习惯用`rune[]`处理.

```go
func addBinary(a string, b string) string {
    lenA, lenB := len(a), len(b)
    curA, curB := lenA-1, lenB-1

    result := []byte{}
    carry := 0
    for curA >= 0 || curB >= 0 {
        temp := carry
        if curA >= 0 {
            temp += int(a[curA] - '0')
        }
        if curB >= 0 {
            temp += int(b[curB] - '0')
        }
        carry = temp / 2
        temp %= 2
        result = append([]byte{'0' + byte(temp)}, result...)

        curA--
        curB--
    }
    if carry > 0 {
        result = append([]byte{'1'}, result...)
    }

    return string(result)
}
```
