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

## Intention
Distributed networks often get confusing, so I decided to code a little helper. Since I am not a programmer, I never used gRPC and wanted to give it a try.

I like it and the helper tool "L7-Snake" is sufficient for a low-level network overview when tools like SDNs (NSX, Nuage, Midonet), Meshes (Istio, LinkerD, etc.), Monitoring (Icinga, Nagios, Datadog, etc.) or other Analysis is limited.

Especially for gated networks (subnet locks) it does serve me well.

Enjoy :)