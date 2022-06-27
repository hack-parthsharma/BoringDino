package dinosaur

import (
  "time"
  "math/rand"
)

type State struct{}
type StateType int

func (t StateType) String() string {
  return [...]string{"Walking", "Shopping", "Eating", "Working"}[t-1]
}

const (
  StateWalking StateType = iota + 1
  StateShopping
  StateEating
  StateWorking
)

func NewState() State {
  return State{}
}

func (s State) Now() StateType {
  rand.Seed(time.Now().UnixNano())
  return StateType(rand.Intn(int(StateWorking))+1)
}

type StateBlock struct {
  State StateType
  StartedAt time.Time
  Prev *StateBlock
}

func NewStateBlock() *StateBlock {
  return &StateBlock{
    Prev: nil,
  }
}

func (l *StateBlock) Update(s StateType, started time.Time) {
  l.StartedAt = started
  l.State = s
}
