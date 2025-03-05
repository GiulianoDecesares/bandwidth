package bandwidth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToConversion(context *testing.T) {
	testCases := []struct {
		name     string
		value    float64
		unit     Bandwidth
		toUnit   Bandwidth
		expected float64
	}{
		{
			name:     "500 B/s conversion to BytePerSecond",
			value:    500,
			unit:     BytePerSecond,
			toUnit:   BytePerSecond,
			expected: 500,
		},
		{
			name:     "10 KB/s conversion to KilobytePerSecond",
			value:    10,
			unit:     KilobytePerSecond,
			toUnit:   KilobytePerSecond,
			expected: 10,
		},
		{
			name:     "5 MB/s conversion to MegabytePerSecond",
			value:    5,
			unit:     MegabytePerSecond,
			toUnit:   MegabytePerSecond,
			expected: 5,
		},
		{
			name:     "2 Gb/s conversion to GigabitPerSecond",
			value:    2,
			unit:     GigabitPerSecond,
			toUnit:   GigabitPerSecond,
			expected: 2,
		},
	}

	for _, testCase := range testCases {
		context.Run(testCase.name, func(context *testing.T) {
			bandwidth := New(testCase.value, testCase.unit)
			actual := bandwidth.To(testCase.toUnit)
			assert.InDelta(context, testCase.expected, actual, 0.001, "Conversion mismatch")
		})
	}
}

func TestUnitString(context *testing.T) {
	testCases := []struct {
		unit     Bandwidth
		expected string
	}{
		{unit: BitPerSecond, expected: "b/s"},
		{unit: KilobitPerSecond, expected: "Kb/s"},
		{unit: MegabitPerSecond, expected: "Mb/s"},
		{unit: GigabitPerSecond, expected: "Gb/s"},
		{unit: BytePerSecond, expected: "B/s"},
		{unit: KilobytePerSecond, expected: "KB/s"},
		{unit: MegabytePerSecond, expected: "MB/s"},
		{unit: GigabytePerSecond, expected: "GB/s"},
		{unit: Bandwidth(123456), expected: "?/?"}, // Unknown unit test case
	}

	for _, testCase := range testCases {
		context.Run(testCase.expected, func(context *testing.T) {
			actual := testCase.unit.UnitString()
			assert.Equal(context, testCase.expected, actual, "UnitString mismatch")
		})
	}
}

func TestString(context *testing.T) {
	testCases := []struct {
		name     string
		value    float64
		unit     Bandwidth
		formatTo Bandwidth
		expected string
	}{
		{
			name:     "500 B/s",
			value:    500,
			unit:     BytePerSecond,
			formatTo: BytePerSecond,
			expected: "500.00 B/s",
		},
		{
			name:     "10 KB/s",
			value:    10,
			unit:     KilobytePerSecond,
			formatTo: KilobytePerSecond,
			expected: "10.00 KB/s",
		},
		{
			name:     "5 MB/s",
			value:    5,
			unit:     MegabytePerSecond,
			formatTo: MegabytePerSecond,
			expected: "5.00 MB/s",
		},
		{
			name:     "2 Gb/s",
			value:    2,
			unit:     GigabitPerSecond,
			formatTo: GigabitPerSecond,
			expected: "2.00 Gb/s",
		},
	}

	for _, testCase := range testCases {
		context.Run(testCase.name, func(context *testing.T) {
			bandwidth := New(testCase.value, testCase.unit)
			actual := bandwidth.String(testCase.formatTo)
			assert.Equal(context, testCase.expected, actual, "String representation mismatch")
		})
	}
}
