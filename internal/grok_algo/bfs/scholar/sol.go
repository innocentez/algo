package scholar

type Queue struct {
	CurrElement string
	Body        []string
}

func Solution() {
	graph := map[string][]string{
		"A": {"B", "C"},
		"B": {"E"},
		"C": {"D", "F"},
		"D": {"E"},
		"F": {"E"},
		"E": {"G"},
	}

	q := &Queue{CurrElement: "A"}
	search(graph, "E", q)
}


func search(graph map[string][]string, target string, queue *Queue) []string {
	q.
}