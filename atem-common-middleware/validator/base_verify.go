package validator

var (
	IdVerify = Rules{
		"ID": {NotEmpty()}}
	PageInfoVerify = Rules{
		"Page":     {NotEmpty()},
		"PageSize": {NotEmpty()}}
)
