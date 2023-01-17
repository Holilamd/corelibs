package corelibs

type OutputFormat struct {
	Success bool
	Message string
	Data    interface{}
	Errors  []struct {
		ErrorFormat
	}
	Code string
}

type ErrorFormat struct {
	Field string
	Error string
}

func ResponseFormat(success bool, message string, data interface{}, errors []struct{ ErrorFormat }) *OutputFormat {
	response := new(OutputFormat)
	response.Success = success
	response.Message = message
	response.Data = data
	response.Errors = errors

	return response
}

func ResponseFormatCode(success bool, message string, data interface{}, errors []struct{ ErrorFormat }, code string) *OutputFormat {
	response := ResponseFormat(success, message, data, errors)
	response.Code = code

	return response
}
