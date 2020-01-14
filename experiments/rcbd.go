package experiments

import (
	"gonum.org/v1/gonum/stat/distuv"

	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// rcbdStat stores statistical information about experiment.
type rcbdStat struct {
	total   SourceOfVariation
	err     SourceOfVariation
	factor  SourceOfVariation
	block   SourceOfVariation
	overall Statistics
}

// RCBDExperiment stores all the information about an experiment.
type RCBDExperiment struct {
	id       int64
	factor   Factor
	block    Block
	subjects []Subject
	stats    rcbdStat
}

// NewRCBDExperiment generates a new experiment and generates the slice of subjects that can
// be randomly assigned as the experiment commences.
func NewRCBDExperiment(id int64, f Factor, b Block) RCBDExperiment {
	var a RCBDExperiment
	a.id = id
	a.factor = f
	a.block = b
	var idCounter int64 = 1
	for f := range a.factor.Levels {
		for b := range a.block.Levels {
			factor := Factor{
				ID:          a.factor.ID,
				Name:        a.factor.Name,
				Description: a.factor.Description,
				Levels: []Level{
					a.factor.Levels[f],
				},
			}
			block := Block{
				ID: a.block.ID,
				Levels: []Level{
					a.block.Levels[b],
				},
			}
			s := newSubject(idCounter, []Factor{factor}, []Block{block})
			a.subjects = append(a.subjects, s)
			idCounter++
		}
	}
	return a
}

// getOpen returns the list of available observations that meet the criteria.
func (re RCBDExperiment) getOpen() []Subject {
	var a []Subject
	for i := range re.subjects {
		if re.subjects[i].open {
			a = append(a, re.subjects[i])
		}
	}
	return a
}

// RandomID returns a random ID from those still available and sets the ID to filled.
func (re *RCBDExperiment) RandomID() int64 {
	avail := re.getOpen()
	rand.Seed(time.Now().Unix())
	sel := rand.Intn(len(avail))
	return re.subjects[sel].id
}

// findObs searches available subjects for a match by id and returns the observations.
func (re *RCBDExperiment) findSubject(id int64) (*Subject, error) {
	for i := range re.subjects {
		if re.subjects[i].id == id {
			return &re.subjects[i], nil
		}
	}
	return nil, fmt.Errorf("Unable to find a subject with id - %d", id)
}

// UpdateObs will update the value for a subject in the experiment.
func (re *RCBDExperiment) UpdateObs(obs Subject) error {
	o, err := re.findSubject(obs.id)
	if err != nil {
		return err
	}
	o.value = obs.value
	o.observed = true
	return nil
}

// BulkUpdate will update multiple subjects at once.
func (re *RCBDExperiment) BulkUpdate(obs []Subject) error {
	for i := range obs {
		if err := re.UpdateObs(obs[i]); err != nil {
			return err
		}
	}
	return nil
}

// updateSources updates the levels and overall values for the experiment using the observed values.
func (re *RCBDExperiment) updateSources() error {
	for i := range re.subjects {
		for b := range re.subjects[i].blocks {
			for bl := range re.subjects[i].blocks[b].Levels {
				block := &re.block
				level, err := block.findLevel(re.subjects[i].blocks[b].Levels[bl].ID)
				if err != nil {
					return err
				}
				level.Stats.updateStats(re.subjects[i].value)
			}
		}
		for f := range re.subjects[i].factors {
			for fl := range re.subjects[i].factors[f].Levels {
				factor := &re.factor
				level, err := factor.findLevel(re.subjects[i].factors[f].Levels[fl].ID)
				if err != nil {
					return err
				}
				level.Stats.updateStats(re.subjects[i].value)
			}
		}
		re.stats.overall.updateStats(re.subjects[i].value)
	}
	return nil
}

// calcDF finds the degrees of freedom for total, treatments, blocks and error.
func (re *RCBDExperiment) calcDF() {
	re.stats.total.df = float64(len(re.subjects)) - 1
	re.stats.block.df = float64(len(re.block.Levels)) - 1
	re.stats.factor.df = float64(len(re.factor.Levels)) - 1
	re.stats.err.df = re.stats.block.df * re.stats.factor.df
}

// calcSSBlock calculates the sum of squares for block effects.
func (re *RCBDExperiment) calcSSBlock() {
	nLevels := float64(len(re.factor.Levels))
	for i := range re.block.Levels {
		b := &re.block.Levels[i]
		b.Stats.calcMean()
		re.stats.block.sumOfSquares += math.Pow(b.Stats.Sum, 2)
	}
	re.stats.block.sumOfSquares = re.stats.block.sumOfSquares / nLevels
	re.stats.block.sumOfSquares -= math.Pow(re.stats.overall.Sum, 2) / float64(re.stats.overall.N)
}

// calcSSTreatments calculates the sum of squares for treatment effects.
func (re *RCBDExperiment) calcSSTreatments() {
	nBlocks := float64(len(re.block.Levels))
	for i := range re.factor.Levels {
		l := &re.factor.Levels[i]
		l.Stats.calcMean()
		re.stats.factor.sumOfSquares += math.Pow(l.Stats.Sum, 2)
	}
	re.stats.factor.sumOfSquares = re.stats.factor.sumOfSquares / nBlocks
	re.stats.factor.sumOfSquares -= math.Pow(re.stats.overall.Sum, 2) / float64(re.stats.overall.N)
}

// calcSS calculates the sum of squares for total, block, treatment and error sources of variation.
func (re *RCBDExperiment) calcSS() {
	re.calcSSBlock()
	re.calcSSTreatments()
	re.stats.total.sumOfSquares = re.stats.overall.SumSquared - (math.Pow(re.stats.overall.Sum, 2) / float64(re.stats.overall.N))
	re.stats.err.sumOfSquares = re.stats.total.sumOfSquares - re.stats.block.sumOfSquares - re.stats.factor.sumOfSquares
}

// calcMS calculates the mean sum of squares for the experiment
func (re *RCBDExperiment) calcMS() {
	re.stats.err.meanSS = re.stats.err.sumOfSquares / re.stats.err.df
	re.stats.block.meanSS = re.stats.block.sumOfSquares / re.stats.block.df
	re.stats.factor.meanSS = re.stats.factor.sumOfSquares / re.stats.factor.df
}

// calcFStat finds the f-statistic value for the treatment in the RCBD experiment.
func (re *RCBDExperiment) calcFStat() {
	re.stats.factor.fStat = re.stats.factor.meanSS / re.stats.err.meanSS
}

// calcPValue returns the p-value for the treatment in an RCBD experiment.
func (re *RCBDExperiment) calcPValue() {
	dist := distuv.F{
		D1: re.stats.factor.df,
		D2: re.stats.err.df,
	}
	re.stats.factor.pValue = 1 - dist.CDF(re.stats.factor.fStat)
}

// Evaluate calculates the statistics for the experiment and stores them in the stats field.
// If impute is set to true, evaluate will impute missing values, if it is set to false,
// the evaluation will not be completed.
func (re *RCBDExperiment) Evaluate(impute bool) error {
	if re.missingValues() && !impute {
		return errors.New("Not all subjects have an observed value")
	}
	if err := re.updateSources(); err != nil {
		return err
	}
	re.stats.overall.calcMean()
	re.calcDF()
	re.calcSS()
	re.calcMS()

	re.calcFStat()

	re.calcPValue()
	return nil
}

// missingValues determines whether all subjects have an observed value.
func (re *RCBDExperiment) missingValues() bool {
	for i := range re.subjects {
		if re.subjects[i].observed == false {
			return true
		}
	}
	return false
}
