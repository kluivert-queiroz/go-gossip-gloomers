package cmd

var challenges = map[string]interface{}{
	"echo":       []string{"-w", "echo", "--bin", "./challenges/go-gossip-gloomers", "--node-count", "1", "--time-limit", "10"},
	"unique-ids": []string{"-w", "unique-ids", "--bin", "./challenges/go-gossip-gloomers", "--time-limit", "30", "--rate", "1000", "--node-count", "3", "--availability", "total", "--nemesis", "partition"},
	"broadcast": map[string]interface{}{
		"1. Single-Node Broadcast":       []string{"-w", "broadcast", "--bin", "./challenges/go-gossip-gloomers", "--node-count", "1", "--time-limit", "20", "--rate", "10"},
		"2. Multi-Node Broadcast":        []string{"-w", "broadcast", "--bin", "./challenges/go-gossip-gloomers", "--node-count", "5", "--time-limit", "20", "--rate", "10"},
		"3. Fault Tolerant Broadcast":    []string{"-w", "broadcast", "--bin", "./challenges/go-gossip-gloomers", "--node-count", "5", "--time-limit", "20", "--rate", "10", "--nemesis", "partition"},
		"4. Efficient Broadcast":         []string{"-w", "broadcast", "--bin", "./challenges/go-gossip-gloomers", "--node-count", "25", "--time-limit", "20", "--rate", "100", "--latency", "100"},
		"5. Efficient Broadcast Part II": []string{"-w", "broadcast", "--bin", "./challenges/go-gossip-gloomers", "--node-count", "25", "--time-limit", "20", "--rate", "100", "--latency", "100"},
	},
	"g-counter": []string{"-w", "g-counter", "--bin", "./challenges/go-gossip-gloomers", "--node-count", "3", "--rate", "100", "--time-limit", "20", "--nemesis", "partition"},
	"kafka":     []string{"-w", "kafka", "--bin", "./challenges/go-gossip-gloomers", "--node-count", "1", "--concurrency", "2n", "--time-limit", "20", "--rate", "1000"},
	"txn-rw-register ": map[string]interface{}{
		"1. Single-Node, Totally Available Transactions":      []string{"-w", "txn-rw-register", "--bin", "./challenges/go-gossip-gloomers", "--node-count", "1", "--time-limit", "20", "--rate", "1000", "--concurrency", "2n", "--consistency-models", "read-uncommitted", "--availability", "total"},
		"2. Totally-Available, Read Uncommitted Transactions": []string{"-w", "txn-rw-register", "--bin", "./challenges/go-gossip-gloomers", "--node-count", "2", "--concurrency", "2n", "--time-limit", "20", "--rate", "1000", "--consistency-models", "read-uncommitted", "--availability", "total", "--nemesis", "partition"},
		"3. Totally-Available, Read Committed Transactions":   []string{"-w", "txn-rw-register", "--bin", "./challenges/go-gossip-gloomers", "--node-count", "2", "--concurrency", "2n", "--time-limit", "20", "--rate", "1000", "--consistency-models", "read-committed", "--availability", "total", "--nemesis", "partition"},
	},
}
