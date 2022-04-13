# l7-snake
Simple gRPC L7 tracer for network troubleshooting.

## How to run
- Compile the binary with (standard go)
- Run it from the CLI
- A default config gets generated (you can kill the app anytime and adjust the config)
  - Config
    - ID = Name of your node
    - LISTENPORT = Port your node listens to
    - TARGETS = List of targets you want to poke (format IP:PORT)
    - ROUTES = Arbitrary tags for human readable tracing
    - TERMINATOR = Decides if your node is the last link of the chain or not
    - INTERVAL = Poking interval (format 1ms,1s,1m,1h)
  - Flags
    - By using the flag "-config" you can set a path for your config file

## Intention
Distributed networks often get confusing, so I decided to code a little helper. Since I am not a programmer, I never used gRPC and wanted to give it a try.

I like it and the helper tool "L7-Snake" is sufficient for a low-level network overview when tools like SDNs (NSX, Nuage, Midonet), Meshes (Istio, LinkerD, etc.), Monitoring (Icinga, Nagios, Datadog, etc.) or other analysis is limited.

Especially for gated networks (subnet locks) it does serve me well.

Enjoy :)

## Further improvement

Everyone is welcome to submit issues.

Further tasks are:
- add SSL/TLS
- cleanup output for multi-target clients
- add Frontend
- implement live route changes (add,delete,change) with REST
- add database connector for timebased analysis
- add file based logging
- add logrotation based on filesize

## Infos

On Windows I built with:
- CGO\_ENABLED=1
- go build -ldflags='-s -w -extldflags "-static"' main.go

On Linux I built with:
- CGO\_ENABLED=0
- go build -a -tags netgo --ldflags '-extldflags "-static"'

For binary compression I used:
- upx --best l7-snake

Hint: _On Windows the compression flag --brute or --ultra-brute throw a false positive with Windows Defender, so I would only use the --best flag._

Protobuf build:
- protoc --go_out=. --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false protoraw/status.proto

## Docker

Inside the Docker directory there is a pre-defined Dockerfile that works on all plattforms. Make sure your binary has no dynamic linking. We only like static linked binaries.

e.g. (after you have built the binary and place it besides the Dockerfile):
- docker build -t l7-snake:v1 .

This will create a local docker image that is ready to use or to push to other registries.

## Kubernetes

Inside the kubernetes directory you can run the create-deployment.sh to create a default deployment, configmap and service.

e.g.:
- ./create-deployment.sh mysnake

This will create a file called "mysnake-dpl.yml" that is ready to be deployed with "kubectl apply -f mysnake-dpl.yml".

Make sure you adjust the imagepath to reflect your environment.

## Schematic
![Alt-Text](./pictures/example2.png)
