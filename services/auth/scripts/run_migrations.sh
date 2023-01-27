#!/usr/bin/env bash

set -e

echo "does it work"

conn_str="postgres://$DB_USERNAME:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME"

/go/bin/goose postgres $conn_str up