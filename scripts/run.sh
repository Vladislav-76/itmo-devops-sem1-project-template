#!/bin/bash

APP_NAME="project_sem"

export DB_USER="validator"
export DB_PASSWORD="val1dat0r"
export DB_NAME="project-sem-1"
export DB_HOST="localhost"
export DB_PORT="5432"

go mod tidy
go build -o $APP_NAME
./$APP_NAME