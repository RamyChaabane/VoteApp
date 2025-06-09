package vote

type Option string

const (
	Cats Option = "Cats"
	Dogs Option = "Dogs"
)

type Vote struct {
	Option Option
}
