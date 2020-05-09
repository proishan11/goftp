package server

const (
	ChangeDir                  = "cd"
	ListDir                    = "ls"
	PresentDir                 = "pwd"
	PermissionError            = "Cannot access file : Permission denied\n"
	CommandNotImplementedError = "Command not implemented\n"
	ServerStartMessage         = "FTP server running. Listening for requests...\n"
	ServerCloseMessage         = "Server stopped...\n"
	ConnectionSuccessMessage   = "Successfully connected...\n"
	OSError                    = "Some error occured\n"
)
