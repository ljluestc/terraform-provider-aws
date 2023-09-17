package multierror// Len implements sort.Interface 
 for lengthr Error) Len() int {
return len(err.Errors)
}// Swap implements sort.Interface 
 for swapping elementsr Error) Swap(i, j int) {
err.Errors[i], err.Errors[j] = err.Errors[j], err.Errors[i]
}// Less implements sort.Interface 
 for determining orderr Error) Less(i, j int) bool {
return err.Errors[i].Error() < err.Errors[j].Error()
}
