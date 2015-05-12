// Generated by: gen
// TypeWriter: slice
// Directive: +gen on MyType

package main

// MyTypeSlice is a slice of type MyType. Use it where you would use []MyType.
type MyTypeSlice []MyType

// Where returns a new MyTypeSlice whose elements return true for func. See: http://clipperhouse.github.io/gen/#Where
func (rcv MyTypeSlice) Where(fn func(MyType) bool) (result MyTypeSlice) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Count gives the number elements of MyTypeSlice that return true for the passed func. See: http://clipperhouse.github.io/gen/#Count
func (rcv MyTypeSlice) Count(fn func(MyType) bool) (result int) {
	for _, v := range rcv {
		if fn(v) {
			result++
		}
	}
	return
}

// GroupByString groups elements into a map keyed by string. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv MyTypeSlice) GroupByString(fn func(MyType) string) map[string]MyTypeSlice {
	result := make(map[string]MyTypeSlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// GroupByInt64 groups elements into a map keyed by int64. See: http://clipperhouse.github.io/gen/#GroupBy
func (rcv MyTypeSlice) GroupByInt64(fn func(MyType) int64) map[int64]MyTypeSlice {
	result := make(map[int64]MyTypeSlice)
	for _, v := range rcv {
		key := fn(v)
		result[key] = append(result[key], v)
	}
	return result
}