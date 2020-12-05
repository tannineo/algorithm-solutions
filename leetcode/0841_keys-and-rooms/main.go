package main

func canVisitAllRooms(rooms [][]int) bool {

	// len(rooms) >= 1

	mapVisited := map[int]bool{}
	countVisited := 0

	// dfs with map, from room 0
	dfsWithMap(rooms, 0, mapVisited, &countVisited)

	return countVisited == len(rooms)
}

func dfsWithMap(rooms [][]int, i int, mapVisited map[int]bool, countVisited *int) {
	if _, ok := mapVisited[i]; ok {
		return // visited
	}

	mapVisited[i] = true // or false, whatever
	(*countVisited)++

	for _, v := range rooms[i] {
		dfsWithMap(rooms, v, mapVisited, countVisited)
	}
}
