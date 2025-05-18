package models

type ProfileResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SkinResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Properties []struct {
		Name      string `json:"name"`
		Value     string `json:"value"`
		Signature string `json:"signature"`
	} `json:"properties"`
}

type SkinURLResponse struct {
	Username string `json:"username"`
	UUID     string `json:"uuid"`
	SkinURL  string `json:"skin_url"`
} 
