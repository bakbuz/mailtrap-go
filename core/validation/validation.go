package validation

import "strings"

const (
	separator string = ", "
)

type ValidationResult struct {
	Errors []ValidationFailure `json:"errors"`
}

type ValidationFailure struct {
	PropertyName string `json:"propertyName"`
	ErrorMessage string `json:"errorMessage"`
}

func NewValidationResult() *ValidationResult {
	return &ValidationResult{
		Errors: []ValidationFailure{},
	}
}

func NewValidationResultFromFailures(failures []ValidationFailure) *ValidationResult {
	errors := make([]ValidationFailure, len(failures))
	for i, f := range failures {
		if f.ErrorMessage != "" {
			errors[i] = f
		}
	}
	return &ValidationResult{
		Errors: errors,
	}
}

func (vr *ValidationResult) IsValid() bool {
	return len(vr.Errors) == 0
}

func (vr *ValidationResult) String() string {
	messages := make([]string, 0, len(vr.Errors))
	for i, f := range vr.Errors {
		messages[i] = f.ErrorMessage
	}
	return strings.Join(messages, separator)
}

func (vr *ValidationResult) Map() map[string][]string {
	result := make(map[string][]string)

	for _, e := range vr.Errors {
		result[e.PropertyName] = append(result[e.PropertyName], e.ErrorMessage)
	}

	return result
}

func (vr *ValidationResult) AddError(propertyName string, errorMessage string) {
	vr.Errors = append(vr.Errors, ValidationFailure{PropertyName: propertyName, ErrorMessage: errorMessage})
}
