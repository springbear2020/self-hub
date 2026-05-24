package dto

type MineWebsitesDTO struct {
	Categories []*MineWebsiteCategory `json:"categories"`
	Sites      []*MineWebsite         `json:"sites"`
}

type MineWebsiteCategory struct {
	Label  string `json:"label"`
	Value  string `json:"value"`
	Extend string `json:"extend"`
}

type MineWebsite struct {
	ID       *int    `json:"id"`
	Name     *string `json:"name"`
	Title    *string `json:"title"`
	Href     *string `json:"href"`
	Icon     *string `json:"icon"`
	Category *string `json:"category"`
}
