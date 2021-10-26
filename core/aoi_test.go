package core

import (
	"fmt"
	"testing"
)

func TestAOIManager(t *testing.T) {
	aoi := NewAOIManager(0, 250, 0, 250, 5, 5)
	fmt.Println(aoi.gridWidth())
	fmt.Println(aoi.gridLength())
	fmt.Println(aoi)
}

func TestAOIManagerGetSurroundGrids(t *testing.T) {
	aoi := NewAOIManager(0, 250, 0, 250, 5, 5)
	for _, grid := range aoi.Grids {
		fmt.Println(grid.Gid,"----",aoi.GetSurroundGridsByGid(grid.Gid))
	}
}
