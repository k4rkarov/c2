package commons

type Message struct {
	AgentId       string
	AgentHostname string
	AgentCWD      string
	Commands      []Command
}
