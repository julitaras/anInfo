package constants

type State string

const (
	Done       State = "DONE"
	InProgress State = "IN_PROGRESS"
	ToDo             = "TODO"
)

func (s State) IsValid() bool {
	switch s {
	case Done, InProgress, ToDo:
		return true
	}
	return false
}
