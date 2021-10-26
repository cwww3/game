package core

import (
	"fmt"
	"sync"
)

type Grid struct {
	sync.RWMutex
	Gid       int
	MinX      int
	MaxX      int
	MinY      int
	MaxY      int
	playerIDs map[int]bool
}

func NewGrid(gID, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		Gid:       gID,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		playerIDs: make(map[int]bool),
	}
}

func (g *Grid) Add(playerID int) {
	g.Lock()
	defer g.Unlock()
	g.playerIDs[playerID] = true
}

func (g *Grid) Remove(playerID int) {
	g.Lock()
	defer g.Unlock()
	delete(g.playerIDs, playerID)
}

func (g *Grid) GetPlayerIds() (playerIDs []int) {
	g.RLock()
	defer g.RUnlock()
	playerIDs = make([]int, 0, len(g.playerIDs))
	for playerID, _ := range g.playerIDs {
		playerIDs = append(playerIDs, playerID)
	}
	return playerIDs
}

func (g *Grid) String() string {
	g.RLock()
	defer g.RUnlock()
	return fmt.Sprintf("Grid id:%d minX:%d maxX:%d minY:%d maxY:%d players:%v\n", g.Gid, g.MinX, g.MaxX, g.MinY, g.MaxY, g.playerIDs)
}
