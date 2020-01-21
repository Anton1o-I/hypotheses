package experiments

import (
	"fmt"
	"math"
)

// Experimenter is an interface that defines calculations to evaluate an experiment.
type Experimenter interface {
	Evaluate(impute bool) error
}

// Leveler defines behavior related to factor and block levels.
type Leveler interface {
	findLevel(id int) (*Level, error)
}

// Factor stores information about factors of an experiment .
type Factor struct {
	ID          int64
	Name        string
	Levels      []Level
	Description string
}

// NewFactor creates a new experiment Factor with specified levels
func NewFactor(n string, d string, l []Level) Factor {
	return Factor{
		Name:        n,
		Levels:      l,
		Description: d,
	}
}

// Level contains information related to levels of a factor.
type Level struct {
	ID          int
	Name        string
	Description string
	Stats       Statistics
}

// updateLevel updates the Treatment.
func (l *Level) updateLevel(v float64) {
	l.Stats.N++
	l.Stats.Sum += v
	l.Stats.SumSquared += math.Pow(v, 2)
}

// NewLevel creates a Level object for a Factor
func NewLevel(l string, d string) Level {
	return Level{
		Name:        l,
		Description: d,
	}
}

// Block stores information on blocks for experiments that use blocking.
type Block struct {
	ID     int
	Levels []Level
	Stats  Statistics
}

// findLevel returns a pointer to a level in a factor by matching to an id.
func (f *Factor) findLevel(id int) (*Level, error) {
	for i := range f.Levels {
		if f.Levels[i].ID == id {
			return &f.Levels[i], nil
		}
	}
	return nil, fmt.Errorf("Unable to find block with id = %d", id)
}

// findLevel returns a pointer to level in a block by matching an id.
func (b *Block) findLevel(id int) (*Level, error) {
	for i := range b.Levels {
		if b.Levels[i].ID == id {
			return &b.Levels[i], nil
		}
	}
	return nil, fmt.Errorf("Unable to find block with id = %d", id)
}

// Statistics stores summary statistics information
type Statistics struct {
	Sum        float64
	SumSquared float64
	N          int64
	Mean       float64
}

// updateStats updates the values stores in stats given a new observation.
func (s *Statistics) updateStats(v float64) {
	s.N++
	s.Sum += v
	s.SumSquared += math.Pow(v, 2)
}

// calcMean calculates the mean value of a statistic.
func (s *Statistics) calcMean() {
	s.Mean = s.Sum / float64(s.N)
}

// SourceOfVariation stores summary information for a source of variation
type SourceOfVariation struct {
	sumOfSquares float64
	meanSS       float64
	df           float64
	fStat        float64
	pValue       float64
}

// Subject stores information on a subject in an experiment.
type Subject struct {
	id       int64
	factors  []Factor
	blocks   []Block
	open     bool //this field specifies whether the subject has been assigned yet.
	value    float64
	observed bool
}

// newRCBDObservation returns an Observation with a hydrated subject but empty value.
func newSubject(id int64, f []Factor, b []Block) Subject {
	return Subject{
		id:      id,
		factors: f,
		blocks:  b,
		open:    true,
	}
}
