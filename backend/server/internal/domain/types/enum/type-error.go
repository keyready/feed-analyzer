package enum

type TypeError string

var (
	ValidationError TypeError = "validation_error"
	DatabaseError             = "database_error"
	ServerError               = "server_error"
)
