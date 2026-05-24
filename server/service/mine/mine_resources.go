package mine

import (
	"fmt"
	"github.com/springbear2020/self-hub/server/model/mine/dto"
)

type MineResourcesService struct{}

func (svc *MineResourcesService) SearchMineResources(uid uint, keyword string) ([]*dto.MineResourcesDTO, error) {
	if keyword == "" {
		return nil, fmt.Errorf("搜索关键词不能为空")
	}

	var results []*dto.MineResourcesDTO

	list, err := mineBlogsService.SearchMineBlogs(uid, keyword)
	if err != nil {
		return nil, err
	}
	results = append(results, list...)

	list, err = mineBooksService.SearchMineBooks(uid, keyword)
	if err != nil {
		return nil, err
	}
	results = append(results, list...)

	list, err = mineProjectsService.SearchMineProjects(uid, keyword)
	if err != nil {
		return nil, err
	}
	results = append(results, list...)

	return results, nil
}
