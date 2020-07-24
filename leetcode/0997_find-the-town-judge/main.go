package main

func findJudge(N int, trust [][]int) int {
	if N == 1 {
		return 1
	}

	trustedCountsMap := make(map[int]int)
	trustOthersMap := make(map[int]struct{})

	for _, v := range trust {
		trustOthersMap[v[0]] = struct{}{}
		trustedCountsMap[v[1]]++
	}
	answer := -1
	for k, v := range trustedCountsMap {
		if v == N-1 {
			if _, ok := trustOthersMap[k]; !ok {
				if answer == -1 {
					answer = k
				} else {
					return -1
				}
			}
		}
	}
	return answer
}

func main() {}
