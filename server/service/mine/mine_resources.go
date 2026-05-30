package mine

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/springbear2020/self-hub/server/model/mine/dto"
)

type MineResourcesService struct{}

func (svc *MineResourcesService) SearchMineResources(uid uint, keyword string) ([]*dto.MineResourcesDTO, error) {
	if keyword == "" {
		return nil, fmt.Errorf("搜索关键词不能为空")
	}

	var mu sync.Mutex
	var results []*dto.MineResourcesDTO

	g, _ := errgroup.WithContext(context.Background())

	// Search blogs
	g.Go(func() error {
		list, err := mineBlogsService.SearchMineBlogs(uid, keyword)
		if err != nil {
			return fmt.Errorf("blogs search error: %w", err)
		}
		mu.Lock()
		results = append(results, list...)
		mu.Unlock()
		return nil
	})

	// Search books
	g.Go(func() error {
		list, err := mineBooksService.SearchMineBooks(uid, keyword)
		if err != nil {
			return fmt.Errorf("books search error: %w", err)
		}
		mu.Lock()
		results = append(results, list...)
		mu.Unlock()
		return nil
	})

	// Search projects
	g.Go(func() error {
		list, err := mineProjectsService.SearchMineProjects(uid, keyword)
		if err != nil {
			return fmt.Errorf("projects search error: %w", err)
		}
		mu.Lock()
		results = append(results, list...)
		mu.Unlock()
		return nil
	})

	// Search archives
	g.Go(func() error {
		list, err := mineArchivesService.SearchMineArchives(uid, keyword)
		if err != nil {
			return fmt.Errorf("archives search error: %w", err)
		}
		mu.Lock()
		results = append(results, list...)
		mu.Unlock()
		return nil
	})

	// Search sentences
	g.Go(func() error {
		list, err := mineSentencesService.SearchMineSentences(uid, keyword)
		if err != nil {
			return fmt.Errorf("sentences search error: %w", err)
		}
		mu.Lock()
		results = append(results, list...)
		mu.Unlock()
		return nil
	})

	if err := g.Wait(); err != nil {
		return results, err
	}

	return results, nil
}
