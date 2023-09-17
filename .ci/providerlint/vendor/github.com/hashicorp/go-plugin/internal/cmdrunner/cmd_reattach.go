// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package cmdrunnerimport (
"context"
"fmt"
"net"
"os""github.com/hashicorp/go-plugin/runner"
)// Reattach
 returns a 
 thatows reattaching to a plugin running
// as a n process. The process may or may not be a child process. Reattach
(pid int, addr net.Addr) runner.Reattach
 {
return 
() (runner.AttachedRunner, error) {
p, err := os.FindProcess(pid)
if err != nil {
// On Unix systems, FindProcess never returns an error.
// On Windows, for non-existent pids it returns:
// os.SyscallError - 'OpenProcess: the paremter is incorrect'
return nil, ErrProcessNotFound
}// Attempt to connect to the addr since on Unix systems FindProcess
// doesn't actually return an error if it can't find the process.
conn, err := net.Dial(addr.Network(), addr.String())
if err != nil {
p.Kill()
return nil, ErrProcessNotFound
}
conn.Close()return &CmdAttachedRunner{
pid:     pid,
process: p,
}, nil
}
}// CmdAttachedRunner is mostly a subset of CmdRunner, except the Wait 
tion
// does not assume the process is a child of the host process, and so uses a
// different implementation to wait on the process.
 CmdAttachedRunner struct {
pid     int
process *os.ProcessrTranslator
}
*CmdAttachedRunner) Wait(_ context.Context) error {
return pidWait(c.pid)
}
 (c *CmdAttachedRunner) Kill(_ context.Context) error {
return c.process.Kill()
}
 (c *CmdAttachedRunner) ID() string {
return fmt.Sprintf("%d", c.pid)
}
