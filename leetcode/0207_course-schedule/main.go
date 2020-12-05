package main

type Schedule struct {
	Val       int
	Prereq    []*Schedule
	TestRound int
}

func canFinish(numCourses int, prerequisites [][]int) bool {

	courses := make([]*Schedule, numCourses)

	// 1. construct graph
	for _, v := range prerequisites {
		if v[0] >= numCourses { // all the course occured should be in [0, numCourses)
			return false
		}
		if v[1] >= numCourses {
			return false
		}

		if courses[v[0]] == nil {
			courses[v[0]] = &Schedule{v[0], []*Schedule{}, 0}
		}

		if courses[v[1]] == nil {
			courses[v[1]] = &Schedule{v[1], []*Schedule{courses[v[0]]}, 0}
		} else {
			courses[v[1]].Prereq = append(courses[v[1]].Prereq, courses[v[0]])
		}
	}

	// 2. dfs
	for _, start := range courses {
		if start != nil {
			// start dfs here
			if dfsHelper(start, 1) {
				return false
			}
		}
	}

	return true
}

func dfsHelper(root *Schedule, testNumer int) bool {
	if root.TestRound != 0 && root.TestRound != testNumer {
		root.TestRound = 0 // clean
		return true
	}

	root.TestRound = testNumer

	for _, v := range root.Prereq {
		if dfsHelper(v, testNumer+1) {
			root.TestRound = 0 // clean
			return true
		}
	}

	root.TestRound = 0 // clean
	return false
}

func main() {
}
