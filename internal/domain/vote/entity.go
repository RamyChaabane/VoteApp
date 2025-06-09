package vote

type Option string

const (
	Cats Option = "Cats"
	Dogs Option = "Dogs"
)

func IsValidOption(option string) bool {
	return option == string(Cats) || option == string(Dogs)
}
