package api

type (
	Bot struct {
		Ok     bool `json:"ok"`
		Result User `json:"result"`
	}

	User struct {
		ID                      int    `json:"id"`
		IsBot                   bool   `json:"is_bot"`
		FirstName               string `json:"first_name"`
		LastName                string `json:"last_name"`
		Username                string `json:"username"`
		CanJoinGroups           bool   `json:"can_join_groups"`
		CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
		SupportsInlineQueries   bool   `json:"supports_inline_queries"`
		LanguageCode            string `json:"language_code"`
		IsPremium               bool   `json:"is_premium"`
		AddToAttachmentMenu     bool   `json:"added_to_attachment_menu"`
	}

	Message struct {
		ChatID                string          `json:"chat_id"`
		Text                  string          `json:"text"`
		ParseMode             string          `json:"parse_mode"`
		Entities              []MessageEntity `json:"entities"`
		DisableWebPagePreview bool            `json:"disable_web_page_preview"`
		DisableNotification   bool            `json:"disable_notification"`
		ProtectNotification   bool            `json:"protect_notification"`
		ReplyToMessageID      bool            `json:"reply_to_message_id"`
		// Reply Markup (( Pendient ))
	}

	MessageEntity struct {
		Type     string `json:"type"`
		Offset   int    `json:"offset"`
		Length   int    `json:"length"`
		URL      string `json:"url"`
		User     User   `json:"user"`
		Language string `json:"language"`
	}
)
