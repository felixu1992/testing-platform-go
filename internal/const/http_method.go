package _const

type HttpMethod string

const (
	GET     = HttpMethod("GET")
	HEAD    = HttpMethod("HEAD")
	POST    = HttpMethod("POST")
	PUT     = HttpMethod("PUT")
	PATCH   = HttpMethod("PATCH")
	DELETE  = HttpMethod("DELETE")
	OPTIONS = HttpMethod("OPTIONS")
	TRACE   = HttpMethod("TRACE")
)

func ValueOf(s string) (HttpMethod, error) {
	switch HttpMethod(s) {
	case GET:
		return GET, nil
	case HEAD:
		return HEAD, nil
	case POST:
		return POST, nil
	case PUT:
		return PUT, nil
	case PATCH:
		return PATCH, nil
	case DELETE:
		return DELETE, nil
	case OPTIONS:
		return OPTIONS, nil
	case TRACE:
		return TRACE, nil
	default:
		return "", nil
	}
}
