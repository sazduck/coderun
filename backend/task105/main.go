package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

func main() {
	Run(os.Stdin, os.Stdout)
}

func Run(r io.Reader, w io.Writer) error {
	grid, stationMap, err := readInput(r)
	if err != nil {
		return err
	}

	// Строим компоненты метро
	components := buildMetroComponents(stationMap, readMetroConnections(r))

	// Прямой маршрут
	forwardPath, forwardVal := findMaxPathForward(grid, stationMap, components)
	if forwardPath == nil {
		fmt.Fprintln(w, 0)
		return nil
	}

	// Удаляем собранные посылки
	grid = removeCollected(grid, forwardPath)

	// Обратный маршрут
	backwardPath, backwardVal := findMaxPathBackward(grid, stationMap, components)
	if backwardPath == nil {
		fmt.Fprintln(w, 0)
		return nil
	}

	fmt.Fprintln(w, forwardVal+backwardVal)
	return nil
}

// ---------- Ввод данных ----------
func readInput(r io.Reader) ([][]int64, map[int][2]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	readInt := func() int {
		scanner.Scan()
		var x int
		fmt.Sscanf(scanner.Text(), "%d", &x)
		return x
	}

	n := readInt()
	grid := make([][]int64, n)
	stationMap := make(map[int][2]int)
	for i := range n {
		grid[i] = make([]int64, n)
		for j := range n {
			val := readInt()
			grid[i][j] = int64(val)
			if val < 0 && val != -1 {
				stationMap[val] = [2]int{i, j}
			}
		}
	}
	return grid, stationMap, nil
}

func readMetroConnections(r io.Reader) [][2]int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	readInt := func() int {
		scanner.Scan()
		var x int
		fmt.Sscanf(scanner.Text(), "%d", &x)
		return x
	}
	m := readInt()
	conns := make([][2]int, m)
	for i := range m {
		conns[i] = [2]int{readInt(), readInt()}
	}
	return conns
}

// ---------- Построение компонент метро ----------
func buildMetroComponents(stationMap map[int][2]int, connections [][2]int) [][]int {
	adj := make(map[int][]int)
	for _, c := range connections {
		a, b := c[0], c[1]
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	visited := make(map[int]bool)
	var components [][]int
	for id := range stationMap {
		if !visited[id] {
			comp := []int{}
			queue := []int{id}
			visited[id] = true
			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]
				comp = append(comp, cur)
				for _, nb := range adj[cur] {
					if !visited[nb] {
						visited[nb] = true
						queue = append(queue, nb)
					}
				}
			}
			components = append(components, comp)
		}
	}
	return components
}

// ---------- Прямой маршрут ----------
type prevInfo struct {
	i, j int
	typ  byte
}

func findMaxPathForward(grid [][]int64, stationMap map[int][2]int, components [][]int) ([][2]int, int64) {
	n := len(grid)
	if grid[0][0] == -1 || grid[n-1][n-1] == -1 {
		return nil, -1
	}

	dp := make([][]int64, n)
	prev := make([][]prevInfo, n)
	for i := range n {
		dp[i] = make([]int64, n)
		prev[i] = make([]prevInfo, n)
		for j := range n {
			dp[i][j] = -1
		}
	}

	startVal := int64(0)
	if grid[0][0] > 0 {
		startVal = grid[0][0]
	}
	dp[0][0] = startVal
	prev[0][0] = prevInfo{-1, -1, 0}

	// Основные циклы
	for i := range n {
		for j := range n {
			if dp[i][j] == -1 {
				continue
			}
			// вправо
			if j+1 < n && grid[i][j+1] != -1 {
				add := int64(0)
				if grid[i][j+1] > 0 {
					add = grid[i][j+1]
				}
				if dp[i][j+1] < dp[i][j]+add {
					dp[i][j+1] = dp[i][j] + add
					prev[i][j+1] = prevInfo{i, j, 'r'}
				}
			}
			// вниз
			if i+1 < n && grid[i+1][j] != -1 {
				add := int64(0)
				if grid[i+1][j] > 0 {
					add = grid[i+1][j]
				}
				if dp[i+1][j] < dp[i][j]+add {
					dp[i+1][j] = dp[i][j] + add
					prev[i+1][j] = prevInfo{i, j, 'd'}
				}
			}
			// переход по метро (только для станций)
			if grid[i][j] < 0 && grid[i][j] != -1 {
				id := int(grid[i][j])
				for _, comp := range components {
					if slices.Contains(comp, id) {
						continue
					}
					for _, st := range comp {
						coords := stationMap[st]
						ni, nj := coords[0], coords[1]
						if ni >= i && nj >= j && (ni != i || nj != j) && grid[ni][nj] != -1 {
							add := int64(0)
							if grid[ni][nj] > 0 {
								add = grid[ni][nj]
							}
							if dp[ni][nj] < dp[i][j]+add {
								dp[ni][nj] = dp[i][j] + add
								prev[ni][nj] = prevInfo{i, j, 'm'}
							}
						}
					}
					break
				}
			}
		}
	}

	if dp[n-1][n-1] == -1 {
		return nil, -1
	}
	path := recoverPath(prev, n-1, n-1)
	return path, dp[n-1][n-1]
}

