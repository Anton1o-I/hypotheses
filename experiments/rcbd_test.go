package experiments

import (
	"math"
	"reflect"
	"testing"
)

var factorLevel1 = Level{
	ID:          1,
	Name:        "levelA",
	Description: "test levelA",
}

var factorLevel2 = Level{
	ID:          2,
	Name:        "levelB",
	Description: "test levelB",
}

var factorLevel3 = Level{
	ID:          3,
	Name:        "levelC",
	Description: "test levelC",
}

var blockLevel1 = Level{
	ID: 1,
}

var blockLevel2 = Level{
	ID: 2,
}
var blockLevel3 = Level{
	ID: 3,
}

var newExpFactors = Factor{
	ID:          1,
	Name:        "test",
	Description: "test factors",
	Levels: []Level{
		Level{
			ID:   1,
			Name: "level1",
		},
		Level{
			ID:   2,
			Name: "level2",
		},
	},
}

var newExpBlocks = Block{
	ID: 1,
	Levels: []Level{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
	},
}

var newExpSubjects = []Subject{
	Subject{
		id: 1,
		factors: []Factor{
			{
				ID:          1,
				Name:        "test",
				Description: "test factors",
				Levels: []Level{
					{
						ID:   1,
						Name: "level1",
					},
				},
			},
		},
		blocks: []Block{
			{
				ID: 1,
				Levels: []Level{
					{
						ID: 1,
					},
				},
			},
		},
		open: true,
	},
	Subject{
		id: 2,
		factors: []Factor{
			{
				ID:          1,
				Name:        "test",
				Description: "test factors",
				Levels: []Level{
					{
						ID:   1,
						Name: "level1",
					},
				},
			},
		},
		blocks: []Block{
			{
				ID: 1,
				Levels: []Level{
					{
						ID: 2,
					},
				},
			},
		},
		open: true,
	},
	Subject{
		id: 3,
		factors: []Factor{
			{
				ID:          1,
				Name:        "test",
				Description: "test factors",
				Levels: []Level{
					{
						ID:   2,
						Name: "level2",
					},
				},
			},
		},
		blocks: []Block{
			{
				ID: 1,
				Levels: []Level{
					{
						ID: 1,
					},
				},
			},
		},
		open: true,
	},
	Subject{
		id: 4,
		factors: []Factor{
			{
				ID:          1,
				Name:        "test",
				Description: "test factors",
				Levels: []Level{
					{
						ID:   2,
						Name: "level2",
					},
				},
			},
		},
		blocks: []Block{
			{
				ID: 1,
				Levels: []Level{
					{
						ID: 2,
					},
				},
			},
		},
		open: true,
	},
}

func TestNewRCBDExperiment(t *testing.T) {
	type args struct {
		id int64
		f  Factor
		b  Block
	}
	tests := []struct {
		name string
		args args
		want RCBDExperiment
	}{
		{
			name: "pass",
			args: args{
				id: 1,
				f:  newExpFactors,
				b:  newExpBlocks,
			},
			want: RCBDExperiment{
				id:     1,
				factor: newExpFactors,
				block: Block{
					ID: 1,
					Levels: []Level{
						{
							ID: 1,
						},
						{
							ID: 2,
						},
					},
				},
				subjects: newExpSubjects,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRCBDExperiment(tt.args.id, tt.args.f, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRCBDExperiment() = \n%+v, want \n%+v", got, tt.want)
			}
		})
	}
}

