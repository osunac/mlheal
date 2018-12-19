# Self-Healing Infrastructure Example

This repository contains an example of a self-healing infrastructure using
Docker containers, a TICK stack and Loud ML.

In order to make all the containers names predictable, all the
interesting stuff is in the `mlheal` subdirectory.


## Requirements

This project requires Docker and Docker Compose.


## Usage

First of all, go to the `mlheal` subdirectory.

Then use Docker Compose to start all the containers.

```
# docker-compose up
```

This will:
- build all images for which a `Dockerfile` is provided;
- download the images for the other containers;
- create virtual network;
- start all the containers.

Some ports are exposed to the host.

Port | Description
---- | --------------------------
8079 | Controller REST API
8080 | Sample application REST API
8077 | Loud ML REST API
8086 | InfluxDB
8888 | Chronograf web interface


To stop all the containers:

```
# docker-compose down
```

If you want to remove all the resources that have been created or downloaded:

```
# docker system prune
# sudo rm -rf ./mlheal/{log,var}
```
