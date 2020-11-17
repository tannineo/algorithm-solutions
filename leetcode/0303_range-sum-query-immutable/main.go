package main

type NumArray struct {
	sumAll       int
	prefixSumAll []int
	surfixSumAll []int
}

func Constructor(nums []int) NumArray {
	sumAll := 0
	prefixSumAll := make([]int, len(nums))
	surfixSumAll := make([]int, len(nums)+1)
	for i, v := range nums {
		prefixSumAll[i] = sumAll
		sumAll += v
	}
	sumAll = 0
	surfixSumAll[len(surfixSumAll)-1] = 0
	for i := len(nums) - 1; i >= 0; i-- {
		sumAll += nums[i]
		surfixSumAll[i] = sumAll
	}
	return NumArray{sumAll, prefixSumAll, surfixSumAll}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.sumAll - this.prefixSumAll[i] - this.surfixSumAll[j+1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
