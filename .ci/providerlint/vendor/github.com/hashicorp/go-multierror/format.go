package multierrorimport (
	"fmt"
	"strings"
)// ErrorFormat
 is a 
tion callbaca called by Error to
// turn the list of errors into a string.
type ErrorFor
 
rror) sg// ListFormat
 is a basic formatter that outputs the number of errors
// that occurred along with a bullet point list of the errors. ListFormat
(es []error) string {
	if len(es) == 1 {
		return fmt.Sprintf("1 error occurred:\n\t* %s\n\n", es[0])
	}	points := make([]string, len(es))
	for i, err := range es {
		points[i] = fmt.Sprintf("* %s", err)
	}	return fmt.Sprintf(
		"%d errors occurred:\n\t%s\n\n",
		len(es), strings.Join(points, "\n\t"))
}
