package status

const (
	CREATING_CODE = iota
	PENDING_CODE
	DELIVERING_CODE
	READY_CODE
	CANCELED_CODE

	CREATING   = "creating"
	PENDING    = "pending"
	DELIVERING = "delivering"
	READY      = "ready"
	CANCELED   = "canceled"
)

func StatusCodeToString(code int32) string {
	switch code {
	case CREATING_CODE:
		return CREATING
	case PENDING_CODE:
		return PENDING
	case DELIVERING_CODE:
		return DELIVERING
	case READY_CODE:
		return READY
	case CANCELED_CODE:
		return CANCELED
	default:
		return ""
	}
}

func StatusToCode(status string) int32 {
	switch status {
	case CREATING:
		return CREATING_CODE
	case PENDING:
		return PENDING_CODE
	case DELIVERING:
		return DELIVERING_CODE
	case READY:
		return READY_CODE
	case CANCELED:
		return CANCELED_CODE
	default:
		return -1
	}
}
