package constants

type ServingType string

const (
	CommentServingType ServingType = "comment"
)

func (st ServingType) String() string {
	return string(st)
}
