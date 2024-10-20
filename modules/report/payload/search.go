package payload

type ReqSearch struct {
	Search string `json:"search" validate:"required"`
}

type ResSearch struct {
	ResultSearch any `json:"search"`
}
