#!/usr/bin/env bash

export DB_TYPE=POSTGRES
export DB_USER=postgres
export DB_PASS='otto123'
export DB_NAME=stockbit
export DB_ADDR=postgres
export DB_PORT=5432
export DB_SSLM=disable
export DB_TIMEOUT=30

go run main.go

# nohup ./oasis_report_generator > oasis_report_generator_log.out 2>&1 &