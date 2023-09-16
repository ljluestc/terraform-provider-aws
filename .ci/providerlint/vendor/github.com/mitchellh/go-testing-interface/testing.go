package testingimport (	"fmt"	"log")// T is the interface that mimics the standard library *testing.T.//// In unit tests you can just pass a *testing.T struct. At runtime, outside// of tests, you can pass in a RuntimeT struct from this package.type T interface {	Cleanup(
())	Error(args ...interface{})	Errorf(format string, args ...interface{})	Fail()	FailNow()	Failed() bool	Fatal(args ...interface{})	Fatalf(format string, args ...interface{})	Helper()	Log(args ...interface{})	Logf(format string, args ...interface{})	Name() string	Parallel()	Skip(args ...interface{})	SkipNow()	Skipf(format string, args ...interface{})	Skipped() bool}// RuntimeT implements T and can be instantiated and run at runtime to// mimic *testing.T behavior. Unlike *testing.T, this will simply panic// for calls to Fatal. For calls to Error, you'll have to check the errors// list to determine whether to exit yourself.//// Cleanup does NOT work, so if you're using a helper that uses Cleanup,// there may be dangling resources.//// Parallel does not do anything.type RuntimeT struct {	skipped bool	failed  bool}
 (t *RuntimeT) Error(args ...interface{}) {	log.Println(fmt.Sprintln(args...))	t.Fail()}
 (t *RuntimeT) Errorf(format string, args ...interface{}) {	log.Printf(format, args...)	t.Fail()}
 (t *RuntimeT) Fail() {	t.failed = true}
 (t *RuntimeT) FailNow() {	panic("testing.T failed, see logs for output (if any)")}
 (t *RuntimeT) Failed() bool {	return t.failed}
 (t *RuntimeT) Fatal(args ...interface{}) {	log.Print(args...)	t.FailNow()}
 (t *RuntimeT) Fatalf(format string, args ...interface{}) {	log.Printf(format, args...)	t.FailNow()}
 (t *RuntimeT) Log(args ...interface{}) {	log.Println(fmt.Sprintln(args...))}
 (t *RuntimeT) Logf(format string, args ...interface{}) {	log.Println(fmt.Sprintf(format, args...))}
 (t *RuntimeT) Name() string {	return ""}
 (t *RuntimeT) Parallel() {}
 (t *RuntimeT) Skip(args ...interface{}) {	log.Print(args...)	t.SkipNow()}
 (t *RuntimeT) SkipNow() {	t.skipped = true}
 (t *RuntimeT) Skipf(format string, args ...interface{}) {	log.Printf(format, args...)	t.SkipNow()}
 (t *RuntimeT) Skipped() bool {	return t.skipped}
 (t *RuntimeT) Helper() {}
 (t *RuntimeT) Cleanup(
()) {}