var expSubjects = []Subject{
	{
		id: 1,
		factors: []Factor{
			{
				ID: 1,
				Levels: []Level{
					factorLevel1,
				},
			},
		},
		blocks: []Block{
			{
				ID: 1,
				Levels: []Level{
					blockLevel1,
				},
			},
		},
		open:     false,
		value:    2,
		observed: true,
	},
	{
		id: 2,
		factors: []Factor{
			{
				ID: 1,
				Levels: []Level{
					factorLevel1,
				},
			},
		},
		blocks: []Block{
			{
				ID: 1,
				Levels: []Level{
					blockLevel2,
				},
			},
		},
		open:     false,
		value:    3,
		observed: true,
	},
	{
		id: 3,
		factors: []Factor{
			{
				ID: 1,
				Levels: []Level{
					factorLevel1,
				},
			},
		},
		blocks: []Block{
			{
				ID: 1,
				Levels: []Level{
					blockLevel3,
				},
			},
		},
		open:     false,
		value:    4,
		observed: true,
	},
	{
		id: 4,
		factors: []Factor{
			{
				ID: 1,
				Levels: []Level{
					factorLevel2,
				},
			},
		},
		blocks: []Block{
			{
				ID: 1,
				Levels: []Level{
					blockLevel1,
				},
			},
		},
		open:     false,
		value:    4,
		observed: true,
	},
	{
		id: 5,
		factors: []Factor{
			{
				ID: 1,
				Levels: []Level{
					factorLevel2,
				},
			},
		},
		blocks: []Block{
			{
				ID: 1,
				Levels: []Level{
					blockLevel2,
				},
			},
		},
		open:     false,
		value:    4,
		observed: true,
	},
	{
		id: 6,
		factors: []Factor{
			{
				ID: 1,
				Levels: []Level{
					factorLevel2,
				},
			},
		},
		blocks: []Block{
			{
				ID: 1,
				Levels: []Level{
					blockLevel3,
				},
			},
		},
		open:     false,
		value:    4,
		observed: true,
	},
	{
		id: 7,
		factors: []Factor{
			{
				ID: 1,
				Levels: []Level{
					factorLevel3,
				},
			},
		},
		blocks: []Block{
			{
				ID: 1,
				Levels: []Level{
					blockLevel1,
				},
			},
		},
		open:     false,
		value:    4,
		observed: true,
	},
	{
		id: 8,
		factors: []Factor{
			{
				ID: 1,
				Levels: []Level{
					factorLevel3,
				},
			},
		},
		blocks: []Block{
			{
				ID: 1,
				Levels: []Level{
					blockLevel2,
				},
			},
		},
		open:     false,
		value:    4,
		observed: true,
	},
	{
		id: 9,
		factors: []Factor{
			{
				ID: 1,
				Levels: []Level{
					factorLevel3,
				},
			},
		},
		blocks: []Block{
			{
				ID: 1,
				Levels: []Level{
					blockLevel3,
				},
			},
		},
		open:     false,
		value:    4,
		observed: true,
	},
}

var testExp = RCBDExperiment{
	id: 1,
	factor: Factor{
		ID:   1,
		Name: "test",
		Levels: []Level{
			factorLevel1,
			factorLevel2,
			factorLevel3,
		},
	},
	block: Block{
		ID: 1,
		Levels: []Level{
			blockLevel1,
			blockLevel2,
			blockLevel3,
		},
	},
	subjects: expSubjects,
}

var evaluateTestExp = RCBDExperiment{
	id: 1,
	factor: Factor{
		ID:   1,
		Name: "test",
		Levels: []Level{
			factorLevel1,
			factorLevel2,
			factorLevel3,
		},
	},
	block: Block{
		ID: 1,
		Levels: []Level{
			blockLevel1,
			blockLevel2,
			blockLevel3,
		},
	},
	subjects: expSubjects,
}

func updateTestExp() RCBDExperiment {
	a := RCBDExperiment{
		id: 1,
		factor: Factor{
			ID:   1,
			Name: "test",
			Levels: []Level{
				factorLevel1,
				factorLevel2,
				factorLevel3,
			},
		},
		block: Block{
			ID: 1,
			Levels: []Level{
				blockLevel1,
				blockLevel2,
				blockLevel3,
			},
		},
		subjects: expSubjects,
	}
	a.factor.Levels[0].Stats = Statistics{
		N:          3,
		Sum:        9,
		SumSquared: 29,
	}
	a.factor.Levels[1].Stats = Statistics{
		N:          3,
		Sum:        12,
		SumSquared: 48,
	}
	a.factor.Levels[2].Stats = Statistics{
		N:          3,
		Sum:        12,
		SumSquared: 48,
	}
	a.block.Levels[0].Stats = Statistics{
		N:          3,
		Sum:        10,
		SumSquared: 36,
	}
	a.block.Levels[1].Stats = Statistics{
		N:          3,
		Sum:        11,
		SumSquared: 41,
	}
	a.block.Levels[2].Stats = Statistics{
		N:          3,
		Sum:        12,
		SumSquared: 48,
	}
	a.stats.overall = Statistics{
		Sum:        33,
		SumSquared: 125,
		N:          9,
	}
	return a
}

func expWithSSDF() RCBDExperiment {
	a := updateTestExp()
	a.stats.total.df = 8
	a.stats.err.df = 4
	a.stats.factor.df = 2
	a.stats.block.df = 2

	a.stats.total.sumOfSquares = 4.0
	a.stats.err.sumOfSquares = 1.3333333333333286
	a.stats.factor.sumOfSquares = 2.0
	a.stats.block.sumOfSquares = 0.6666666666666714

	return a
}

