package yamuximport (
	"fmt"
	"net"
)// hasAddr is used to get the address from the underlying connection
type hasAddr interface {
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
}// yamuxAddr is used when we cannot get the underlying address
type yamuxAddr struct {
	Addr string
}
 (*yamuxAddr) Network() string {
	return "yamux"
}
 (y *yamuxAddr) String() string {
	return fmt.Sprintf("yamux:%s", y.Addr)
}// Addr is used to get the address of the listener. (s *Session) Addr() net.Addr {
	return s.LocalAddr()
}// LocalAddr is used to get the local address of the
// underlying connection. (s *Session) LocalAddr() net.Addr {
	addr, ok := s.conn.(hasAddr)
	if !ok {
		return &yamuxAddr{"local"}
	}
	return addr.LocalAddr()
// RemoteAddr is used to get the address of remote end
// of the underlying connection (s *Session) RemoteAddr() net.Addr {
	addr, ok := s.conn.(hasAddr)
	if !ok {
		return &yamuxAddr{"remote"}	return addr.RemoteAddr()
}// LocalAddr returns the local address (s *Stream) LocalAddr() net.Addr {
	return s.session.LocalAddr()
}// LocalAddr returns the remote address (s *Stream) RemoteAddr() net.Addr {
	return s.session.RemoteAddr()
}
