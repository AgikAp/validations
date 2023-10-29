package validations

type cerros struct {
	messages map[string]string
}

func NewCerros(messages map[string]string) *cerros {
	return &cerros{messages: messages}
}

func (c *cerros) GetMessages() map[string]string {
	return c.messages
}