func expWithMS() RCBDExperiment {
	a := expWithSSDF()
	a.stats.err.meanSS = 0.33333333333333215
	a.stats.factor.meanSS = 1
	a.stats.block.meanSS = 0.3333333333333357

	return a
}

func expWithF() RCBDExperiment {
	a := expWithMS()

	a.stats.factor.fStat = 1 / 0.33333333333333215
	return a
}

func expWithPValue() RCBDExperiment {
	a := expWithF()

	a.stats.factor.pValue = 0.15999999999999925
	for f := range a.factor.Levels {
		a.factor.Levels[f].Stats.calcMean()
	}
	for b := range a.block.Levels {
		a.block.Levels[b].Stats.calcMean()
	}
	a.stats.overall.calcMean()
	return a
}

func TestRCBDExperiment_updateSources(t *testing.T) {
	type fields struct {
		id       int64
		factor   Factor
		block    Block
		subjects []Subject
		stats    rcbdStat
	}
	tests := []struct {
		name    string
		exp     *RCBDExperiment
		wantErr bool
		want    RCBDExperiment
	}{
		{
			name:    "pass",
			exp:     &testExp,
			wantErr: false,
			want:    updateTestExp(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := tt.exp
			if err := re.updateSources(); (err != nil) != tt.wantErr {
				t.Errorf("RCBDExperiment.updateSources() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(re, &tt.want) {
					t.Errorf("UpdateObs() = \n%+v, want \n%+v", re, &tt.want)
				}
			}
		})
	}
}

