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

// Creates a new Duration from the specified number of nanoseconds.
func DurationFromNanos(nanos U64) Duration {
	return DurationNew(nanos/U64(_NANOS_PER_SEC), U32(nanos%U64(_NANOS_PER_SEC)))
}

// Returns true if this Duration spans no time.
func (self Duration) IsZero() bool {
	return self.seconds == 0 && self.nanoseconds == 0
}

// Returns the number of whole seconds contained by this Duration.
// The returned value does not include the fractional (nanosecond) part of the duration, which can be obtained using subsec_nanos.
func (self Duration) AsSecs() U64 {
	return self.seconds
}

// Returns the fractional part of this Duration, in whole milliseconds.
// This method does not return the length of the duration when represented by milliseconds.
// The returned number always represents a fractional portion of a second (i.e., it is less than one thousand).
func (self Duration) SubsecMillis() U32 {
	return self.nanoseconds / _NANOS_PER_MILLI
}

// Returns the fractional part of this Duration, in whole microseconds.
// This method does not return the length of the duration when represented by microseconds.
// The returned number always represents a fractional portion of a second (i.e., it is less than one million).
func (self Duration) SubsecMicros() U32 {
	return self.nanoseconds / _NANOS_PER_MICRO
}

// Returns the fractional part of this Duration, in nanoseconds.
// This method does not return the length of the duration when represented by nanoseconds.
// The returned number always represents a fractional portion of a second (i.e., it is less than one billion).
func (self Duration) SubsecNanos() U32 {
	return self.nanoseconds
}

// Returns the total number of whole milliseconds contained by this Duration.
func (self Duration) AsMillis() U128 {
	return U128_FromU64(self.seconds).Mul(U128_FromU64(U64(_MILLIS_PER_SEC))).Add(U128_FromU64(U64(self.nanoseconds / _NANOS_PER_MILLI)))
}
