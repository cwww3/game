package core

import "fmt"

type AOIManager struct {
	MinX   int
	MaxX   int
	MinY   int
	MaxY   int
	CountX int
	CountY int
	Grids  map[int]*Grid
}

func NewAOIManager(minX, maxX, minY, maxY, countX, countY int) *AOIManager {
	aioManager := &AOIManager{
		MinX:   minX,
		MaxX:   maxX,
		MinY:   minY,
		MaxY:   maxY,
		CountX: countX,
		CountY: countY,
		Grids:  make(map[int]*Grid),
	}

	width := aioManager.gridWidth()
	length := aioManager.gridLength()
	for y := 0; y < countY; y++ {
		for x := 0; x < countX; x++ {
			gid := y*countX + x
			minX := x * width
			maxX := (x + 1) * width
			minY := y * length
			maxY := (y + 1) * length
			grid := NewGrid(gid, minX, maxX, minY, maxY)
			aioManager.Grids[gid] = grid
		}
	}
	return aioManager
}

func (m *AOIManager) gridWidth() int {
	return (m.MaxX - m.MinX) / m.CountX
}

func (m *AOIManager) gridLength() int {
	return (m.MaxY - m.MinY) / m.CountY
}

func (m *AOIManager) GetSurroundGridsByGid(gid int) []*Grid {
	if _, ok := m.Grids[gid]; !ok {
		return nil
	}

	grids := make([]*Grid, 0, 9)
	gids := make([]int, 0, 3)
	grids = append(grids, m.Grids[gid])
	gids = append(gids, gid)

	idx := gid % m.CountX
	if idx > 0 {
		grids = append(grids, m.Grids[gid-1])
		gids = append(gids, gid-1)
	}
	if idx < m.CountX-1 {
		grids = append(grids, m.Grids[gid+1])
		gids = append(gids, gid+1)
	}

	for _, v := range gids {
		idy := v / m.CountY
		if idy > 0 {
			grids = append(grids, m.Grids[v-m.CountX])
		}
		if idy < m.CountY-1 {
			grids = append(grids, m.Grids[v+m.CountX])
		}
	}

	return grids
}

func (m *AOIManager) GetGidByPos(x, y float32) int {
	if x < 0 || x > float32(m.MaxX) {
		return 0
	}
	if y < 0 || y > float32(m.MaxY) {
		return 0
	}
	width := m.gridWidth()
	length := m.gridLength()
	idx := int(x) / width * width
	idy := int(y) / length * length
	return idx + idy*m.CountX
}

func (m *AOIManager) GetSurroundPlayerIDsByPos(x, y float32) []int {
	gid := m.GetGidByPos(x, y)
	grids := m.GetSurroundGridsByGid(gid)
	playerIDs := make([]int, 0, 10)
	for _, grid := range grids {
		playerIDs = append(playerIDs, grid.GetPlayerIds()...)
	}
	return playerIDs
}

func (m *AOIManager) String() string {
	s := fmt.Sprintf("aio minX:%d maxX:%d minY:%d maxY:%d\n", m.MinX, m.MaxX, m.MinY, m.MaxY)
	for _, grid := range m.Grids {
		s += grid.String()
	}
	return s
}
