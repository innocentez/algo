package sbex1

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

	res := bfs(graph, "a")
	fmt.Println(res)
}

func bfs(graph map[string][]string, start string) map[string]struct{} {
	visited := map[string]struct{}{}
	queue := []string{start}
	visited[start] = struct{}{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, child := range graph[node] {
			if _, ok := visited[child]; !ok {
				visited[child] = struct{}{}
				queue = append(queue, child)
			}
		}
	}

	return visited
}
