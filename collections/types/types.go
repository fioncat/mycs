package types

import "bytes"

type ScoreKey struct {
	Key   string
	Score float64
}

type ScoreValue struct {
	Value interface{}
	Score float64
}

func ScoreValue2Key(values []ScoreValue) []ScoreKey {
	keys := make([]ScoreKey, len(values))
	for i, val := range values {
		keys[i] = ScoreKey{
			Key:   val.Value.(string),
			Score: val.Score,
		}
	}
	return keys
}

type Range struct {
	Min, Max float64

	Minex, Maxex bool
}

func (r *Range) IsValid() bool {
	if r.Min > r.Max {
		return false
	}
	if r.Minex || r.Maxex {
		return r.Max != r.Min
	}
	return true
}

func (r *Range) GteMin(value float64) bool {
	if r.Minex {
		return value > r.Min
	}
	return value >= r.Min
}

func (r *Range) LteMax(value float64) bool {
	if r.Maxex {
		return value < r.Max
	}
	return value <= r.Max
}

type Comparable interface {
	Compare(b interface{}) int
}

func Compare(a, b interface{}) int {
	switch va := a.(type) {
	case string:
		vb := b.(string)
		if va == vb {
			return 0
		}
		if va < vb {
			return -1
		}
		return 1

	case []byte:
		return bytes.Compare(va, b.([]byte))

	default:
		ca, ok := a.(Comparable)
		if !ok {
			panic("Compare: a must implement types.Comparable interface")
		}
		return ca.Compare(b)
	}
}
