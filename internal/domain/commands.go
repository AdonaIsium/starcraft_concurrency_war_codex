package domain

import "time"

// CommandType enumerates available commands.
type CommandType string

const (
    CommandMove  CommandType = "MOVE"
    CommandAttack CommandType = "ATTACK"
    CommandScout  CommandType = "SCOUT"
)

// Command targets a squad with an action.
// Concurrency note:
// - Commands should be sent via a channel to the engine.
// - Engine processes commands at the start of each tick.
type Command struct {
    ID        ID
    Type      CommandType
    IssuedAt  time.Time
    PlayerID  ID
    SquadID   ID
    // Params for different command types (use oneOf in validation)
    TargetPos *Position // for MOVE/SCOUT
    TargetSquad *ID     // for ATTACK
}