// ---------- Обратный маршрут ----------
func findMaxPathBackward(grid [][]int64, stationMap map[int][2]int, components [][]int) ([][2]int, int64) {
	n := len(grid)
	if grid[0][0] == -1 || grid[n-1][n-1] == -1 {
		return nil, -1
	}

	dp := make([][]int64, n)
	prev := make([][]prevInfo, n)
	for i := range n {
		dp[i] = make([]int64, n)
		prev[i] = make([]prevInfo, n)
		for j := range n {
			dp[i][j] = -1
		}
	}

	startVal := int64(0)
	if grid[n-1][n-1] > 0 {
		startVal = grid[n-1][n-1]
	}
	dp[n-1][n-1] = startVal
	prev[n-1][n-1] = prevInfo{-1, -1, 0}

	// Обратный обход: от нижнего правого к верхнему левому
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if dp[i][j] == -1 {
				continue
			}
			// влево
			if j-1 >= 0 && grid[i][j-1] != -1 {
				add := int64(0)
				if grid[i][j-1] > 0 {
					add = grid[i][j-1]
				}
				if dp[i][j-1] < dp[i][j]+add {
					dp[i][j-1] = dp[i][j] + add
					prev[i][j-1] = prevInfo{i, j, 'l'}
				}
			}
			// вверх
			if i-1 >= 0 && grid[i-1][j] != -1 {
				add := int64(0)
				if grid[i-1][j] > 0 {
					add = grid[i-1][j]
				}
				if dp[i-1][j] < dp[i][j]+add {
					dp[i-1][j] = dp[i][j] + add
					prev[i-1][j] = prevInfo{i, j, 'u'}
				}
			}
			// переход по метро (обратный: ni <= i, nj <= j)
			if grid[i][j] < 0 && grid[i][j] != -1 {
				id := int(grid[i][j])
				for _, comp := range components {
					if slices.Contains(comp, id) {
						continue
					}
					for _, st := range comp {
						coords := stationMap[st]
						ni, nj := coords[0], coords[1]
						if ni <= i && nj <= j && (ni != i || nj != j) && grid[ni][nj] != -1 {
							add := int64(0)
							if grid[ni][nj] > 0 {
								add = grid[ni][nj]
							}
							if dp[ni][nj] < dp[i][j]+add {
								dp[ni][nj] = dp[i][j] + add
								prev[ni][nj] = prevInfo{i, j, 'm'}
							}
						}
					}
					break
				}
			}
		}
	}

	if dp[0][0] == -1 {
		return nil, -1
	}
	path := recoverPath(prev, 0, 0)
	return path, dp[0][0]
}

// ---------- Вспомогательные ----------

func recoverPath(prev [][]prevInfo, endI, endJ int) [][2]int {
	path := [][2]int{}
	i, j := endI, endJ
	for !(i == -1 && j == -1) {
		path = append(path, [2]int{i, j})
		p := prev[i][j]
		i, j = p.i, p.j
	}
	// разворачиваем, чтобы путь шёл от начала к концу
	for l, r := 0, len(path)-1; l < r; l, r = l+1, r-1 {
		path[l], path[r] = path[r], path[l]
	}
	return path
}

func removeCollected(grid [][]int64, path [][2]int) [][]int64 {
	newGrid := make([][]int64, len(grid))
	for i := range grid {
		newGrid[i] = make([]int64, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	for _, cell := range path {
		if newGrid[cell[0]][cell[1]] > 0 {
			newGrid[cell[0]][cell[1]] = 0
		}
	}
	return newGrid
}
