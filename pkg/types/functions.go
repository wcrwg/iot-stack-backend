package types

func (message *IotMessage) AddTag(key string, value string) *IotMessage {
	if message.Tags == nil {
		message.Tags = make(map[string]string, 0)
	}

	message.Tags[key] = value

	return message
}

func (message *IotMessage) AddField(key string, value interface{}) *IotMessage {
	if message.Fields == nil {
		message.Fields = make(map[string]interface{}, 0)
	}

	message.Fields[key] = value

	return message
}
