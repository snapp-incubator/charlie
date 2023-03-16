package model

// Question model.
// each question is for a class.
// each question has a content file and test case directory.
// each question has a set of extensions.
type Question struct {
	Base
	ClassID     uint   `json:"class_id"`
	ContentFile string `json:"content_file"`
	TestCaseDir string `json:"test_case_dir"`
	Extensions  string `json:"extensions"`
}
