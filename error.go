package main

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Param   string `json:"param,omitempty"`
}

func (e *Error) Error() string {
	return e.Message
}
