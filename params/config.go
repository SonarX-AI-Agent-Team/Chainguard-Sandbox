package params

// ChainConfig holds minimal synthetic chain configuration.
type ChainConfig struct {
	ChainID      int
	NetworkName  string
	BlockGasLimit uint64
}

// DefaultConfig returns a stable placeholder configuration for tests.
func DefaultConfig() ChainConfig {
	return ChainConfig{
		ChainID:      7001,
		NetworkName:  "sandbox",
		BlockGasLimit: 30_000_000,
	}
}
