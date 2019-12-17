# LANGAR - A web service to work with Docker daemon

Langar is a web service that manage Docker daemon.
Langar (لنگر in Persian) means Anchor.

# Contents
* [Motivations](https://github.com/SakkuCloud/langar#motivations)
* [How to use](https://github.com/SakkuCloud/langar#how-to-use)
* [Configuration](https://github.com/SakkuCloud/langar#configuration)
* [Endpoints](https://github.com/SakkuCloud/langar#endpoints)
  * [Get all networks](https://github.com/SakkuCloud/langar#get-all-networks)
  * [Get network info](https://github.com/SakkuCloud/langar#get-network-info)
  * [Create network](https://github.com/SakkuCloud/langar#create-network)
  * [Delete network](https://github.com/SakkuCloud/langar#delete-network)
* [To do](https://github.com/SakkuCloud/langar#to-do)

# Motivations
In SAKKU team we have several modules that needs to make actions in Docker daemon in several servers. Actions like managing container networks, list of containers, manage of swarm and so on. So we need a manager to listen in endpoints and make this actions in Docker daemon servers to prevent issues like:
- Opening Docker API to public networks.
- No control in actions that every module can do with Docker API.

# How to use
Building from source:
```sh
$ go build -o /usr/bin/langar github.com/SakkuCloud/langar
```

Running with CLI (stay foreground):
```sh
$ langar
```

Running in debug mode:
```sh
$ langar -debug=true
```

Langar uses some options. These options are listed below.

key   | name        | default                |
----- | ----------- | ---------------------- |
debug | Debug Mode  | false                  |
| c     | Config File | /etc/langar/config.yml |

For production, it's better to use a Systemd service to run Langar.
A simple Systemd service shown below. Save this in `/lib/systemd/system/langar.service` 
> ```sh
> [Unit]
> Description=LANGAR - A web service that manage Docker daemon
> After=network.target
>
> [Service]
> Type=simple
> Restart=on-failure
> TimeoutStopSec=10
> RestartSec=5
> ExecStart=/usr/bin/langar
>
> [Install]
> WantedBy=multi-user.target
>```

Run and enable service:
```sh
$ systemctl enable langar
$ systemctl start langar
```

# Configuration
Langar uses both YAML format and OS Environment for config. You can see [config.yml.example](https://github.com/SakkuCloud/langar/blob/master/config.yml.example) for a sample config file.
You can pass config file with:
```sh
langar -c config.yml
```
Below table describes available config file.

| config         | env                   | required | default             | describe |
| ---------------| --------------------- | :------: | ------------------- | ------------------------------------------------------- |
| port           | LANGAR_PORT           | NO       | 3000                | server will run on this port                           |
| logfile        | LANGAR_LOGFILE        | NO       | /var/log/langar.log | logs will store in this file                           |
| sentrydsn      | LANGAR_SENTRYDSN      | NO       |                     | DSN of Sentry                                          |
| accesskey      | LANGAR_ACCESSKEY      | YES      |                     | value of service http header to authorize requests     |
| secretkey      | LANGAR_SECRETKEY      | YES      |                     | value of service-key http header to authorize requests |
| docker host    | LANGAR_DOCKER_SOCKET  | NO       | /run/docker.sock    | Docker socket address                                  |


# Endpoints
### Get all networks
Get all networks in docker daemon (and swarm cluster). A complete curl requests shown below.
```sh
curl -X GET \
 http://127.0.0.1:3000/api/network \
 -H 'service: my-awesome-accesss' \
 -H 'service-key: Super$3crT' 
```

### Get network info
Get network info in docker daemon (and swarm cluster). A complete curl requests shown below. you can get network by Id or Name.
```sh
curl -X GET \
 http://127.0.0.1:3000/api/network/1234 \
 -H 'service: my-awesome-accesss' \
 -H 'service-key: Super$3crT' 
```

### Create network
Create network info in docker daemon (and swarm cluster). A complete curl requests shown below.
```sh
curl -X POST \
 http://127.0.0.1:3000/api/network \
 -H 'Content-Type: application/json' \
 -H 'service: my-awesome-accesss' \
 -H 'service-key: Super$3crT' \
 -d '{
"Name":"test-net",
"CheckDuplicate":true,
"Attachable":true,
"Driver":"overlay",
"IPAM":{
"Driver": "default",
"Config":[
{
"Subnet":"172.20.0.0/16",
"IPRange":"172.20.10.0/24",
"Gateway":"172.20.10.11"
}
]
},
"Internal":true
}
'
```

### Delete network
Delete network in docker daemon (and swarm cluster). A complete curl requests shown below. you can delete network by Id or Name.
```sh
curl -X DELETE \
 http://127.0.0.1:3000/api/network/1234 \
 -H 'service: my-awesome-accesss' \
 -H 'service-key: Super$3crT' 
```


# To do
* Add container to network
* Remove container to network
