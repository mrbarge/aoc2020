package main

import (
	"aoc2020/helper"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tile struct {
	data [][]bool
	id   int
}

func readTiles(data []string) []Tile {
	readingTile := false

	tiles := make([]Tile, 0)
	var tmpTile Tile
	for _, line := range data {
		if strings.HasPrefix(line, "Tile") {
			tileId, _ := strconv.Atoi(strings.Split(strings.Split(line, " ")[1], ":")[0])
			readingTile = true
			tmpTile = Tile{
				id: tileId,
			}
			continue
		}

		if line == "" {
			readingTile = false
			tiles = append(tiles, tmpTile)
			continue
		}

		if readingTile {
			row := make([]bool, len(line))
			for i, c := range line {
				if c == '#' {
					row[i] = true
				}
			}
			tmpTile.data = append(tmpTile.data, row)
		}
	}
	if readingTile {
		tiles = append(tiles, tmpTile)
	}
	return tiles
}

func (t Tile) Print() {
	fmt.Printf("Tile: %d\n", t.id)
	for _, row := range t.data {
		for _, col := range row {
			if col {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func (t Tile) FlipHorizontal() Tile {
	r := newTile(t.id, len(t.data[0]), len(t.data))
	for y := 0; y < len(t.data); y++ {
		for x := 0; x < len(t.data[y]); x++ {
			r.data[y][len(t.data[y])-1-x] = t.data[y][x]
		}
	}
	return r
}

func (t Tile) FlipVertical() Tile {
	r := newTile(t.id, len(t.data[0]), len(t.data))
	for y := 0; y < len(t.data); y++ {
		for x := 0; x < len(t.data[y]); x++ {
			r.data[len(t.data)-1-y][x] = t.data[y][x]
		}
	}
	return r
}

func (t Tile) Rotate() Tile {
	r := newTile(t.id, len(t.data[0]), len(t.data))
	for y := 0; y < len(t.data); y++ {
		for x := 0; x < len(t.data[y]); x++ {
			r.data[y][x] = t.data[len(t.data)-x-1][y]
		}
	}
	return r
}

func newTile(id int, x int, y int) Tile {
	data := make([][]bool, 0)
	for i := 0; i < y; i++ {
		data = append(data, make([]bool, x))
	}
	return Tile{
		id:   id,
		data: data,
	}
}

func (t Tile) Variants() []Tile {
	ret := make([]Tile, 0)
	tmp := t
	ret = append(ret, tmp)
	tmp = tmp.Rotate()
	ret = append(ret, tmp)
	tmp = tmp.Rotate()
	ret = append(ret, tmp)
	tmp = tmp.Rotate()
	ret = append(ret, tmp)
	tmp = t.FlipVertical()
	ret = append(ret, tmp)
	tmp = tmp.Rotate()
	ret = append(ret, tmp)
	tmp = tmp.Rotate()
	ret = append(ret, tmp)
	tmp = tmp.Rotate()
	ret = append(ret, tmp)
	return ret
}

func (t Tile) IsUpperLeft(tiles []Tile) bool {

	tileLeftMatch := 0
	tileUpperMatch := 0
	for _, tile := range tiles {
		if tile.id == t.id {
			continue
		}

		if tile.id == tileLeftMatch || tile.id == tileUpperMatch {
			continue
		}

		tmpTile := tile
		tmpOrig := t
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {

				// find left match
				if tmpTile.AdjacentHorizontal(tmpOrig) || tmpTile.AdjacentHorizontal(tmpOrig.FlipHorizontal()) {
					//fmt.Printf("Tile %d Found match 1 with tile %d on rotation %d,%d\n", t.id, tmpTile.id,x,y)
					tileLeftMatch = tmpTile.id
				}
				hf := tmpTile.FlipHorizontal()
				if hf.AdjacentHorizontal(tmpOrig) || hf.AdjacentHorizontal(tmpOrig.FlipHorizontal()) {
					//fmt.Printf("Tile %d Found match 2 with tile %d on rotation %d,%d\n", t.id, tmpTile.id,x,y)
					tileLeftMatch = tmpTile.id
				}
				vf := tmpTile.FlipVertical()
				if vf.AdjacentHorizontal(tmpOrig) || vf.AdjacentHorizontal(tmpOrig.FlipHorizontal()) {
					//fmt.Printf("Tile %d Found match 3 with tile %d on rotation %d,%d\n", t.id, tmpTile.id,x,y)
					tileLeftMatch = tmpTile.id
				}

				// find upper match
				if tmpTile.AdjacentVertical(tmpOrig) || tmpTile.AdjacentVertical(tmpOrig.FlipVertical()) {
					//fmt.Printf("Tile %d Found match 1 with tile %d on rotation %d,%d\n", t.id, tmpTile.id,x,y)
					tileUpperMatch = tmpTile.id
				}
				hf = tmpTile.FlipHorizontal()
				if hf.AdjacentVertical(tmpOrig) || hf.AdjacentVertical(tmpOrig.FlipVertical()) {
					tileUpperMatch = tmpTile.id
				}
				vf = tmpTile.FlipVertical()
				if vf.AdjacentVertical(tmpOrig) || vf.AdjacentVertical(tmpOrig.FlipVertical()) {
					tileUpperMatch = tmpTile.id
				}
				tmpOrig = tmpOrig.Rotate()
			}
			tmpTile = tmpTile.Rotate()
		}
		if tileLeftMatch > 0 && tileUpperMatch > 0 {
			return false
		}
	}
	return true
}

func (t Tile) AdjacentHorizontal(n Tile) bool {
	dimX := len(t.data[0]) - 1
	for y := 0; y < len(t.data); y++ {
		if t.data[y][dimX] != n.data[y][0] {
			return false
		}
	}
	return true
}

func (t Tile) AdjacentVertical(n Tile) bool {
	dimY := len(t.data) - 1
	for x := 0; x < len(t.data[0]); x++ {
		if t.data[dimY][x] != n.data[0][x] {
			return false
		}
	}
	return true
}

func problem(data []string) (p1answer int, p2answer int) {

	inputTiles := readTiles(data)

	allTiles := make(map[int][]Tile)

	// part 1
	for _, tile := range inputTiles {
		allTiles[tile.id] = tile.Variants()
	}
	grid := assemble(allTiles)
	p1answer = grid[0][0].id * grid[0][11].id * grid[11][0].id * grid[11][11].id

	// part 2
	p2Grid := makePartTwoGrid(grid)
	p2Grids := p2Grid.Variants()
	for _, p2grid := range p2Grids {
		if hasMonsters(p2grid) {
			p2answer = countNonMonster(p2grid)
		}
	}

	return p1answer, p2answer
}

func isMonster(t Tile, x int, y int) bool {
	if t.data[y-1][x] &&
		t.data[y][x+1] &&
		t.data[y][x+4] &&
		t.data[y-1][x+5] &&
		t.data[y-1][x+6] &&
		t.data[y][x+7] &&
		t.data[y][x+10] &&
		t.data[y-1][x+11] &&
		t.data[y-1][x+12] &&
		t.data[y][x+13] &&
		t.data[y][x+16] &&
		t.data[y-1][x+17] &&
		t.data[y-1][x+18] &&
		t.data[y-2][x+18] &&
		t.data[y-1][x+19] {
		return true
	}
	return false
}

func countNonMonster(t Tile) (sum int) {
	for y := 2; y < 96; y++ {
		for x := 0; x < 76; x++ {
			if isMonster(t, x, y) {
				t.data[y-1][x] = false
				t.data[y][x+1] = false
				t.data[y][x+4] = false
				t.data[y-1][x+5] = false
				t.data[y-1][x+6] = false
				t.data[y][x+7] = false
				t.data[y][x+10] = false
				t.data[y-1][x+11] = false
				t.data[y-1][x+12] = false
				t.data[y][x+13] = false
				t.data[y][x+16] = false
				t.data[y-1][x+17] = false
				t.data[y-1][x+18] = false
				t.data[y-2][x+18] = false
				t.data[y-1][x+19] = false
			}
		}
	}
	for y := 0; y < 96; y++ {
		for x := 0; x < 96; x++ {
			if t.data[y][x] {
				sum++
			}
		}
	}
	return sum
}

func makePartTwoGrid(g [][]Tile) Tile {

	grid := make([][]bool, 0)
	for i := 0; i < 12*8; i++ {
		grid = append(grid, make([]bool, 12*8))
	}

	for y := 0; y < 12; y++ {
		for x := 0; x < 12; x++ {
			gdata := g[y][x]
			for i := 1; i < 9; i++ {
				for j := 1; j < 9; j++ {
					if gdata.data[i][j] {
						grid[(y*8)+(i-1)][(x*8)+(j-1)] = true
					}
				}
			}
		}
	}

	return Tile{
		data: grid,
		id:   0,
	}
}

func hasMonsters(t Tile) bool {
	for y := 2; y < 96; y++ {
		for x := 0; x < 76; x++ {
			if isMonster(t, x, y) {
				return true
			}
		}
	}
	return false
}

func assemble(tiles map[int][]Tile) [][]Tile {
	for tileid, _ := range tiles {
		for _, tile := range tiles[tileid] {
			grid := make([][]Tile, 0)
			for x := 0; x < 12; x++ {
				grid = append(grid, make([]Tile, 12))
			}
			grid[0][0] = tile
			takenTiles := []int{tileid}
			success, g := canMatch(grid, helper.Coord{0, 0}, takenTiles, tiles)
			if success {
				return g
			}
		}
	}
	return nil
}

func canMatch(grid [][]Tile, c helper.Coord, takenTiles []int, allTiles map[int][]Tile) (bool, [][]Tile) {

	//fmt.Printf("Calling canMatch %v %v\n", c, takenTiles)
	if len(takenTiles) == len(allTiles) {
		return true, grid
	}

	current := helper.Coord{X: c.X, Y: c.Y}
	if c.X == 11 {
		current.X = 0
		current.Y++
	} else {
		current.X++
	}

	for tileid, _ := range allTiles {
		if helper.ContainsInt(tileid, takenTiles) {
			continue
		}

		for _, tile := range allTiles[tileid] {
			if current.X > 0 {
				// horizontal match with left
				lastTile := grid[c.Y][c.X]
				if lastTile.AdjacentHorizontal(tile) {
					grid[current.Y][current.X] = tile
					takenTiles = append(takenTiles, tile.id)
					return canMatch(grid, current, takenTiles, allTiles)
				}
			} else {
				// vertical match with above
				lastTile := grid[c.Y][0]
				if lastTile.AdjacentVertical(tile) {
					grid[current.Y][current.X] = tile
					takenTiles = append(takenTiles, tile.id)
					return canMatch(grid, current, takenTiles, allTiles)
				}
			}
		}
	}
	return false, nil
}

func main() {
	fh, _ := os.Open("input.txt")
	data, _ := helper.ReadLines(fh, false)
	p1, p2 := problem(data)
	fmt.Printf("Part one: %d\n", p1)
	fmt.Printf("Part two: %d\n", p2)

}
