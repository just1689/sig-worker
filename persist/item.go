package persist

var ItemEntity = "items"

type Item struct {
	ID        string `json:"id"`
	ExtraPK   string `json:"extra_pk"`
	JobID     string `json:"job_id"`
	Title     string `json:"title"`
	Value     string `json:"v"`
	CreatedBy string `json:"created_by"`
}
