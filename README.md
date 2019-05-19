# Event Store Microservice

**ATTENTION:** this setup assumes that the code is always executed inside
a Docker container and is not meant for running code on the host machine
directly. 

To learn more: [https://justgo.rocks](https://justgo.rocks)

## Quickstart

1. Build project and start inside a container: 

    ```
    > make
    ```

1. Check logs to verify things are running properly:

    ```
    > make logs
    ```

    If you see `Starting microservice on internal port` as the last entry in 
    the log things should be A-OK. However, the port indicated there is
    internal to the Docker container and not a port you can test the service
    at. You need to run `make ps` to detect the external port (see below).

1. Find the port the server attached to by running:

   ```
   > make ps
   ```

   which will have an output that looks something like 

   ```
     Name                   Command               State            Ports
   --------------------------------------------------------------------------------
   ms-gevent-api   ./wait-for.sh -t 60 ms-gev ...   Up      0.0.0.0:32773->3535/tcp
   ```

   Whatever you see instead of `0.0.0.0:32773` is the host/port that your
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


## License 

[MIT](LICENSE)
