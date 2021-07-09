package survey

type Survey struct {
	ID     int    `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
	Name   string `json:"name,omitempty"`
}

type Question struct {
	ID       int    `json:"id,omitempty"`
	Title    string `json:"title"`
	SurveyId int    `json:"survey_id,omitempty"`
}

type SurveyWithQuestion struct {
	ID        int    `json:"id,omitempty"`
	Status    string `json:"status,omitempty"`
	Name      string `json:"name,omitempty"`
	Questions []Question
}
