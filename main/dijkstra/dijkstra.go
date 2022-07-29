package main

import "fmt"

const (
	MAXVEX = 9 // 限制顶点个数，也可以不要
	INF    = 0x3f3f3f3f
)

func dijkstra(D [][]int, start int) []int {
	n := len(D)            // 顶点个数
	vis := make([]bool, n) // 标记已经访问过的节点
	dist := make([]int, n) // 标记从start到其他节点的最短路径

	for i := 0; i < n; i++ {
		dist[i] = D[start][i] //遍历节点
	}

	vis[start] = true // 初始节点一定是被访问过的

	for i := 0; i < n; i++ {
		minDist := INF // 表示从起点到其他未被访问节点的最短路径
		minNode := 0   // 最短路径的节点编号

		// 遍历n个节点, 找未被访问且距离最小的点, 记录为minNode
		for j := 0; j < n; j++ {
			if !vis[j] && dist[j] < minDist {
				minDist = dist[j]
				minNode = j
			}
		}
		// 判断从start -> minNode -> j 和 start -> j的距离哪个小
		for j := 0; j < n; j++ {
			if !vis[j] && dist[j] > dist[minNode]+D[minNode][j] {
				dist[j] = dist[minNode] + D[minNode][j]
			}
		}
		vis[minNode] = true
	}
	return dist
}

func main() {
	var D = [][]int{
		{0, 100, 1200, INF, INF, INF},
		{100, 0, 900, 300, INF, INF},
		{1200, 900, 0, 400, 500, INF},
		{INF, 300, 400, 0, 1300, 1400},
		{INF, INF, 500, 1300, 0, 1500},
		{INF, INF, INF, 1400, 1500, 0},
	}
	dist := dijkstra(D, 0)
	fmt.Println(dist[5])
}
