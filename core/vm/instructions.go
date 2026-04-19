package vm

// OpCode is a placeholder instruction identifier.
type OpCode byte

const (
	STOP OpCode = 0x00
	ADD  OpCode = 0x01
)

// ExecuteStep provides a minimal instruction executor stub.
func ExecuteStep(op OpCode) string {
	switch op {
	case STOP:
		return "halt"
	case ADD:
		return "add"
	default:
		return "unknown"
	}
}
