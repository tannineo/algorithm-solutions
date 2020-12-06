package main

func findOrder(numCourses int, prerequisites [][]int) []int {

	if numCourses == 0 {
		return []int{}
	}
	ans := []int{}
	if len(prerequisites) == 0 {
		for i := 0; i < numCourses; i++ {
			ans = append(ans, i)
		}
		return ans
	}

	// init edges
	graphNN := make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		graphNN[i] = []int{}
	}
	for _, v := range prerequisites {
		graphNN[v[0]] = append(graphNN[v[0]], v[1])
	}

	// record visited status
	// 0 - not visited
	// 1 - visiting in this dfs
	// 2 - visited
	visitState := make([]int, numCourses)

	// iterate and dfs
	for i := 0; i < numCourses; i++ {
		if dfsWithVisited(i, graphNN, &visitState, &ans) {
			return []int{}
		}
	}

	// reverse
	revAns := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		revAns[i] = ans[len(ans)-1-i]
	}
	return revAns
}

// return true when loop is detected
func dfsWithVisited(cur int, graphNN [][]int, visited, ans *[]int) bool {
	if (*visited)[cur] == 1 {
		return true
	}
	if (*visited)[cur] == 2 {
		return false
	}

	(*visited)[cur] = 1 // mark as visiting

	for _, v := range graphNN[cur] {
		if dfsWithVisited(v, graphNN, visited, ans) {
			return true
		}
	}

	(*visited)[cur] = 2 // mark as visited
	*ans = append([]int{cur}, (*ans)...)

	return false
}

func main() {
	findOrder(2, [][]int{[]int{1, 0}})
}
