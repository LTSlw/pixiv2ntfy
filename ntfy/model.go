package ntfy

type Priority int

const (
	PriorityMin Priority = iota + 1
	PriorityLow
	PriorityDefault
	PriorityHigh
	PriorityMax
)

type Message struct {
	Title       string
	Message     string
	Priority    int
	Tags        []string
	Markdown    bool
	Delay       string
	Actions     []Action
	Click       string
	Filename    string
	Attach      string
	File        []byte
	Icon        string
	Email       string
	Call        string
	NoCache     bool
	NoFirebase  bool
	UnifiedPush bool
}
