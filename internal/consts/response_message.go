package consts

const (
	// status 405
	NotAllowedMessage       string = "Method not allowed"
	MethodNotAllowedMessage string = "Requested path doesn't accept the method"

	// status 404
	NotFoundMessage     string = "Not found"
	PageNotFoundMessage string = "Page you are looking for is not found"

	// status 500
	InternalServerErrorMessage string = "Error occurred in server, please try again later"

	// status 400
	BadRequestMessage string = "Bad request"
	// complete this message with the sub path index
	InvalidParamTypeMessage string = "Invalid param type for sub path index "
)
