package dto

import "sort"

type DashboardStatDTO []*DashboardStatItem

type DashboardStatItem struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Unit  string `json:"unit"`
	Sort  int    `json:"sort"`
	Icon  string `json:"icon"`
}

func (d DashboardStatDTO) Len() int {
	return len(d)
}

func (d DashboardStatDTO) Less(i, j int) bool {
	return d[i].Sort < d[j].Sort
}

func (d DashboardStatDTO) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d DashboardStatDTO) Ascending() {
	sort.Sort(d)
}
