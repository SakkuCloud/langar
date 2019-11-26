# LANGAR - A web service to work with Docker daemon

[![N|Solid](https://avatars3.githubusercontent.com/u/44247427?s=200&v=4)](https://github.com/SakkuCloud)

Langar is a web service that manage Docker daemon.
Langar (لنگر in Persian) means Anchor.

# Contents
* [Motivations](https://github.com/SakkuCloud/langar#motivations)

# Motivations
In SAKKU team we have several modules that needs to make actions in Docker daemon in several servers. Actions like managing container networks, list of containers, manage of swarm and so on. So we need a manager to listen in endpoints and make this actions in Docker daemon servers to prevent issues like:
- Opening Docker API to public networks.
- No control in actions that every module can do with Docker API.