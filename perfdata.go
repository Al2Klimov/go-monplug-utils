package go_monplug_utils

import (
	"math"
	"strconv"
	"strings"
)

var posInf = math.Inf(1)
var negInf = math.Inf(-1)

type PerfdataCollection []Perfdata

func (self PerfdataCollection) String() string {
	if len(self) < 1 {
		return ""
	}

	result := make([]string, len(self))
	for i, perfdat := range self {
		result[i] = perfdat.String()
	}

	return " |" + strings.Join(result, " ")
}

type Perfdata struct {
	Label, UOM string
	Value      float64
	Warn, Crit OptionalThreshold
	Min, Max   OptionalNumber
}

func (self *Perfdata) String() string {
	builder := strings.Builder{}

	builder.WriteByte('\'')
	builder.WriteString(self.Label)
	builder.WriteByte('\'')

	builder.WriteByte('=')

	builder.WriteString(perfFloat(self.Value))
	builder.WriteString(self.UOM)

	return strings.TrimRight(
		strings.Join(
			[]string{builder.String(), self.Warn.String(), self.Crit.String(), self.Min.String(), self.Max.String()},
			";",
		),
		";",
	)
}

type OptionalThreshold struct {
	IsSet, Inverted bool
	Start, End      float64
}

func (self *OptionalThreshold) String() string {
	if self.IsSet {
		builder := strings.Builder{}

		if self.Inverted {
			builder.WriteByte('@')
		}

		if self.Start == 0 {
			if self.End == posInf {
				builder.WriteString("0:")
			} else {
				builder.WriteString(perfFloat(self.End))
			}
		} else {
			if self.Start == negInf {
				builder.WriteByte('~')
			} else {
				builder.WriteString(perfFloat(self.Start))
			}

			builder.WriteByte(':')

			if self.End != posInf {
				builder.WriteString(perfFloat(self.End))
			}
		}

		return builder.String()
	}

	return ""
}

type OptionalNumber struct {
	IsSet bool
	Value float64
}

func (self *OptionalNumber) String() string {
	if self.IsSet {
		return perfFloat(self.Value)
	}

	return ""
}

func perfFloat(x float64) string {
	if math.IsNaN(x) {
		x = 0
	} else if math.IsInf(x, 0) {
		if x > 0 {
			x = math.MaxFloat64
		} else {
			x = -math.MaxFloat64
		}
	}

	return strconv.FormatFloat(x, 'f', -1, 64)
}
