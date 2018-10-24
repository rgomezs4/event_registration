#!/bin/bash

export APP_PORT="3158"
export APP_NAME="events"
export APP_DB_DRIVER="postgres"
export APP_DB_SOURCE="postgres://postgres:Wblake91@10.232.6.1:5432/events?sslmode=disable"
export APP_KEY="secret"

./event_registration

