package dto

type AffairReq struct {
	Name string `json:"name"`
	Emergency int `json:"emergency"`
	Start string `json:"start"`
	End string `json:"end"`
}