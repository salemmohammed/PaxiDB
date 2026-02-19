package main

import (
	"flag"
	"github.com/salemmohammed/PaxiDB/paxos"
	"sync"
	"github.com/salemmohammed/PaxiDB"
	"github.com/salemmohammed/PaxiDB/log"
)

var algorithm = flag.String("algorithm", "", "Distributed algorithm")
var id = flag.String("id", "", "ID in format of Zone.Node.")
var simulation = flag.Bool("sim", false, "simulation mode")
var master = flag.String("master", "", "Master address.")

func replica(id PaxiDB.ID) {
	if *master != "" {
		PaxiDB.ConnectToMaster(*master, false, id)
	}

	log.Infof("node %v starting...", id)

	switch *algorithm {
	case "paxos":
		paxos.NewReplica(id).Run()

	default:
		panic("Unknown algorithm")
	}
}
func main() {
	PaxiDB.Init()

	if *simulation {
		var wg sync.WaitGroup
		wg.Add(1)
		PaxiDB.Simulation()
		for id := range PaxiDB.GetConfig().Addrs {
			n := id
			go replica(n)
		}
		wg.Wait()
	} else {
		replica(PaxiDB.ID(*id))
	}
}
