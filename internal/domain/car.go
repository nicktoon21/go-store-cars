package domain

type Car struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
	Brand Brand  `json:"brand"`
	Model Model  `json:"model"`
}

type Brand struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Model struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
