#!/usr/bin/env bash

set -e

yellow="\033[1;33m"
noc="\033[0m"

printf "${yellow}INIT:${noc} 🚀 run migrations script\n"

conn_str="postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME"

/go/bin/goose -dir '/app/migrations' postgres $conn_str up