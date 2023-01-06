package methods

import "github.com/babbageLabs/brinch/bin/core/types"

func Validate[V types.IValidatable](validatables []V) (bool, error) {
	for _, validatable := range validatables {
		isValid, err := validatable.Validate()
		if !isValid {
			return false, err
		}
	}

	return true, nil
}

func Marshal(params *types.Params) ([]byte, error) {
	// TODO implement this
	return nil, nil
}
