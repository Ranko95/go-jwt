#!/bin/bash

COMPOSE_HTTP_TIMEOUT=3000 docker-compose --compatibility -f ./deploy/local/docker-compose.yml up --build --remove-orphans
