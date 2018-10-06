package go_monplug_utils

import (
	"math"
	"strconv"
	"strings"
)

type perfdataStatus uint8

const ok perfdataStatus = 0
const warning perfdataStatus = 1
const critical perfdataStatus = 2

var posInf = math.Inf(1)
var negInf = math.Inf(-1)

type PerfdataCollection []Perfdata

func (self PerfdataCollection) calcStatus() perfdataStatus {
	status := ok

	for _, part := range self {
		if partStatus := part.calcStatus(); partStatus > status {
			status = partStatus
		}
	}

	return status
}

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

func (self *Perfdata) calcStatus() perfdataStatus {
	if self.Crit.contains(self.Value) {
		return critical
	}

	if self.Warn.contains(self.Value) {
		return warning
	}

	return ok
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

func (self *OptionalThreshold) contains(value float64) bool {
	if self.IsSet {
		return (self.Start <= value && value <= self.End) == self.Inverted
	}

	return false
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
