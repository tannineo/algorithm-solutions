# 0278_first-bad-version

二分法

```go
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return            true if current version is bad
 *                    false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
  if n == 1 { // 1 <= n
    return 1
  }

  left := 1
  right := n

  for left < right {
    guess := (left + right) / 2
    if isBadVersion(guess) {
      // answer is in [left, guess]
      right = guess
    } else {
      // answer is in [guess+1, right]
      left = guess + 1
    }
  }

  return left
}

```
