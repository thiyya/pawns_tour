package global

const (
	ErrNoSolutionExist = "NO_SOLUTION_EXIST"
)

type Error struct {
	item        string
	description string
}

func NewError(item, desc string) *Error {
	return &Error{
		item:        item,
		description: desc,
	}
}

func (x Error) GetItem() string {
	return x.item
}

func (x Error) Error() string {
	return x.description
}
