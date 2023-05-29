package enum

import "strings"

type ReportCategory string

const (
	ReportCategoryIrrelevant ReportCategory = "irrelevant"
	ReportCategoryObscene    ReportCategory = "obscene"
	ReportCategoryIllegal    ReportCategory = "illegal"
	ReportCategoryInsulting  ReportCategory = "insulting"
)

func (c ReportCategory) OneOf(cats ...ReportCategory) bool {
	for _, cat := range cats {
		if cat == c {
			return true
		}
	}
	return false
}

var (
	reportCategories = map[string]ReportCategory{
		"irrelevant": ReportCategoryIrrelevant,
		"obscene":    ReportCategoryObscene,
		"illegal":    ReportCategoryIllegal,
		"insulting":  ReportCategoryInsulting,
	}
)

func ParseStringToReportCategory(s string) (ReportCategory, bool) {
	c, ok := reportCategories[strings.ToLower(s)]
	return c, ok
}
