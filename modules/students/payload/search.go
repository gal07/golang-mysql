package payload

type ReqSearch struct {
	Search string `json:"search" validate:"required,alpha"`
}

type ResSearch struct {
	ResultSearch any `json:"search"`
}
