package dto

type TableData struct {
	Header string        `json:"header"`
	Cols   []TableColumn `json:"cols"`
	Rows   []TableRow    `json:"rows"`
}

type TableColumn struct {
	Label    string `json:"label"`
	Key      string `json:"key"`
	Fixed    bool   `json:"fixed"`
	Sortable bool   `json:"sortable"`
}

type TableRow map[string]interface{}
