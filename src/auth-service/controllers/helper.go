package controllers

type MessageType struct {
	Message string `json:"Message"`
}

func MakeMessageForSending(str string) MessageType {
	return MessageType{Message: str}
}
