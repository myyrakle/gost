package gost

type Duration struct {
	seconds     U64
	nanoseconds U32
}

const _NANOS_PER_SEC U32 = 1_000_000_000
const _NANOS_PER_MILLI U32 = 1_000_000
const _NANOS_PER_MICRO U32 = 1_000
const _MILLIS_PER_SEC U32 = 1_000
const _MICROS_PER_SEC U32 = 1_000_000

// Creates a new Duration from the specified number of whole seconds and additional nanoseconds.
// If the number of nanoseconds is greater than 1 billion (the number of nanoseconds in a second), then it will carry over into the seconds provided.
func DurationNew(secs U64, nanos U32) Duration {
	seconds := U64(nanos / _NANOS_PER_SEC)

	return Duration{
		seconds:     secs.CheckedAdd(seconds).Expect("overflow in DurationNew"),
		nanoseconds: nanos % U32(1e9),
	}
}
