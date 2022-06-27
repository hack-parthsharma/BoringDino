package dinosaur

import (
  "time"
  "math/rand"
)

type Mood struct{}
type MoodType int

const (
  MoodSunny MoodType = iota + 1
  MoodWindy
  MoodRainy
  MoodCloudy
)

func (t MoodType) String() string {
  return [...]string{"Sunny", "Windy", "Rainy", "Cloudy"}[t-1]
}

func NewMood() Mood {
  return Mood{}
}

func (m Mood) Now() MoodType {
  rand.Seed(time.Now().UnixNano())
  return MoodType(rand.Intn(int(MoodCloudy))+1)
}
