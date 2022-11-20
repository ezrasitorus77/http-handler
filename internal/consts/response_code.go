package consts

const (
	// GENERAL
	RCNotFound            string = "404"
	RCBadRequest          string = "400"
	RCMethodNotAllowed    string = "405"
	RCInternalServerError string = "500"

	// SPECIFIC
	// status 400
	RCInvalidPathFormat     string = "090"
	RCInvalidAddRoute       string = "091"
	RCOverlappingPathFormat string = "092"
	RCInvalidParamType      string = "093"
)
