package payload

type ReqInsert struct {
	Name  string `json:"name" validate:"required"`
	Age   uint64 `json:"age" validate:"required,gte=0,lte=50,numeric,number"`
	Grade uint64 `json:"grade" validate:"required,gte=0,lte=20,numeric,number"`
}
