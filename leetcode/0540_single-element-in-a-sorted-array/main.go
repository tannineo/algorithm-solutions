package main

func singleNonDuplicate(nums []int) int {
	lenNums := len(nums)
	if lenNums == 1 {
		return nums[0]
	}

	left := 0
	right := lenNums

	// [left, ..., guess, ... , right)
	for left < right {
		guess := (left + right) / 2

		if (guess == 0 || nums[guess] != nums[guess-1]) && (guess == lenNums-1 || nums[guess] != nums[guess+1]) {
			return nums[guess]
		}

		if guess+1 < lenNums && nums[guess] == nums[guess+1] {
			lenLeft := guess - left
			if lenLeft%2 == 0 { // discard left part
				left = guess + 2
			} else { // discard right part
				right = guess
			}
		} else if guess > 0 { // must be nums[guess] == nums[guess-1]
			lenLeft := guess - left - 1
			if lenLeft%2 == 0 { // discard left part
				left = guess + 1
			} else { // discard right part
				right = guess - 1
			}
		}
	}

	return -1 // won't be here
}

func main() {}
