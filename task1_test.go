package mrctech

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// left out
var p11 = P{1, 1}

// right out
var p56 = P{5, 6}

// double peak
var p34 = P{3, 4}

var p12 = P{1, 2}
var p13 = P{1, 3}
var p14 = P{1, 4}
var p15 = P{1, 5}
var p16 = P{1, 6}
var p17 = P{1, 7}
var p18 = P{1, 8}
var p19 = P{1, 9}

var p23 = P{2, 3}
var p24 = P{2, 4}
var p25 = P{2, 5}
var p26 = P{2, 6}
var p27 = P{2, 7}
var p28 = P{2, 8}
var p29 = P{2, 9}

func TestDurationPeak(t *testing.T) {
	assert.Equal(t, 3, DurationPeak([]P{p11, p12, p13, p14, p15}, 2, 4, true))
	assert.Equal(t, 3, DurationPeak([]P{p12, p13, p14, p15, p56}, 2, 4, true))
	assert.Equal(t, 3, DurationPeak([]P{p11, p12, p13, p14, p15, p56}, 2, 4, true))
	assert.Equal(t, 3, DurationPeak([]P{p12, p13, p14, p15}, 2, 4, true))
	assert.Equal(t, 6, DurationPeak([]P{p11, p12, p13, p14, p15, p23, p24, p25}, 2, 4, true))
	assert.Equal(t, 3, DurationPeak([]P{p12, p13, p14, p15, p34}, 2, 4, true))
	assert.Equal(t, 1, DurationPeak([]P{p34}, 2, 4, true))
}
