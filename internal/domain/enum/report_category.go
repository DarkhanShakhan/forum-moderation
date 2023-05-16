package enum

type ReportCategory string

const (
	ReportCategoryIrrelevant ReportCategory = "irrelevant"
	ReportCategoryObscene    ReportCategory = "obscene"
	ReportCategoryIllegal    ReportCategory = "illegal"
	ReportCategoryInsulting  ReportCategory = "insulting"
)
