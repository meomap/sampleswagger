#!/bin/bash
swagger generate server -t gen -f ./swagger.yaml --exclude-main -A SimplePost
swagger generate client -t gen -f ./swagger.yaml -A SimplePost
