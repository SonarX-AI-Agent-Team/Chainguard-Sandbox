package sandbox

// Engine is a minimal consensus engine stub.
type Engine struct {
	Name   string
	Period uint64
}

// DefaultEngine returns a deterministic configuration.
func DefaultEngine() Engine {
	return Engine{
		Name:   "sandbox",
		Period: 3,
	}
}
