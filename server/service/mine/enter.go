package mine

import "github.com/springbear2020/self-hub/server/service/system"

type ServiceGroup struct {
	MineBooksService
	MineBlogsService
	MineProjectsService
	MineResourcesService
	MineWebsitesService
	MineArchivesService
	MineSentencesService
}

var (
	mineBooksService     = MineBooksService{}
	mineBlogsService     = MineBlogsService{}
	mineProjectsService  = MineProjectsService{}
	mineArchivesService  = MineArchivesService{}
	mineSentencesService = MineSentencesService{}
	dictionaryService    = system.DictionaryService{}
)

const (
	searchPageSize = 5

	searchTypeBlog     = "blog"
	searchTypeBook     = "book"
	searchTypeProject  = "project"
	searchTypeArchive  = "archive"
	searchTypeSentence = "sentence"
)
