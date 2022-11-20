package consts

const (
	HSTS string = "Strict-Transport-Security"
	CSP  string = "Content-Security-Policy"
	CT   string = "Content-Type"

	JSON string = "application/json"
	HTML string = "text/html; charset=utf-8"

	// Default value
	// max-age serves in seconds
	DefaultHSTS string = "max-age=63072000; includeSubDomains; preload"
	DefaultCSP  string = "default-src 'self'; img-src https://*"

	AccesControlAllowMethods string = "Access-Control-Allow-Methods"
)
