package payload

type ResUpdate struct {
	Name  string `json:"name"`
	Age   uint64 `json:"age"`
	Grade uint64 `json:"grade"`
}
