package verify

// String ...
type String string

// Empty ...
func (r String) Empty() bool {
	return len(string(r)) == 0
}
