package other

import "demo/util"

// 调用链路最长耗时
// (A,B,100)表示A调用B耗时100ms 求调用链路下最长的耗时

type invokeNode struct {
	startName string
	endName   string
	cost      int
}

func getLongestCost(nodes []*invokeNode) int {
	if len(nodes) == 0 {
		return 0
	}

	ret := 0

	var dfs func([]*invokeNode, int, int, []string)
	dfs = func(invokeNodes []*invokeNode, cost, currentIndex int, path []string) {
		ret = util.GetMaxInt(ret, cost)

		for i, v := range invokeNodes {
			// 如果当前节点不是调用链路的下游 或当前节点正在处理 则跳过
			if v.startName != path[len(path)-1] || i == currentIndex {
				continue
			}

			path = append(path, v.startName, v.endName)
			dfs(invokeNodes, cost+v.cost, i, path)
			path = path[:len(path)-2]
		}

	}

	for i, v := range nodes {
		path := make([]string, 0)
		path = append(path, v.startName, v.endName)
		dfs(nodes, v.cost, i, path)
	}

	return ret

}
