# 0733_flood-fill

题目太直白了, flood fill.

```go
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

  originalColor := image[sr][sc]
  if originalColor == newColor { // no need to change color
    return image
  }
  image[sr][sc] = newColor
  // up
  if sr-1 >= 0 && image[sr-1][sc] == originalColor {
    floodFill(image, sr-1, sc, newColor)
  }
  // down
  if sr+1 < len(image) && image[sr+1][sc] == originalColor {
    floodFill(image, sr+1, sc, newColor)
  }
  // left
  if sc-1 >= 0 && image[sr][sc-1] == originalColor {
    floodFill(image, sr, sc-1, newColor)
  }
  // right
  if sc+1 < len(image[0]) && image[sr][sc+1] == originalColor {
    floodFill(image, sr, sc+1, newColor)
  }

  return image
}
```
