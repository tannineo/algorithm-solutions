const _ = require('lodash')

burteForce = nums => {
  const requireTimes = nums.length / 2

  for (let toTest of nums) {
    let count = 0
    for (let element of nums) {
      if (element === toTest) {
        ++count
      }
    }
    if (count > requireTimes) {
      return toTest
    }
  }

  return 0
}

hashmapWay = nums => {
  const requireTimes = nums.length / 2
  const hashmap = new Map()

  for (let element of nums) {
    tempCnt = hashmap.get(element)
    if (!tempCnt) {
      hashmap.set(element, 1)
    } else {
      hashmap.set(element, tempCnt + 1)
    }
  }

  for (let [key, value] of hashmap) {
    if (value > requireTimes) return key
  }

  return 0
}

randomized = nums => {
  const requireTimes = nums.length / 2

  let flag = true
  let result = 0
  while (flag) {
    const toTest = nums[_.random(0, nums.length)]
    let count = 0
    for (let element of nums) {
      if (toTest === element) ++count
    }
    if (count > requireTimes) {
      flag = false
      result = toTest
    }
  }
  return result
}

votingBit = nums => {
  const requireTimes = nums.length / 2

  let bitsCount = new Array(32)
  _.fill(bitsCount, 0)

  for (let element of nums) {
    let bit = 0
    while (element !== 0) {
      if ((element & 1) === 1) ++bitsCount[bit]
      element >>>= 1 // right shift with sign
      ++bit
    }
  }

  let result = 0
  for (let i = bitsCount.length - 1; i >= 0; --i) {
    result <<= 1
    if (bitsCount[i] > requireTimes) result |= 1
  }

  return result
}

votingBoyerMoore = nums => {
  let count = 0
  let result
  for (let element of nums) {
    if (count === 0) {
      count = 1
      result = element
    } else if (element === result) {
      ++count
    } else {
      --count
    }
  }

  return result
}

fullSortWay = nums => {
  const median = Math.floor(nums.length / 2)

  nums.sort()

  return nums[median]
}

partialSortWay = nums => {
  // TODO
  return 0
}

divideAndConquer = function (nums, start, end) {
  if (_.isUndefined(start) && _.isUndefined(end)) {
    let [result, _] = divideAndConquer(nums, 0, nums.length)
    return result
  }

  if (start === end) return null
  if (start === end - 1) return [nums[start], 1]

  const mid = Math.floor((start + end) / 2)

  const obsLeft = divideAndConquer(nums, start, mid)
  const obsRight = divideAndConquer(nums, mid, end)

  if (obsLeft === null && obsRight === null) {
    return null // should not happen with leagal input
  } else if (obsLeft === null) {
    return obsRight
  } else if (obsRight === null) {
    return obsLeft
  }

  if (obsLeft[0] === obsRight[0]) return [obsLeft[0], obsLeft[1] + obsRight[1]]

  const leftCount = nums.slice(start, end).filter(e => e === obsLeft[0]).length
  const rightCount = nums.slice(start, end).filter(e => e === obsRight[0])
    .length

  return leftCount > rightCount
    ? [obsLeft[0], leftCount]
    : [obsRight[0], rightCount]
}

/**
 * @param {number[]} nums
 * @return {number}
 */
const majorityElement = votingBoyerMoore
