package bolingboulders

type Cube struct {
	X, Y, Z int
}

func CalculateSurfaceArea(cubes []Cube) int {
	surfaceArea := 0
	// Iterate over cubes to check the neighbour cube
	for _, cube := range cubes {
		for _, dir := range [][]int{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}} {
			var adjecent Cube
			adjecent.X = cube.X + dir[0]
			adjecent.Y = cube.Y + dir[1]
			adjecent.Z = cube.Z + dir[2]
			if !isCubePresentToTheNext(cubes, adjecent) {
				surfaceArea++
			}
		}
	}
	return surfaceArea
}

// check all six side
func isCubePresentToTheNext(cubes []Cube, coordinate Cube) bool {
	for _, cube := range cubes {
		if cube.X == coordinate.X && cube.Y == coordinate.Y && cube.Z == coordinate.Z {
			return true
		}
	}
	return false
}
