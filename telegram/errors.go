package telegram

type UnprocessableMessageError struct{}

func (e UnprocessableMessageError) Error() string {
	return "Cannot process message"
}
