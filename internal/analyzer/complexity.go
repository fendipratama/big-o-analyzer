package analyzer

type Complexity struct {
	Time  string
	Space string
}

func MergeComplexity(a, b Complexity) Complexity {
	if a.Time == "O(n^2)" || b.Time == "O(n^2)" {
		return Complexity{Time: "O(n^2)", Space: "O(1)"}
	}
	if a.Time == "O(n)" || b.Time == "O(n)" {
		return Complexity{Time: "O(n)", Space: "O(1)"}
	}
	return Complexity{Time: "O(1)", Space: "O(1)"}
}
