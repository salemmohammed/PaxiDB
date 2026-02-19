package main

import (
	//"encoding/binary"
	"flag"
	"github.com/salemmohammed/PaxiDB/log"
	"github.com/salemmohammed/PaxiDB/paxos"

	"github.com/salemmohammed/PaxiDB"
)

var id = flag.String("id", "", "node id this client connects to")
var algorithm = flag.String("algorithm", "", "Client API type [paxos]")
var load = flag.Bool("load", false, "Load K keys into DB")
var master = flag.String("master", "", "Master address.")
var delta = flag.Int("delta", 0, "value of delta.")


type db struct {
	PaxiDB.Client
}

func (d *db) Init() error {
	return nil
}

func (d *db) Stop() error {
	return nil
}


func (d *db) Write(k int, v []byte) error {
	key := PaxiDB.Key(k)
	//value := make([]byte, binary.MaxVarintLen64)
	//value := make([]byte, 10000)
	//binary.ByteOrder(v)
	//binary.PutUvarint(value, uint64(v))
	err := d.PutMUL(key, v)
	//err := d.Put(key, value)
	return err
}

func main() {
	PaxiDB.Init()

	if *master != "" {
		PaxiDB.ConnectToMaster(*master, true, PaxiDB.ID(*id))
	}

	d := new(db)
	switch *algorithm {
	case "paxos":
		d.Client = paxos.NewClient(PaxiDB.ID(*id))
	default:
		d.Client = PaxiDB.NewHTTPClient(PaxiDB.ID(*id))
	}

	b := PaxiDB.NewBenchmark(d)
	if *load {
		log.Debugf("Load in Clinet is started")
		b.Load()
	} else {
		log.Debugf("Run in Clinet is started")
		b.Run()
	}
}
