// Code generated by "stringer -type=MetricKind"; DO NOT EDIT.

package export

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CounterKind-0]
	_ = x[GaugeKind-1]
	_ = x[MeasureKind-2]
}

const _MetricKind_name = "CounterKindGaugeKindMeasureKind"

var _MetricKind_index = [...]uint8{0, 11, 20, 31}

func (i MetricKind) String() string {
	if i < 0 || i >= MetricKind(len(_MetricKind_index)-1) {
		return "MetricKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MetricKind_name[_MetricKind_index[i]:_MetricKind_index[i+1]]
}