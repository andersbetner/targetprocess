package targetprocess

var (
	// ErrorCodes is a map of error number to string
	ErrorCodes = map[int]string{
		200: "Success. Request was handled correctly.",
		401: "Unauthorized. Wrong or missed credentials.",
		400: "Bad format. Incorrect parameter or query string.",
		403: "Forbidden. A user has insufficient rights to perform an action.",
		404: "Requested Entity not found.",
		500: "Internal server error. Targetprocess messed up.",
		501: "Not implemented. The requested action is either not supported or not implemented yet.",
	}
)
