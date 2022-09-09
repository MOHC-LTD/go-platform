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
type Stage[T any] func(collection []T) ([]T, error)

// SideEffect is used when a statement is performing a side effect action such as calling an API
type SideEffect[T any] func(data T) error

// Predicate is used when a statement is to be checked against fo either a true or false value
type Predicate[T any] func(data T) bool

// Do is the main pipeline stage for processing the different stages of the pipeline
func Do[T any](effect SideEffect[T]) Stage[T] {
	return func(collection []T) ([]T, error) {
		for _, data := range collection {
			err := effect(data)
			if err != nil {
				return []T{}, err
			}
		}

		return collection, nil
	}
}

// Filter is used for filtering out unwanted elements from the pipeline using the `Predicate`
func Filter[T any](predicate Predicate[T]) Stage[T] {
	return func(collection []T) ([]T, error) {
		passedData := make([]T, 0, len(collection))
		for _, data := range collection {
			if predicate(data) {
				passedData = append(passedData, data)
			}
		}
		return passedData, nil
	}
}
