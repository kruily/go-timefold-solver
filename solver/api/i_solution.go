package api

type ISolution interface {

	// GetScore 获取当前解决方案的评分
	GetScore() IScore
	// SetScore 设置解决方案的评分
	SetScore(score IScore)
	// GetProblemFacts 获取问题事实（不可变的问题数据）
	GetProblemFacts() []interface{}
}
