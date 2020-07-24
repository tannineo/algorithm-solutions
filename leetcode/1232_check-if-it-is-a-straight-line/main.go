package main

func gcd(a, b int) int {
	c := a % b
	for c != 0 {
		a = b
		b = c
		c = a % b
	}
	return b
}

func checkStraightLine(coordinates [][]int) bool {
	lenC := len(coordinates)
	if lenC == 2 {
		return true
	}

	// find a, b
	a := coordinates[0]
	var b []int
	i := 1
	for ; i < lenC; i++ {
		if coordinates[i][0] != a[0] || coordinates[i][1] != a[1] {
			b = coordinates[i]
			break
		}
	}

	if b == nil || i == lenC {
		return true
	}

	// calculate gradient
	// set the change of x (a->b) to >= 0
	if a[0] > b[0] {
		a, b = b, a
	}
	x := b[0] - a[0]
	y := b[1] - a[1]
	if x == 0 { // vertical
		for i++; i < lenC; i++ {
			if coordinates[i][0] != a[0] {
				return false
			}
		}
	} else if y == 0 { // horizontal
		for i++; i < lenC; i++ {
			if coordinates[i][1] != a[1] {
				return false
			}
		}
	} else {
		// because all elements are int
		// we need to cut x, y into the `finest` distance
		minuxFlag := false
		if y < 0 {
			minuxFlag = true
			y = -y
		}
		gcdXY := gcd(x, y)
		x = x / gcdXY
		y = y / gcdXY
		if minuxFlag {
			y = -y
		}

		// gradient = y/x
		for i++; i < lenC; i++ {
			disX := a[0] - coordinates[i][0]
			disY := a[1] - coordinates[i][1]
			if disX%x != 0 || disY%y != 0 {
				return false
			}
			step := disX / x
			if step*y != disY {
				return false
			}
		}
	}

	return true
}

func main() {

	println(gcd(319, 377)) // 29

	println(gcd(-319, 377)) // 29

}
