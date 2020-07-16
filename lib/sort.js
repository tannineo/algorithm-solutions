// sorting algorithms ref:
// https://www.geeksforgeeks.org/sorting-algorithms/

/**
 *
 *
 * ref https://stackoverflow.com/questions/29145520/how-is-nth-element-implemented
 *
 * @param {Array} array array to sort
 * @param {number} first range, [first, last)
 * @param {number} nth the element to find
 * @param {number} last range, [first, last)
 * @param {*} depth
 */
const introselect = function (array, first, nth, last, depth) {
  // insertion_sort
}

/**
 * return the nth element of array after a partial sorting
 *
 * array will be partial sorted (side effect)
 * the elements before nth is smaller than nth
 * the elements after nth is bigger then nth
 *
 * ref https://stackoverflow.com/questions/29145520/how-is-nth-element-implemented
 *
 * @param {Array} array array to sort
 * @param {number} first range, [first, last)
 * @param {number} nth the element to find
 * @param {number} last range, [first, last)
 * @returns {number} the nth element
 */
const nthElement = function (array, first, nth, last) {
  if (nth >= last || nth < first) {
    return null
  }

  introselect(array, first, nth, last, Math.log2((last - first) * 2))
}

module.export = {
  nthElement,
}
