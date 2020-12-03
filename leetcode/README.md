# leetcode

Problems list.

- [leetcode](#leetcode)
  - [important](#important)
  - [data structures](#data-structures)
    - [linked list](#linked-list)
    - [stack](#stack)
    - [tree](#tree)
    - [graph](#graph)
    - [heap](#heap)
  - [maths](#maths)
  - [dp](#dp)
  - [string](#string)
  - [bitwise](#bitwise)
  - [divide and conquer](#divide-and-conquer)
  - [misc](#misc)

## important

- [0023_merge-k-sorted-lists](./0023_merge-k-sorted-lists)
  - HARD: linked list, min-heap
- [0033_search-in-rotated-sorted-array](./0033_search-in-rotated-sorted-array)
  - binary search
- [0034_find-first-and-last-position-of-element-in-sorted-array](./0034_find-first-and-last-position-of-element-in-sorted-array)
  - binary search
- [0053_maximum-subarray](./0053_maximum-subarray)
  - dp, divide and conquer
- [0096_unique-binary-search-trees](./0096_unique-binary-search-trees)
  - BST, maths
- [0137_single-number-ii](./0137_single-number-ii)
  - bitwise
- [0142_linked-list-cycle-ii](./0142_linked-list-cycle-ii)
  - linked list, Floyd's Tortoise and Hare, loop detection
- [0148_sort-list](./0148_sort-list)
  - linked list, merge sort, divide and conquer
- [0169_majority-element](./0169_majority-element)
  - hashmap, divide and conquer, voting, randomized, partial sorting ...
- [0222_count-complete-tree-nodes](./0222_count-complete-tree-nodes)
  - binary tree, post-order traversal
- [0230_kth-smallest-element-in-a-bst](./0230_kth-smallest-element-in-a-bst)
  - BST, in-order traversal
- [0260_single-number-iii](./0260_single-number-iii)
  - bitwise
- [0279_perfect-squares](./0279_perfect-squares)
  - dp, maths
- [0287_find-the-duplicate-number](./0287_find-the-duplicate-number)
  - Floyd's Tortoise and Hare, loop detection
- [0300_longest-increasing-subsequence](./0300_longest-increasing-subsequence)
  - dp, binary search
- [0435_non-overlapping-intervals](./0435_non-overlapping-intervals)
  - dp, greedy
- [0441_arranging-coins](./0441_arranging-coins)
  - solving quadratic equation in one unknown
- [0450_delete-node-in-a-bst](./0450_delete-node-in-a-bst)
  - delete a node in a BST, with complex if-else statements
- [0518_coin-change-2](./0518_coin-change-2)
  - dp
- [0525_contiguous-array](./0525_contiguous-array)
  - "coloring the balloons"
- [0572_subtree-of-another-tree](./0572_subtree-of-another-tree)
  - tree, preorder-traversal, preprocessing
- [0621_task-scheduler](./0621_task-scheduler)
  - greedy
- [0704_binary-search](./0704_binary-search)
  - CLASSIC binary search
- [0912_sort-an-array](./0912_sort-an-array)
  - various kinds of sorting
- [0918_maximum-sum-circular-subarray](./0918_maximum-sum-circular-subarray)
  - dp
- [1218_longest-arithmetic-subsequence-of-given-difference](./1218_longest-arithmetic-subsequence-of-given-difference)
  - dp

## data structures

### linked list

- [0002_add-two-numbers](./0002_add-two-numbers)
- [0021_merge-two-sorted-lists](./0021_merge-two-sorted-lists)
- [0024_swap-nodes-in-pairs](./0024_swap-nodes-in-pairs)
- [0141_linked-list-cycle](./0141_linked-list-cycle)
- [0146_lru-cache](,/0146_lru-cache): double linked list with map
- [0147_insertion-sort-list](/0147_insertion-sort-list)
- [0203_remove-linked-list-elements](./0203_remove-linked-list-elements)
- [0206_reverse-linked-list](./0206_reverse-linked-list)
- [0237_delete-node-in-a-linked-list](./0237_delete-node-in-a-linked-list)
- [0328_odd-even-linked-list](./0328_odd-even-linked-list)
- [0430_flatten-a-multilevel-doubly-linked-list](./0430_flatten-a-multilevel-doubly-linked-list)
- [0445_add-two-numbers-ii](./0445_add-two-numbers-ii)
- [0707_design-linked-list](./0707_design-linked-list)
- [0876_middle-of-the-linked-list](./0876_middle-of-the-linked-list)
- [1429_first-unique-number](./1429_first-unique-number): double linked list with map

### stack

- [0020_valid-parentheses](./0020_valid-parentheses)
- [0155_min-stack](./0155_min-stack)
- [0901_online-stock-span](./0901_online-stock-span)

### tree

- [0096_binary-tree-inorder-traversal](./0096_binary-tree-inorder-traversal)
- [0098_validate-binary-search-tree](./0098_validate-binary-search-tree)
- [0099_recover-binary-search-tree](./0099_recover-binary-search-tree): HARD(?), BST, in-order traversal
- [0100_same-tree](./0100_same-tree)
- [0101_symmetric-tree](./0101_symmetric-tree)
- [0102_binary-tree-level-order-traversal](./0102_binary-tree-level-order-traversal)
- [0104_maximum-depth-of-binary-tree](./0104_maximum-depth-of-binary-tree)
- [0107_binary-tree-level-order-traversal-ii](./0107_binary-tree-level-order-traversal-ii)
- [0108_convert-sorted-array-to-binary-search-tree](./0108_convert-sorted-array-to-binary-search-tree)
- [0110_balanced-binary-tree](./0110_balanced-binary-tree)
- [0111_minimum-depth-of-binary-tree](./0111_minimum-depth-of-binary-tree)
- [0112_path-sum](./0112_path-sum)
- [0113_path-sum-ii](./0113_path-sum-ii)
- [0124_binary-tree-maximum-path-sum](./0124_binary-tree-maximum-path-sum): HARD
- [0129_sum-root-to-leaf-numbers](./0129_sum-root-to-leaf-numbers)
- [0144_binary-tree-preorder-traversal](./0144_binary-tree-preorder-traversal)
- [0145_binary-tree-postorder-traversal](./0145_binary-tree-postorder-traversal)
- [0208_implement-trie-prefix-tree](./0208_implement-trie-prefix-tree)
- [0226_invert-binary-tree](./0226_invert-binary-tree)
- [0235_lowest-common-ancestor-of-a-binary-search-tree](./0235_lowest-common-ancestor-of-a-binary-search-tree)
- [0236_lowest-common-ancestor-of-a-binary-tree](./0236_lowest-common-ancestor-of-a-binary-tree)
- [0257_binary-tree-paths](./0257_binary-tree-paths)
- [0270_closest-binary-search-tree-value](./0270_closest-binary-search-tree-value)
- [0297_serialize-and-deserialize-binary-tree](./0297_serialize-and-deserialize-binary-tree): HARD(?)
- [0315_count-of-smaller-numbers-after-self](./0315_count-of-smaller-numbers-after-self): HARD, BST
- [0337_house-robber-iii](./0337_house-robber-iii)
- [0429_n-ary-tree-level-order-traversal](./0429_n-ary-tree-level-order-traversal)
- [0437_path-sum-iii](./0437_path-sum-iii)
- [0449_serialize-and-deserialize-bst](./0449_serialize-and-deserialize-bst)
- [0501_find-mode-in-binary-search-tree](./0501_find-mode-in-binary-search-tree)
- [0508_most-frequent-subtree-sum](./0508_most-frequent-subtree-sum)
- [0530_minimum-absolute-difference-in-bst](./0530_minimum-absolute-difference-in-bst)
- [0543_diameter-of-binary-tree](./0543_diameter-of-binary-tree)
- [0589_n-ary-tree-preorder-traversal](./0589_n-ary-tree-preorder-traversal)
- [0617_merge-two-binary-trees](./0617_merge-two-binary-trees)
- [0662_maximum-width-of-binary-tree](./0662_maximum-width-of-binary-tree)
- [0669_trim-a-binary-search-tree](./0669_trim-a-binary-search-tree)
- [0687_longest-univalue-path](./0687_longest-univalue-path)
- [0700_search-in-a-binary-search-tree](./0700_search-in-a-binary-search-tree)
- [0701_insert-into-a-binary-search-tree](./0701_insert-into-a-binary-search-tree)
- [0814_binary-tree-pruning](./0814_binary-tree-pruning)
- [0872_leaf-similar-trees](./0872_leaf-similar-trees)
- [0965_univalued-binary-tree](./0965_univalued-binary-tree)
- [0968_binary-tree-cameras](./0968_binary-tree-cameras): HARD, post-order traversal, dp
- [0987_vertical-order-traversal-of-a-binary-tree](./0987_vertical-order-traversal-of-a-binary-tree)
- [0993_cousins-in-binary-tree](./0993_cousins-in-binary-tree)
- [1008_construct-binary-search-tree-from-preorder-traversal](./1008_construct-binary-search-tree-from-preorder-traversal)
- [1214_two-sum-bsts](./1214_two-sum-bsts)
- [1302_deepest-leaves-sum](./1302_deepest-leaves-sum)
- [1325_delete-leaves-with-a-given-value](./1325_delete-leaves-with-a-given-value)
- [1430_check-if-a-string-is-a-valid-sequence-from-root-to-leaves-path-in-a-binary-tree](./1430_check-if-a-string-is-a-valid-sequence-from-root-to-leaves-path-in-a-binary-tree)

### graph

- [0133_clone-graph](./0133_clone-graph)
- [0138_copy-list-with-random-pointer](./0138_copy-list-with-random-pointer)
- [0207_course-schedule](./0207_course-schedule)
- [0886_possible-bipartition](./0886_possible-bipartition): basic dfs in graph

### heap

- [1046_last-stone-weight](./1046_last-stone-weight)

## maths

- [0070_climbing-stairs](./0070_climbing-stairs)
- [0258_add-digits](./0258_add-digits)
- [0342_power-of-four](./0342_power-of-four)
- [0367_valid-perfect-square](./0367_valid-perfect-square): Newton's Method
- [0528_random-pick-with-weight](./0528_random-pick-with-weight)
- [0976_largest-perimeter-triangle](./0976_largest-perimeter-triangle)
- [1232_check-if-it-is-a-straight-line](./1232_check-if-it-is-a-straight-line): gcd

## dp

- [0055_jump-game](./0055_jump-game)
- [0062_unique-paths](./0062_unique-paths)
- [0064_minimum-path-sum](./0064_minimum-path-sum)
- [0072_edit-distance](./0072_edit-distance): HARD
- [0264_ugly-number-ii](./0264_ugly-number-ii)
- [0368_largest-divisible-subset](./0368_largest-divisible-subset)
- [0673_number-of-longest-increasing-subsequence](./0673_number-of-longest-increasing-subsequence)
- [0678_valid-parenthesis-string](./0678_valid-parenthesis-string)
- [0712_minimum-ascii-delete-sum-for-two-strings](./0712_minimum-ascii-delete-sum-for-two-strings)
- [0746_min-cost-climbing-stairs](./0746_min-cost-climbing-stairs)
- [1035_uncrossed-lines](./1035_uncrossed-lines)
- [1048_longest-string-chain](./1048_longest-string-chain): 0300-like dp with string ops.
- [1143_longest-common-subsequence](./1143_longest-common-subsequence)

## string

- [0560_subarray-sum-equals-k](./0560_subarray-sum-equals-k)

## bitwise

- [0190_reverse-bits](./0190_reverse-bits)
- [0201_bitwise-and-of-numbers-range](./0201_bitwise-and-of-numbers-range)
- [0231_power-of-two](./0231_power-of-two)
- [0338_counting-bits](./0338_counting-bits)
- [0461_hamming-distance](./0461_hamming-distance): turning right most 1
- [0476_number-complement](./0476_number-complement)
- [0957_prison-cells-after-n-days](./0957_prison-cells-after-n-days): preprocess

## divide and conquer

- [0153_find-minimum-in-rotated-sorted-array](./0153_find-minimum-in-rotated-sorted-array)
- [0154_find-minimum-in-rotated-sorted-array-ii](./0154_find-minimum-in-rotated-sorted-array-ii)
- [0973_k-closest-points-to-origin](./0973_k-closest-points-to-origin)

## misc

- [0003_longest-substring-without-repeating-characters](./0003_longest-substring-without-repeating-characters)
- [0005_longest-palindromic-substring](./0005_longest-palindromic-substring)
- [0006_zigzag-conversion](./0006_zigzag-conversion)
- [0035_search-insert-position](./0035_search-insert-position): binary search
- [0042_trapping-rain-water](./0042_trapping-rain-water): HARD(?)
- [0049_group-anagrams](./0049_group-anagrams)
- [0066_plus-one](./0066_plus-one)
- [0067_add-binary](./0067_add-binary)
- [0075_sort-colors](./0075_sort-colors)
- [0122_best-time-to-buy-and-sell-stock-ii](./0122_best-time-to-buy-and-sell-stock-ii)
- [0125_valid-palindrome](./0125_valid-palindrome)
- [0130_surrounded-regions](./0130_surrounded-regions): flood fill
- [0151_reverse-words-in-a-string](./0151_reverse-words-in-a-string)
- [0171_excel-sheet-column-number](./0171_excel-sheet-column-number)
- [0200_number-of-islands](./0200_number-of-islands)
- [0202_happy-number](./0202_happy-number)
- [0221_maximal-square](./0221_maximal-square)
- [0238_product-of-array-except-self](./0238_product-of-array-except-self)
- [0274_h-index](./0274_h-index)
- [0275_h-index-ii](./0275_h-index-ii)
- [0278_first-bad-version](./0278_first-bad-version): binary search
- [0283_move-zeroes](./0283_move-zeroes)
- [0303_range-sum-query-immutable](./0303_range-sum-query-immutable): prefix sum
- [0344_reverse-string](./0344_reverse-string)
- [0359_logger-rate-limiter](./0359_logger-rate-limiter)
- [0380_insert-delete-getrandom-o1](./0380_insert-delete-getrandom-o1): array with map
- [0383_ransom-note](./0383_ransom-note)
- [0387_first-unique-character-in-a-string](./0387_first-unique-character-in-a-string)
- [0392_is-subsequence](./0392_is-subsequence)
- [0402_remove-k-digits](./0402_remove-k-digits): greedy
- [0406_queue-reconstruction-by-height](./0406_queue-reconstruction-by-height): greedy
- [0412_fizz-buzz](./0412_fizz-buzz)
- [0438_find-all-anagrams-in-a-string](./0438_find-all-anagrams-in-a-string)
- [0442_find-all-duplicates-in-an-array](./0442_find-all-duplicates-in-an-array)
- [0451_sort-characters-by-frequency](./0451_sort-characters-by-frequency)
- [0463_island-perimeter](./0463_island-perimeter)
- [0468_validate-ip-address](./0468_validate-ip-address): unreal problem, wrong testcases
- [0520_detect-capital](./0520_detect-capital)
- [0540_single-element-in-a-sorted-array](./0540_single-element-in-a-sorted-array)
- [0567_permutation-in-string](./0567_permutation-in-string)
- [0733_flood-fill](./0733_flood-fill)
- [0771_jewels-and-stones](./0771_jewels-and-stones)
- [0844_backspace-string-compare](./0844_backspace-string-compare)
- [0986_interval-list-intersections](./0986_interval-list-intersections)
- [0994_rotting-oranges](./0994_rotting-oranges): flood fill
- [0997_find-the-town-judge](./0997_find-the-town-judge)
- [1029_two-city-scheduling](./1029_two-city-scheduling): greedy?
- [1137_n-th-tribonacci-number](./1137_n-th-tribonacci-number)
- [1277_count-square-submatrices-with-all-ones](./1277_count-square-submatrices-with-all-ones)
- [1286_iterator-for-combination](./1286_iterator-for-combination)
- [1426_counting-elements](./1426_counting-elements)
- [1427_perform-string-shifts](./1427_perform-string-shifts)
- [1428_leftmost-column-with-at-least-a-one](./1428_leftmost-column-with-at-least-a-one)
- [1529_bulb-switcher-iv](./1529_bulb-switcher-iv): greedy?
- [1572_matrix-diagonal-sum](./1572_matrix-diagonal-sum)
