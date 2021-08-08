# The theory of Docker Networking

At the highest level, Docker networking comprises these three major components:

- The Container Network Model(CNM)
- libnetwork
- Drivers


The CNM is the design specification. It outlines the fundamental building blocks of Docker network.

libnetwork is a real-world implementation of the CNM, and is used by Docker. It's written in GO and implements the core components
outlined in the CNM.

Drivers extend the model by implementing specific network topologies such as VXLAN overlay networks.

The following figure shows how they fit together at a very high level.

