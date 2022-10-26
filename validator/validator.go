package validator

type Validator struct {
	Errors map[string]string
}

func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

// If there is no error, it means all validations are passed
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check a condition and add an error to error list if condition is not met
func (v *Validator) Validate(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}
