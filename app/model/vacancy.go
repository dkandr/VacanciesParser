package model

type Vacancy struct {
	DatePosted     string     `json:"datePosted"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	Identifier     Identifier `json:"identifier"`
	ValidThrough   string     `json:"validThrough"`
	EmploymentType string     `json:"employmentType"`
}

type Identifier struct {
	Type  string `json:"@type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
