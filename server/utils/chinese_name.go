package utils

import (
	"sort"
	"unicode/utf8"

	"github.com/mozillazg/go-pinyin"
)

// 常见复姓表
var compoundSurnames = []string{
	"欧阳", "司马", "上官", "诸葛", "东方", "夏侯",
	"皇甫", "尉迟", "公孙", "慕容", "宇文", "长孙",
	"司徒", "轩辕", "令狐", "钟离", "闾丘", "子车",
	"亓官", "宗政", "濮阳", "公羊", "东郭", "南宫",
	"独孤", "太史", "申屠", "公冶", "西门", "公良",
	"拓跋", "呼延", "夹谷", "谷梁", "宰父", "端木",
	"公仲", "公上", "公门", "公山", "公坚", "左丘",
	"公伯", "公祖", "公乘", "公皙", "公广",
}

// 获取姓氏（支持复姓）
func getSurname(name string) string {
	if utf8.RuneCountInString(name) < 2 {
		return name
	}
	// 优先检查是否是复姓
	if len(name) >= 6 { // UTF-8 下一个汉字 3 字节
		firstTwo := name[:6]
		for _, s := range compoundSurnames {
			if s == firstTwo {
				return s
			}
		}
	}
	// 否则取单字姓
	r, _ := utf8.DecodeRuneInString(name)
	return string(r)
}

// 转换姓氏为拼音
func surnameToPinyin(surname string) string {
	a := pinyin.NewArgs()
	py := pinyin.LazyPinyin(surname, a)
	if len(py) == 0 {
		return surname
	}
	return py[0]
}

// SortChineseNames 按姓氏拼音升序排序
func SortChineseNames(names []string) {
	sort.Slice(names, func(i, j int) bool {
		si := surnameToPinyin(getSurname(names[i]))
		sj := surnameToPinyin(getSurname(names[j]))
		return si < sj
	})
}
