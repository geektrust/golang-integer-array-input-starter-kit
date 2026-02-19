#!/bin/bash
go clean && go build -o app && ./app "$@"