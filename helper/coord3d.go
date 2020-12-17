package helper

import "fmt"

type Coord3D struct {
	X int
	Y int
	Z int
	V int
}

func (c Coord3D) AsString() string {
	return fmt.Sprintf("%d/%d/%d", c.X,c.Y,c.Z)
}