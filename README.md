# Revercity
## Reverse Proxy server

## Running
* Compile with `go build`
* Run with args: `localhost:8080 7070`
* It will run with mapping newly created port `7070` into local port `8080`


## What is for?
* Could map another server port when port is changing but you want single
* Could be used in combination with `ssh -R` port forwarding to make server with low configuration run as reverse proxy

## No warranty
* Author of this product is not taking any responsibility of damage this program may cause. This program is provided __as-is__ without any warranties. Use with caution and think about security.

## Usages
* Connect with SSH (Forward local `8080` or whatever port):
```bash
# 7777           - means new port to forward on remote
# localhost:8080 - means port on YOUR CURRENT local machine
ssh -R 7777:localhost:8080 your_username@your_remote_server
```
* Then run the server.
    * Let's assume that built program name is `revercity` or `revercity.exe`
    * This example will map `8080` port to `7070` (Let's imagine that `7070` is a public port)
```bash
# Forward current 8080 to public 7070
./reversity localhost:8080 7070

# Forward google site as 7070 port
./reversity google.com:80 7070

# Forward minecraft server port as 7070
./reversity localhost:25565 7070
```