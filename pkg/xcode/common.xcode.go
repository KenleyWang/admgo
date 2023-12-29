package xcode

var (
	OK                 = add(0, "ok")
	NoLogin            = add(101, "no login")
	RequestErr         = add(400, "INVALID_ARGUMENT")
	Unauthorized       = add(401, "UNAUTHENTICATED")
	AccessDenied       = add(403, "PERMISSION_DENIED")
	NotFound           = add(404, "not found")
	MethodNotAllowed   = add(405, "METHOD_NOT_ALLOWED")
	Canceled           = add(498, "Canceled")
	ServerErr          = add(500, "INTERNAL_ERROR")
	ServiceUnavailable = add(503, "UNAVAILABLE")
	Deadline           = add(504, "DEADLINE_EXCEEDED")
	LimitExceed        = add(509, "RESOURCE_EXHAUSTED")
	BadRequests        = add(40001, "bad requests")
)

var (
	BadRequestsCode = func(err error) error { return New(40001, err.Error()) }
)
