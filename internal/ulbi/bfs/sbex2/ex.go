package sbex2

import "fmt"

func Example() {
	graph := map[string][]string{
		"a": {"b", "c"},
		"b": {"e"},
		"c": {"d", "f"},
		"d": {"g"},
		"f": {"e"},
		"g": {},
	}

	resMap := bfs(graph, "a", "f")
	fmt.Println(resMap)
}

func bfs(graph map[string][]string, start string, end string) bool {
	visited := map[string]struct{}{}
	queue := []string{start}
	visited[start] = struct{}{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, child := range graph[node] {
			if child == end {
				return true
			}

			if _, ok := visited[child]; !ok {
				visited[child] = struct{}{}
				queue = append(queue, child)
			}
		}
	}

	return false
}
