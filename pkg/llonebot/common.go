package llonebot

func NewTextMsg(content string) Message {
	return Message{
		Type: Text,
		Data: MessageData{
			Text: content + "\n",
		},
	}
}

func NewImageMsg(file string) Message {
	return Message{
		Type: Image,
		Data: MessageData{
			File: file + "\n",
		},
	}
}
