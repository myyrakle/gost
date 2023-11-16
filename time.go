package gost

// A Duration type to represent a span of time, typically used for system timeouts.
// Each Duration is composed of a whole number of seconds and a fractional part represented in nanoseconds. If the underlying system does not support nanosecond-level precision, APIs binding a system timeout will typically round up the number of nanoseconds.
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

// Creates a new Duration from the specified number of whole seconds.
func DurationFromSecs(secs U64) Duration {
	return DurationNew(secs, 0)
}

// Creates a new Duration from the specified number of milliseconds.
func DurationFromMillis(millis U64) Duration {
	return DurationNew(millis/U64(_MILLIS_PER_SEC), U32(millis%U64(_MILLIS_PER_SEC))*_NANOS_PER_MILLI)
}

// Creates a new Duration from the specified number of microseconds.
func DurationFromMicros(micros U64) Duration {
	return DurationNew(micros/U64(_MICROS_PER_SEC), U32(micros%U64(_MICROS_PER_SEC))*_NANOS_PER_MICRO)
}
