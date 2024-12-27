package utils

type ValidatorTypes int

const (
	Required ValidatorTypes = iota
	Min
	Max
	EndsWith
	StartsWith
	Email
	Contains
	Regexp
	FederalDocument
)

func ContainsTypes(slice []ValidatorTypes, search ValidatorTypes) bool {
	for _, v := range slice {
		if v == search {
			return true
		}
	}
	return false
}
