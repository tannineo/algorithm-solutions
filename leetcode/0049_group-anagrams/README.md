# 0049_group-anagrams

output 是一个二维slice, 不用在意里面元素的顺序.

## 首次尝试 TLE

`func groupAnagramsTLE(strs []string) [][]string`

对每一行, 维护他们的特征map, 这有点像字母的 bag-of-words.

用了go的`reflect`的`func DeepEqual(x, y interface{}) bool`去比较两个map是否一致, 超时.

## 优化后 AC

`func groupAnagrams(strs []string) [][]string`

更换思路:

- 对每个字符串sort, 然后比较字符串

Accepted. 说明费时的例子: `len(strs)`比较大, 遍历它很费时.
