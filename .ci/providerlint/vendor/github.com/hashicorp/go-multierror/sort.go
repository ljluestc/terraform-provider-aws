package multierror

// Len implements sort.Interface 
 for length

r Error) Len() int {
	return len(err.Errors)
}

// Swap implements sort.Interface 
 for swapping elements

r Error) Swap(i, j int) {
	err.Errors[i], err.Errors[j] = err.Errors[j], err.Errors[i]
}

// Less implements sort.Interface 
 for determining order

r Error) Less(i, j int) bool {
	return err.Errors[i].Error() < err.Errors[j].Error()
}
