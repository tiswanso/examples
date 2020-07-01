# Kiknos VPN Endpoint NSE
This folder contains a Helm chart for the Kiknos VPN Endpoint NSE. This document describes the most
important parts that may be needed in order to tweak the NSE deployment.


## Important Helm options
The default values for Helm options present in [values.yaml](values.yaml) file can be generally preserved,
except the following options that usually have to be tweaked for each deployment:

| Helm Option                          | Example Value                      | Description |
| ------------------------------------ | ---------------------------------- | ----------- |
| `strongswan.network.localSubnet`     | `172.31.23.0/24`                   | Local IP subnet exposed via the IPSec tunnel to the remote peers. This value will be also used as the local NSE's IPAM prefix pool. |
| `strongswan.network.remoteSubnets`   | `{172.31.23.0/24,172.31.100.0/24}` | Array of remote IP subnets accessible via IPSec tunnels. This usually matches to an array of the local subnets of individual remote IPSec peers. |
| `strongswan.network.remoteAddr`      | `81.82.123.124`                    | (optional) IP address of the remote IPSec peer. Required only if this NSE acts as an IPSec client initiating a connection to the remote gateway/peer. Leave empty if this NSE acts as an IPSec gateway. |
| `strongswan.secrets.ikePreSharedKey` | `Vpp123`                           | Pre-shared-key used to authenticate with remote IPSec peers. |


## StrongSwan configuration files
There are two StrongSwan configuration files embedded in this helm chart. Most often they don't have to be
touched, since they are parametrized using the aforementioned helm options, but the parts that are not exposed
via the helm options can still be tweaked if needed:

- [strongswan-cfg.yaml](templates/strongswan-cfg.yaml): contains generic StrongSwan configuration,
 such as the count of StrongSwan threads or timeout values,
- [responder-cfg.yaml](templates/responder-cfg.yaml): contains IPSec connection configuration,
 such as authentication and encryption algorithms.


## VPP startup configuration file
VPP startup configuration generally does not require any tweaks, but it is located in
[vpp-cfg.yaml](templates/vpp-cfg.yaml) file in case that any changes are required.
