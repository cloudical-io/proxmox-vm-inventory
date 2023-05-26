# Proxmox Inventory Tool

This tool generates a list of vms in one or more proxmox clusters and consolidates the information into a JSON format.

## Features

### Proxmox API

Calls the Proxmox API to generate a List of VM's and their IP addresses in a given set of clusters.

At the moment can only retreive IP's of virtual machines that support cloud init.

### HTTP endpoint

The Tool Provides some http endpoints to retreive the generated JSON

| Path | HTTP Method | Form Parameters | Response |
|--|--|--|--|
|`/`| `GET` | - | Catch all returning `NotImplemented` |
|`/all`| `GET` | - | JSON containing all clusters and all VMs |
|`/cluster`| `POST` | `cluster=<CLUSTER_NAME>` | JSON containing VMs of a specific cluster |

#### Examples

```sh
# Get the whole list
curl "localhost:8080/all"

# Get the cluster specific list
curl "localhost:8080/cluster" --form-string 'cluster=MY_CLUSTER'
```

## Config

You'll have to create an config file as the `example.yaml` and fill out the values as needed.

The config file has to contain a `clusters` list in YAML format. All other options can be passed as ENV vars or as CLI Flag.

```yaml
# -- Cluster list to fetch info from
clusters:
  - # Name must be unique
    name: "CUSTOMN CLUSTER NAME"
    # API key for the proxmox cluster
    apikey: "aaaaaaaaa-bbb-cccc-dddd-ef0123456789"
    # API user
    apiuser: "root@pam!monitoring"
    # API URL
    apihost: "https://10.0.0.1:8006"
```

### CLI Flags

You can also pass most of the values as CLI Flag or Environment Variable.

Note that the Config File always has to be specified.

|Flag | Long Flag | Default Value | ENV_VAR | Usage
|--|--|--|--|--|
|-f|--config-file||INV_CLUSTER_FILE|YAML file containing your config values. Values set here override all commandline flags and environment vars|
|-l|--log-level|INFO|INV_LOG_LEVEL|Set the Log Level / verbosity|
|-t|--timeout|10|INV_TIMEOUT|Time in seconds before a request times out|
|-i|--fetch-interval|300|INV_INTERVAL|Interval at whicht to refetch all VMs|

## Build Instructions

To build the binary yourself you'll have to have `go 1.X` installed.

```sh
CGO_ENABLED=0 go mod download
CGO_ENABLED=0 go build -o pveinventory cmd/main.go
```