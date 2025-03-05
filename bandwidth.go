package bandwidth

import "fmt"

type Bandwidth float64

const (
	BitPerSecond     Bandwidth = 1
	KilobitPerSecond Bandwidth = 1000 * BitPerSecond
	MegabitPerSecond Bandwidth = 1000 * KilobitPerSecond
	GigabitPerSecond Bandwidth = 1000 * MegabitPerSecond

	BytePerSecond     Bandwidth = 8 * BitPerSecond
	KilobytePerSecond Bandwidth = 1000 * BytePerSecond
	MegabytePerSecond Bandwidth = 1000 * KilobytePerSecond
	GigabytePerSecond Bandwidth = 1000 * MegabytePerSecond
)

func New(value float64, unit Bandwidth) Bandwidth {
	return Bandwidth(value * float64(unit))
}

func (bandwidth Bandwidth) To(unit Bandwidth) float64 {
	return float64(bandwidth) / float64(unit)
}

func (bandwidth Bandwidth) String(unit Bandwidth) string {
	return fmt.Sprintf("%.2f %s", bandwidth.To(unit), unit.UnitString())
}

func (bandwidth Bandwidth) UnitString() string {
	switch bandwidth {
	case BitPerSecond:
		return "b/s"
	case KilobitPerSecond:
		return "Kb/s"
	case MegabitPerSecond:
		return "Mb/s"
	case GigabitPerSecond:
		return "Gb/s"
	case BytePerSecond:
		return "B/s"
	case KilobytePerSecond:
		return "KB/s"
	case MegabytePerSecond:
		return "MB/s"
	case GigabytePerSecond:
		return "GB/s"
	default:
		return "?/?"
	}
}
