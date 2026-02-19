package paxos
import (
	"encoding/gob"
	"fmt"
	"github.com/salemmohammed/PaxiDB"
)
func init() {
	gob.Register(P1a{})
	gob.Register(P1b{})
	gob.Register(P2a{})
	gob.Register(P2b{})
	gob.Register(P3{})
}
type P1a struct {
	Ballot PaxiDB.Ballot
}
func (m P1a) String() string {
	return fmt.Sprintf("P1a {b=%v}", m.Ballot)
}
type CommandBallot struct {
	Command PaxiDB.Command
	Ballot  PaxiDB.Ballot
}
func (cb CommandBallot) String() string {
	return fmt.Sprintf("cmd=%v b=%v", cb.Command, cb.Ballot)
}
type P1b struct {
	Ballot PaxiDB.Ballot
	ID     PaxiDB.ID               // from node id
	Log    map[int]CommandBallot // uncommitted logs
}
func (m P1b) String() string {
	return fmt.Sprintf("P1b {b=%v id=%s log=%v}", m.Ballot, m.ID, m.Log)
}
type P2a struct {
	Ballot  PaxiDB.Ballot
	Slot    int
	Command PaxiDB.Command
}
func (m P2a) String() string {
	return fmt.Sprintf("P2a {b=%v s=%d cmd=%v}", m.Ballot, m.Slot, m.Command)
}
type P2b struct {
	Ballot PaxiDB.Ballot
	ID     PaxiDB.ID // from node id
	Slot   int
}
func (m P2b) String() string {
	return fmt.Sprintf("P2b {b=%v id=%s s=%d}", m.Ballot, m.ID, m.Slot)
}
type P3 struct {
	Ballot  PaxiDB.Ballot
	Slot    int
	Command PaxiDB.Command
}
func (m P3) String() string {
	return fmt.Sprintf("P3 {b=%v s=%d cmd=%v}", m.Ballot, m.Slot, m.Command)
}
