package payload

type ReqMassCreate struct {
	Total int `json:"total" validate:"required,numeric"`
}

type ResMassCreate struct {
	TotalInput   int    `json:"total_create"`
	LastInsertID string `json:"last_insert_id"`
}
