package telegram

// represents adapter for relevant data of telego.Update
type TgInteraction struct {
	// raw text of the received update
	RawText string
	// parsed command in the update (without leading /)
	Cmd string
	// args passed with command, if any
	CmdArgs []string
	// inline selection data
	InlineSelectionData string
	// message that triggered the inline selection prompt
	InlineSelectionTrigger string
}

func (tgi TgInteraction) IsCommand() bool {
	return tgi.Cmd != ""
}

func (tgi TgInteraction) IsInline() bool {
	return tgi.InlineSelectionData != ""
}

// represents adapter for relevant data to create telego.SendMessageParams from bot
type TgBotResponse struct {
	// bot text response, if any
	ResponseText string
	// inline response options, if any
	InlineOptions []InlineOption
}

type InlineOption struct {
	// human-readable text of inline element
	Description string
	// payload of inline element sent back into TgInteraction.InlineSelectionData
	SelectionData string
}

func (tgbr TgBotResponse) IsInline() bool {
	return tgbr.InlineOptions != nil
}
