# PaxiDB

**PaxiDB** is a modular and extensible framework written in Go for implementing, benchmarking, and evaluating consensus protocols. It provides a unified environment for testing and comparing protocols under standardized conditions.

## Overview

This framework helps researchers and engineers to:

- Implement new consensus protocols or extend existing ones.
- Benchmark different protocols using the same environment.
- Analyze performance metrics like latency, throughput, and commit rate.

## Implemented Protocols

- Paxos

Each protocol is implemented as a pluggable module with a standard interface.

## Requirements

- Go 1.15+

## Installation

```bash
git clone https://github.com/salemmohammed/Distributed_DB.git
cd Distributed_DB
```

## Build

Run the build script from inside the `bin/` directory:

```bash
cd bin
./build.sh
```

This compiles the server, client, and cmd binaries.

## Configuration

Before running, make sure `config.json` is present in the directory you run the server from. A sample config is provided in `bin/config.json`. Copy it to the repo root:

```bash
cp bin/config.json .
```

The config defines node addresses, HTTP endpoints, and benchmark settings:

```json
{
    "address": {
        "1.1": "tcp://127.0.0.1:1735",
        "1.2": "tcp://127.0.0.1:1736",
        "1.3": "tcp://127.0.0.1:1737",
        "1.4": "tcp://127.0.0.1:1738"
    },
    "http_address": {
        "1.1": "http://127.0.0.1:8080",
        "1.2": "http://127.0.0.1:8081",
        "1.3": "http://127.0.0.1:8082",
        "1.4": "http://127.0.0.1:8083"
    },
    "policy": "majority",
    "threshold": 3,
    "benchmark": {
        "T": 60,
        "N": 0,
        "K": 1000,
        "W": 1,
        "Concurrency": 1,
        "Distribution": "uniform"
    }
}
```

## Usage

Run a server node from the repo root:

```bash
go run ./server/server.go -id <node_id> -algorithm <protocol>
```

**Flags:**

- `-id`: Node ID matching one of the IDs in `config.json` (e.g., `1.1`, `1.2`)
- `-algorithm`: One of `pbft`, `paxos`, `tendermint`, `streamlet`

**Example — running a 4-node Paxos cluster:**

Open 4 terminals and run one command in each:

```bash
# Terminal 1
go run ./server/server.go -id 1.1 -algorithm paxos

# Terminal 2
go run ./server/server.go -id 1.2 -algorithm paxos

# Terminal 3
go run ./server/server.go -id 1.3 -algorithm paxos

# Terminal 4
go run ./server/server.go -id 1.4 -algorithm paxos
```

Then run the client:

```bash
go run ./client/client.go
```

## Publications

This framework is featured in:

- **Bottlenecks in Blockchain Consensus Protocols**  
  [IEEE Xplore](https://ieeexplore.ieee.org/document/9524210)

## Contributing

We welcome contributions! Please open an issue or submit a pull request to get started.
