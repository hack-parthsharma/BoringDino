package dinosaur

import (
  "math"
  "math/rand"
)

const (
  InitialWeight = 30.0
  InitialBalance = 10.0
  InitialSteps = 0
)

// Step change per second
func stepsChangeRate() float64 {
  return math.Max(float64(rand.Intn(6)), 1)
}

// Weight gaining per second by eating
func weightGainRate() float64 {
  return math.Max(rand.Float64(), 0.01)
}

// Weight loss per step
func weightLossRate() float64 {
  return 0.1
}
// Earning per second by working
func workingEarningRate() float64 {
  return math.Max(rand.Float64(), 0.01)
}

// Shopping spend per second
func shoppingSpendRate() float64 {
  return rand.Float64() / 10
}
