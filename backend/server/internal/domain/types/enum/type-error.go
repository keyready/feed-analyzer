package enum

type TypeError string

var (
	ValidationError TypeError = "validation_error"
	DatabaseError   TypeError = "database_error"
	ServerError     TypeError = "server_error"
)
