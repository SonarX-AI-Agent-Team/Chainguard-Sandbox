package ethapi

// API is a minimal RPC surface used only for synthetic validation.
type API struct {
	Network string
}

// Health returns a stable placeholder status.
func (api API) Health() map[string]string {
	return map[string]string{
		"network": api.Network,
		"status":  "ok",
	}
}
