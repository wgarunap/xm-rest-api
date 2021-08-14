package domain

type Company struct {
	Name    string `json:"name"`
	Code    string `json:"code"`
	Country string `json:"country"`
	Website string `json:"website"`
	Phone   int    `json:"phone"`
}
