package paxos
import (
	"github.com/salemmohammed/PaxiDB"
)
type Client struct {
	*PaxiDB.HTTPClient
	ballot PaxiDB.Ballot
}
func NewClient(id PaxiDB.ID) *Client {
	return &Client{
		HTTPClient: PaxiDB.NewHTTPClient(id),
	}
}
func (c *Client) Put(key PaxiDB.Key, value PaxiDB.Value) error {
	c.HTTPClient.CID++
	_, meta, err := c.RESTPut(c.ID, key, value)
	if err == nil {
		b := PaxiDB.NewBallotFromString(meta[HTTPHeaderBallot])
		if b > c.ballot {
			c.ballot = b
		}
	}

	return err
}