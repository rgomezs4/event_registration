#!/bin/bash

go build -o event_registration *.go
echo "binary built succesfuly ./event_registration"
zip events_be_binaries.zip event_registration ignored_paths run.sh
echo "generated release zip events_be_binaries.zip"