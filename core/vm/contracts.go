package vm

// Precompile is a minimal synthetic contract descriptor.
type Precompile struct {
	Name    string
	Address string
}

// DefaultPrecompiles returns a small placeholder set.
func DefaultPrecompiles() []Precompile {
	return []Precompile{
		{Name: "identity", Address: "0x0000000000000000000000000000000000000004"},
		{Name: "sha256", Address: "0x0000000000000000000000000000000000000002"},
	}
}
