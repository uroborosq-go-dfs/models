package tcp_operation_code

const (
	SendFile byte = iota + 1
	RequestFile
	RequestList
	RequestSize
	RemoveFile
)
