package mine

import (
	"github.com/springbear2020/self-hub/server/global"
	"github.com/springbear2020/self-hub/server/model/common"
	"github.com/springbear2020/self-hub/server/model/mine"
	"github.com/springbear2020/self-hub/server/model/mine/dto"
	"github.com/springbear2020/self-hub/server/model/mine/request"
)

type MineSentencesService struct{}

func (mineSentencesService *MineSentencesService) CreateMineSentences(uid uint, mineSentences *mine.MineSentences) (err error) {
	mineSentences.UserId = uid
	err = global.GVA_DB.Create(mineSentences).Error
	return err
}

func (mineSentencesService *MineSentencesService) DeleteMineSentences(uid uint, id string) (err error) {
	return mineSentencesService.DeleteMineSentencesByIds(uid, []string{id})
}

func (mineSentencesService *MineSentencesService) DeleteMineSentencesByIds(uid uint, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]mine.MineSentences{}, "user_id = ? AND id in ?", uid, ids).Error
	return err
}

func (mineSentencesService *MineSentencesService) UpdateMineSentences(uid uint, mineSentences mine.MineSentences) (err error) {
	err = global.GVA_DB.Model(&mine.MineSentences{}).Where("user_id = ? AND id = ?", uid, mineSentences.Id).Updates(&mineSentences).Error
	return err
}

func (mineSentencesService *MineSentencesService) GetMineSentences(uid uint, id string) (mineSentences mine.MineSentences, err error) {
	err = global.GVA_DB.Where("user_id = ? AND id = ?", uid, id).First(&mineSentences).Error
	return
}

func (mineSentencesService *MineSentencesService) GetMineSentencesInfoList(uid uint, info request.MineSentencesSearch) (list []mine.MineSentences, total int64, err error) {
	// 查询条件
	db := global.GVA_DB.Model(&mine.MineSentences{}).Where("user_id = ? ", uid)
	if info.Title != nil && *info.Title != "" {
		db = db.Where("title LIKE ?", "%"+*info.Title+"%")
	}
	if info.Author != nil && *info.Author != "" {
		db = db.Where("author LIKE ?", "%"+*info.Author+"%")
	}
	if info.Dynasty != nil && *info.Dynasty != "" {
		db = db.Where("dynasty LIKE ?", "%"+*info.Dynasty+"%")
	}
	if info.Content != nil && *info.Content != "" {
		db = db.Where("content LIKE ?", "%"+*info.Content+"%")
	}
	if info.CategoryId != nil {
		db = db.Where("category_id = ?", *info.CategoryId)
	}

	// 记录总数
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	// 分页条件
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	// 查询记录
	var mineSentences []mine.MineSentences
	err = db.Order("sort_value desc").Find(&mineSentences).Error
	return mineSentences, total, err
}

func (mineSentencesService *MineSentencesService) SearchMineSentences(uid uint, keyword string) (list []*dto.MineResourcesDTO, err error) {
	db := global.GVA_DB.
		Table("mine_sentences").
		Model(&dto.MineResourcesDTO{}).
		Where("user_id = ? AND (title LIKE ? OR content LIKE ?)", uid, "%"+keyword+"%", "%"+keyword+"%")
	err = db.Select("id, title, content as description, link").Order("sort_value desc").Limit(searchPageSize).Find(&list).Error
	if err != nil {
		return
	}

	for _, item := range list {
		item.Type = searchTypeSentence
	}
	return
}

func (mineSentencesService *MineSentencesService) GetMineSentencesStat(uid uint) ([]*common.CountDTO, error) {
	var categories []*common.CountDTO
	err := global.GVA_DB.Raw(`
	SELECT category_id AS label, COUNT(1) AS total
	FROM mine_sentences
	WHERE user_id = ?
	GROUP BY label
	ORDER BY total DESC;
	`, uid).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}
