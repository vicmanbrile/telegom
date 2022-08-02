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
)
