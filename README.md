# fargate-default-backend

A trivial web service that returns http code 200 for the root (/) and /health, but 404 for all others. The purpose is to be the default container used by our supported fargate templating. Running this simple container will allow for validation and troubleshooting of the fargate service. It is available in linux for amd64 and arm64. 

The container is hosted in the public github repository:

`docker pull ghcr.io/warnermedia/fargate-default-backend:latest`

Configuration via the following env variables (all are optional):

* PORT (default: 8001) - the port to listen on
* ENABLE_LOGGING (default: false) - set to true to log requests to stdout 

# About
Copyright (c) Warner Media, LLC. All rights reserved. Licensed under the MIT license.
See the LICENSE file for license information.
