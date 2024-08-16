package telegram

type UnprocessableMessageError struct{}

func (e UnprocessableMessageError) Error() string {
	return "Cannot process message"
}

type MissingCommandError struct{ Command string }

func (e MissingCommandError) Error() string {
	return "Cannot process command: " + e.Command
}
