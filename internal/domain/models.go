package domain

import "time"

// ID type alias for clarity; you can switch to UUID strings.
type ID string

// Player represents a user in the game.
type Player struct {
    ID        ID
    Name      string
    CreatedAt time.Time
}

// Position is a grid coordinate.
type Position struct {
    X int
    Y int
}

// UnitType represents various unit archetypes (Marine, Zealot, etc.).
type UnitType string

const (
    UnitMarine UnitType = "Marine"
    UnitZealot UnitType = "Zealot"
)

// Unit models a single combatant.
type Unit struct {
    ID        ID
    SquadID   ID
    Type      UnitType
    HP        int
    Attack    int
    Range     int
    Vision    int
    Position  Position
    CreatedAt time.Time
}

// Squad is a group of units controlled by a player.
type Squad struct {
    ID        ID
    PlayerID  ID
    Name      string
    Units     []Unit
    Position  Position // representative/centroid position on the map
    CreatedAt time.Time
}

// Tile models a map tile (terrain, walkable, etc.).
type Tile struct {
    Walkable bool
}

// World holds the current game state.
type World struct {
    Width   int
    Height  int
    Tiles   [][]Tile
    Players map[ID]Player
    Squads  map[ID]Squad
    // Optional: Tick number for timing-based logic
    Tick uint64
}

// Snapshot contains a read-only view of the world for queries.
// Implementation detail:
// - Use deep copies or immutable structures to avoid data races.
type Snapshot struct {
    World World
}

