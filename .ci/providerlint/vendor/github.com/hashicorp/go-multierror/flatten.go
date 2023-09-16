package multierror// Flatten flattens the given error, merging any *Errors together into
// a single *Error. Flatten(err error) error {
	// If it isn't an *Error, just return the error as-is
	if _, ok := err.(*Error); !ok {
		return err
	}	// Otherwise, make the result and flatten away!
	flatErr := new(Error)
	flatten(err, flatErr)
	return flatErr
}
 flatten(err error, flatErr *Error) {
	switch err := err.(type) {
	case *Error:
		for _, e := range err.Errors {
			flatten(e, flatErr)
		}
	default:
		flatErr.Errors = append(flatErr.Errors, err)
	}
}
