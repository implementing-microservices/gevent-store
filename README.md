# justgo-microservice
[![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/inadarei/justgo-microservice/issues)
[![Go project version](https://badge.fury.io/go/github.com%2Finadarei%2Fjustgo-microservice.svg)](https://badge.fury.io/go/github.com%2Finadarei%2Fjustgo-microservice)
[![Go Report Card](https://goreportcard.com/badge/github.com/inadarei/justgo-microservice)](https://goreportcard.com/report/github.com/inadarei/justgo-microservice)

Microservice Template for JustGo (https://git.justgo.rocks)

Skeleton project for jump-starting a Go-powered microservice development with
Docker, code hot-reloading and Go best-practices.

The only requirement for running this project is a properly set-up Docker
environment and (optionally) GNU make. You can also run commands in the 
Makefile manually (they are simple), if you don't have make.

**ATTENTION:** this setup assumes that the code is always executed inside
a Docker container and is not meant for running code on the host machine
directly. 

To learn more: [https://justgo.rocks](https://justgo.rocks)

## Quickstart

1. Get code with [justgo]() (preferred) or by checking this repo out, locally.
2. Build project and start inside a container: 

    ```
    > make
    ```

3. Check logs to verify things are running properly:

    ```
    > make logs
    ```

    If you see `Starting microservice on internal port` as the last entry in 
    the log things should be A-OK. However, the port indicated there is
    internal to the Docker container and not a port you can test the service
    at. You need to run `make ps` to detect the external port (see below).

4. Find the port the server attached to by running:

   ```
   > make ps
   ```

   which will have an output that looks something like 

   ```
     Name                   Command               State            Ports
   --------------------------------------------------------------------------------
   ms-helloworld   CompileDaemon -build=scrip ...   Up      0.0.0.0:32770->3737/tcp
   ```

   Whatever you see instead of `0.0.0.0:32770` is the host/port that your
   microservice started at. Type it in your browser or Postman or another
   HTTP client of your choice to verify that the service is responding.

## Usage

```
# run:
> make [start]

# stop:
> make stop

# follow logs:
> make logs

# show container status:
> make ps

# jump into the Docker container
> make shell

# To add a dependency, just modify go.mod
# and hot reloader will do the rest! THAT EASY

# build a release (production-appropriate) Docker image "from scratch":
> make build-release

# Run the release build:
> make run-release

```

## Contributing
Contributions are always welcome, no matter how large or small. Substantial feature requests should be proposed as an [RFC](https://github.com/apiaryio/api-blueprint-rfcs/blob/master/template.md). Before contributing, please read the [code of conduct](https://github.com/inadarei/justgo-microservice/blob/master/CODE_OF_CONDUCT.md).

See [Contributing](CONTRIBUTING.md).

## License 

[MIT](LICENSE)
