package domain

type Vehicle struct {
	Url   string `json:"url"`
	Title string
	Price string
	Year  int
}

type MakeModelResponse struct {
	OEMs []MakeModelResponseOEM `json:"makeModels"`
}

type MakeModelResponseOEM struct {
	Title  string                   `json:"make"`
	Models []MakeModelResponseModel `json:"children"`
}

type MakeModelResponseModel struct {
	Title string `json:"model"`
}
