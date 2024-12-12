package tabu

import "github.com/kruily/go-timefold-solver/solver/api"

type TabuItem struct {
	hash      string     // 哈希值
	iteration int        // 迭代次数
	score     api.IScore // 分数
}
