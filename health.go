package dinosaur


type Health struct {
  Steps int
  Weight float64
}


func NewHealth() *Health {
  return &Health {
    Weight: InitialWeight,
    Steps: InitialSteps,
  }
}
