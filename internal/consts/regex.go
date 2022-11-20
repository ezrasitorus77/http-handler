package consts

const (
	ParamPrefixRegex string = "->d:|->s:|->f:"
	ParamKeyRegex    string = `\w+[^\:]`

	FullPathRegex string = `^\/$|^\/\w$|^\/\w[a-zA-Z0-9\-\/->]+$`
	SubPathRegex  string = `(^(->d:)|^(->s:)|^(->f:))(\w+$|[a-zA-Z0-9\-]+\w$)|^\-$|^(\w$)|^\w(\w+$|[a-zA-Z0-9\-]+\w$)`

	StringRegex string = `.*`
	IntRegex    string = `\d+$`
	FloatRegex  string = `^\d+c{1,}\d+$`
)
