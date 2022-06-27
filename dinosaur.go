package dinosaur

import (
  "fmt"
  "time"
  "math/rand"
)


type Dinosaur struct {
  mood Mood
  state State
  stateChanged chan StateBlock
  stateUpdated chan bool
  currState *StateBlock
  wallet *Wallet
  health *Health
  stateHandler map[StateType]interface{}
  running bool
}

func NewDinosaur() *Dinosaur {
  rand.Seed(time.Now().UnixNano())
  ret := &Dinosaur{
    mood: NewMood(),
    wallet: NewWallet(),
    state: NewState(),
    running: true,
    currState: NewStateBlock(),
    health: NewHealth(),
  }

  ret.currState.Prev = NewStateBlock()
  ret.currState.Update(ret.state.Now(), time.Now())
  ret.currState.Prev.Update(ret.currState.State, ret.currState.StartedAt)

  ret.stateHandler = make(map[StateType]interface{})
  ret.stateHandler[StateWorking] = ret.workingHandler
  ret.stateHandler[StateShopping] = ret.shoppingHandler
  ret.stateHandler[StateWalking] = ret.walkingHandler
  ret.stateHandler[StateEating] = ret.eatingHandler

  ret.stateChanged = make(chan StateBlock, 1)
  ret.stateUpdated = make(chan bool, 1)

  go func() {
    ret.StateDispatch()
  }()
  return ret
}

func (d Dinosaur) Measure() {
  d.currState.Prev.Update(d.currState.State, d.currState.StartedAt)
  d.currState.Update(d.state.Now(), time.Now())
  d.stateChanged <- *d.currState
}

func (d Dinosaur) Mood() MoodType {
  return d.mood.Now()
}

func (d Dinosaur) String() string {
  return fmt.Sprintf("Little dinosaur looks %s while %s, balance: %.2f, steps: %d, weight: %.2f, now it's %s",
    d.Mood(), d.currState.Prev.State, d.Balance(), d.health.Steps, d.health.Weight, d.currState.State)
}

func (d Dinosaur) walkingHandler(curr StateBlock) {
  dur := time.Now().Sub(curr.Prev.StartedAt)
  steps := int(dur.Seconds() * stepsChangeRate())
  d.health.Steps += steps
  d.health.Weight -= float64(steps) * weightLossRate()
}

func (d Dinosaur) eatingHandler(curr StateBlock) {
  d.health.Weight += time.Now().Sub(curr.Prev.StartedAt).Seconds() * weightGainRate()
}

func (d Dinosaur) workingHandler(curr StateBlock) {
  d.wallet.Update(time.Now().Sub(curr.Prev.StartedAt).Seconds() * workingEarningRate())
}

func (d Dinosaur) shoppingHandler(curr StateBlock) {
  d.wallet.Update(-(d.wallet.Balance() * shoppingSpendRate()))
}

func (d Dinosaur) StateDispatch() {
  for {
    select {
      case last := <-d.stateChanged:
        if h, ok := d.stateHandler[last.Prev.State]; ok {
          h.(func(StateBlock))(last)
          d.stateUpdated <- true
        }
      case <- d.stateUpdated:
        fmt.Println(d)
    }
  }
}

func (d Dinosaur) Balance() float64 {
  return d.wallet.Balance()
}

func (d Dinosaur) Close() {
  d.running = false
}
