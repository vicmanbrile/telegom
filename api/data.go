package api

type (
	Bot struct {
		Ok     bool `json:"ok"`
		Result User `json:"result"`
	}

	Updates struct {
		Ok     bool     `json:"ok"`
		Update []Update `json:"result"`
	}

	Update struct {
		UpdateID          int         `json:"update_id"`
		Message           Message     `json:"Message"`
		EditedMessage     Message     `json:"edited_message"`
		ChannelPost       Message     `json:"channel_post"`
		EditedChannelPost Message     `json:"edited_channel_post"`
		InlineQuery       InlineQuery `json:"inline_query"`
		// ChosenInlineResult
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

	Chat struct {
		ID                 int       `json:"id"`
		Type               string    `json:"type"`
		Title              string    `json:"title"`
		Username           string    `json:"username"`
		FirstName          string    `json:"first_name"`
		LastName           string    `json:"last_name"`
		Photo              ChatPhoto `json:"photo"`
		Bio                string    `json:"bio"`
		HasPrivateForwards bool      `json:"has_private_forwards"`
		JoinToSendMessage  bool      `json:"join_to_send_messages"`
		JoinByRequest      bool      `json:"join_by_request"`
		Description        string    `json:"description"`
		InviteLink         string    `json:"invite_link"`
		// PinnedMessage      Message         `json:"pinned_message"`
		Permissions ChatPermissions `json:"permissions"`
	}

	ChatPhoto struct {
		SmallFileID       string `json:"small_file_id"`
		SmallFileUniqueID string `json:"small_file_unique_id"`
		BigFileId         string `json:"big_file_id"`
		BigFileUniqueID   string `json:"big_file_unique_id"`
	}

	ChatPermissions struct {
		CanSendMessages      bool
		CanSendMediaMessages bool
		CanSendPoll          bool
	}

	Message struct {
		ChatID                string          `json:"chat_id"`
		From                  User            `json:"from"`
		SenderChat            Chat            `json:"sender_Chat"`
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

	InlineQuery struct {
		ID       string   `json:"id"`
		From     User     `json:"from"`
		Query    string   `json:"query"`
		Offset   string   `json:"offset"`
		ChatType string   `json:"chat_type"`
		Location Location `json:"location"`
	}

	ChosenlnLineResult struct {
		ResultID        string   `json:"result_id"`
		From            User     `json:"from"`
		Location        Location `json:"location"`
		InlineMessageID string   `json:"inline_message_id"`
		Query           string   `json:"query"`
	}

	CallbackQuery struct {
		ID               string  `json:"id"`
		From             User    `json:"from"`
		Message          Message `json:"message"`
		InLineMesssageID string  `json:"inline_messsage_id"`
		ChatInstance     string  `json:"chat_instance"`
		Data             string  `json:"data"`
		GameShortName    string  `json:"game_short_name"`
	}

	ShippingQuery struct {
		ID   string `json:"id"`
		From User   `json:"from"`
	}

	// Types Sending

	Location struct {
		Longitude            float64 `json:"longitude"`
		Latitude             float64 `json:"latitude"`
		HorizontalAccuracy   float64 `json:"horizontal_accuracy"`
		LivePeriod           int     `json:"live_period"`
		Heading              int     `json:"heading"`
		ProximityAlertRadius int     `json:"proximity_alert_radius"`
	}
)
