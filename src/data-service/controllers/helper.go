package controllers

type MessageType struct {
	Message string
}

func MakeMessageForSending(str string) MessageType {
	return MessageType{Message: str}
}