func TestRCBDExperiment_calcDF(t *testing.T) {
	tests := []struct {
		name string
		exp  RCBDExperiment
		want rcbdStat
	}{
		{
			name: "testDF Pass",
			exp:  updateTestExp(),
			want: rcbdStat{
				total: SourceOfVariation{
					df: 8,
				},
				err: SourceOfVariation{
					df: 4,
				},
				factor: SourceOfVariation{
					df: 2,
				},
				block: SourceOfVariation{
					df: 2,
				},
				overall: Statistics{
					N:          9,
					Sum:        33,
					SumSquared: 125,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := tt.exp
			re.calcDF()
			if !reflect.DeepEqual(re.stats, tt.want) {
				t.Errorf("calcDF() = %v, want %v", re.stats, tt.want)
			}
		})
	}
}

func TestRCBDExperiment_calcSS(t *testing.T) {
	tests := []struct {
		name string
		exp  RCBDExperiment
		want rcbdStat
	}{
		{
			name: "testDF Pass",
			exp:  updateTestExp(),
			want: rcbdStat{
				overall: Statistics{
					Sum:        33,
					N:          9,
					Mean:       33.0 / 9,
					SumSquared: 125,
				},
				total: SourceOfVariation{
					sumOfSquares: 125 - (math.Pow(33, 2) / 9),
				},
				err: SourceOfVariation{
					sumOfSquares: 1.3333333333333286,
				},
				factor: SourceOfVariation{
					sumOfSquares: 2.0,
				},
				block: SourceOfVariation{
					sumOfSquares: 0.6666666666666714,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := tt.exp
			re.stats.overall.calcMean()
			re.calcSS()
			if !reflect.DeepEqual(re.stats, tt.want) {
				t.Errorf("calcSS() = %+v,want = %v", re.stats, tt.want)
			}
		})
	}
}

func TestRCBDExperiment_calcMS(t *testing.T) {
	tests := []struct {
		name string
		exp  RCBDExperiment
		want rcbdStat
	}{
		{
			name: "testDF Pass",
			exp:  expWithSSDF(),
			want: rcbdStat{
				overall: Statistics{
					Sum:        33,
					N:          9,
					SumSquared: 125,
				},

				total: SourceOfVariation{
					df:           8,
					sumOfSquares: 125 - (math.Pow(33, 2) / 9),
				},
				err: SourceOfVariation{
					sumOfSquares: 1.3333333333333286,
					df:           4,
					meanSS:       0.33333333333333215,
				},
				factor: SourceOfVariation{
					df:           2,
					sumOfSquares: 2,
					meanSS:       1,
				},
				block: SourceOfVariation{
					df:           2,
					sumOfSquares: 0.6666666666666714,
					meanSS:       0.3333333333333357,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := tt.exp
			re.calcMS()
			if !reflect.DeepEqual(re.stats, tt.want) {
				t.Errorf("calcSS() = \n%v,want = \n%v", re.stats, tt.want)
			}
		})
	}
}

func TestRCBDExperiment_calcFStat(t *testing.T) {
	tests := []struct {
		name string
		exp  RCBDExperiment
		want rcbdStat
	}{
		{
			name: "test FStat Pass",
			exp:  expWithMS(),
			want: rcbdStat{
				overall: Statistics{
					Sum:        33,
					N:          9,
					SumSquared: 125,
				},

				total: SourceOfVariation{
					df:           8,
					sumOfSquares: 125 - (math.Pow(33, 2) / 9),
				},
				err: SourceOfVariation{
					sumOfSquares: 1.3333333333333286,
					df:           4,
					meanSS:       0.33333333333333215,
				},
				factor: SourceOfVariation{
					df:           2,
					sumOfSquares: 2,
					meanSS:       1,
					fStat:        1.0 / 0.33333333333333215,
				},
				block: SourceOfVariation{
					df:           2,
					sumOfSquares: 0.6666666666666714,
					meanSS:       0.3333333333333357,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := tt.exp
			re.calcFStat()
			if !reflect.DeepEqual(re.stats, tt.want) {
				t.Errorf("calcFStat() = \n%+v,want = \n%+v", re.stats, tt.want)
			}
		})
	}
}

func TestRCBDExperiment_calcPValue(t *testing.T) {
	tests := []struct {
		name string
		exp  RCBDExperiment
		want rcbdStat
	}{
		{
			name: "testDF Pass",
			exp:  expWithF(),
			want: rcbdStat{
				overall: Statistics{
					Sum:        33,
					N:          9,
					SumSquared: 125,
				},

				total: SourceOfVariation{
					df:           8,
					sumOfSquares: 125 - (math.Pow(33, 2) / 9),
				},
				err: SourceOfVariation{
					sumOfSquares: 1.3333333333333286,
					df:           4,
					meanSS:       0.33333333333333215,
				},
				factor: SourceOfVariation{
					df:           2,
					sumOfSquares: 2,
					meanSS:       1,
					fStat:        1.0 / 0.33333333333333215,
					pValue:       0.15999999999999925,
				},
				block: SourceOfVariation{
					df:           2,
					sumOfSquares: 0.6666666666666714,
					meanSS:       0.3333333333333357,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := tt.exp
			re.calcPValue()
			if !reflect.DeepEqual(re.stats, tt.want) {
				t.Errorf("calcSS() = \n%v,want = \n%v", re.stats, tt.want)
			}
		})
	}
}

func TestRCBDExperiment_Evaluate(t *testing.T) {
	tests := []struct {
		name string
		exp  RCBDExperiment
		want RCBDExperiment
	}{
		{
			name: "testDF Pass",
			exp:  evaluateTestExp,
			want: expWithPValue(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := &tt.exp
			re.Evaluate(false)
			if !reflect.DeepEqual(re, &tt.want) {
				t.Errorf("calcDF() = \n%+v, want \n%+v", re, &tt.want)
			}
		})
	}
}
func TestRCBDExperiment_UpdateObs(t *testing.T) {
	type fields struct {
		id       int64
		factor   Factor
		block    Block
		subjects []Subject
		stats    rcbdStat
	}
	type args struct {
		obs Subject
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		want    []Subject
	}{
		{
			name: "good update",
			fields: fields{
				id: 1,
				subjects: []Subject{
					Subject{
						id:      1,
						factors: []Factor{},
						blocks:  []Block{},
					},
				},
			},
			args: args{
				obs: Subject{
					id:    1,
					value: 10,
				},
			},
			wantErr: false,
			want: []Subject{
				Subject{
					id:       1,
					value:    10,
					observed: true,
					factors:  []Factor{},
					blocks:   []Block{},
				},
			},
		},
		{
			name: "test error",
			fields: fields{
				id:     1,
				factor: Factor{},
				block:  Block{},
				subjects: []Subject{
					Subject{
						id: 1,
					},
				},
			},
			args: args{
				obs: Subject{
					id:    2,
					value: 10,
				},
			},
			wantErr: true,
			want:    []Subject{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			re := &RCBDExperiment{
				id:       tt.fields.id,
				factor:   tt.fields.factor,
				block:    tt.fields.block,
				subjects: tt.fields.subjects,
				stats:    tt.fields.stats,
			}
			if err := re.UpdateObs(tt.args.obs); (err != nil) != tt.wantErr {
				t.Errorf("RCBDExperiment.UpdateObs() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if !reflect.DeepEqual(re.subjects, tt.want) {
					t.Errorf("UpdateObs() = \n%v, want \n%v", re, tt.want)
				}
			}
		})
	}
}
