package main

type Error struct {
	Status  int    `json:"status,omitempty"`
	Message string `json:"message"`
	Param   string `json:"param,omitempty"`
}

func (e *Error) Error() string {
	return e.Message
}
