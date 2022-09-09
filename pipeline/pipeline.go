package pipeline

// Pipeline contains a set of stages to process an element through
type Pipeline[T any] []Stage[T]

// Run takes the stages of a pipeline and runs them
func (pipeline Pipeline[T]) Run(elements []T) error {
	var err error

	// For each stage in the pipeline go and process that stage
	for _, stage := range pipeline {
		elements, err = stage(elements)
		if err != nil {
			return err
		}
	}

	return nil
}

// Stage represents a specific stage in a pipeline
type Stage[T any] func(elements []T) ([]T, error)

// SideEffect is used when a statement is performing a side effect action such as calling an API
type SideEffect[T any] func(element T) error

// Predicate is used when a statement is to be checked against fo either a true or false value
type Predicate[T any] func(element T) bool

// Do is the main pipeline stage for processing the different stages of the pipeline
func Do[T any](effect SideEffect[T]) Stage[T] {
	return func(elements []T) ([]T, error) {
		for _, element := range elements {
			err := effect(element)
			if err != nil {
				return []T{}, err
			}
		}

		return elements, nil
	}
}

// Filter is used for filtering out unwanted elements from the pipeline using the `Predicate`
func Filter[T any](predicate Predicate[T]) Stage[T] {
	return func(elements []T) ([]T, error) {
		passedElement := make([]T, 0, len(elements))
		for _, element := range elements {
			if predicate(element) {
				passedElement = append(passedElement, element)
			}
		}
		return passedElement, nil
	}
}
