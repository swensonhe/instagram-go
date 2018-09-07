package instagram

type Error string

const (
	ErrInternal = Error("internal")
)

func (e Error) Error() string {
	return string(e)
}
