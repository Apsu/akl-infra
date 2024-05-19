package mini

type Ngram struct {
	chars []rune
	freq  float64
}

type Metric int

const (
	Sfb Metric = iota
	Sft
	Sfr
	Alt
	AltSfs
	Red
	BadRed
	RedSfs
	BadRedSfs
	InOne
	OutOne
	InRoll
	OutRoll
	Unknown
)

const MetricNum = 14

var Table [4096]Metric

var mapStrMetric = map[string]Metric{
	"sfb":         Sfb,
	"sft":         Sft,
	"alt":         Alt,
	"alt-sfs":     AltSfs,
	"red":         Red,
	"red-sfs":     RedSfs,
	"bad-red":     BadRed,
	"bad-red-sfs": BadRedSfs,
	"inoneh":      InOne,
	"outoneh":     OutOne,
	"inroll":      InRoll,
	"outroll":     OutRoll,
	"unknown":     Unknown,
}

var mapStrFinger = map[string]uint16{
	"LP": 0,
	"LR": 1,
	"LM": 2,
	"LI": 3,
	"LT": 4,
	"RT": 5,
	"RI": 6,
	"RM": 7,
	"RR": 8,
	"RP": 9,
}

var mapMetricStr = map[Metric]string{
	Sfb:       "sfb",
	Sft:       "sft",
	Alt:       "alt",
	AltSfs:    "alt-sfs",
	Red:       "red",
	BadRed:    "bad-red",
	RedSfs:    "red-sfs",
	BadRedSfs: "bad-red-sfs",
	InOne:     "inoneh",
	OutOne:    "outoneh",
	InRoll:    "inroll",
	OutRoll:   "outroll",
	Unknown:   "unknown",
}

var mapFingerStr = map[uint16]string{
	0: "LP",
	1: "LR",
	2: "LM",
	3: "LI",
	4: "LT",
	5: "RT",
	6: "RI",
	7: "RM",
	8: "RR",
	9: "RP",
}
