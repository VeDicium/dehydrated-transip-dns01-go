#!/usr/bin/env bash

# Hook script for dns-01 challenge via TransIP API
# Made via Golang, but added bash script for simplification
#
# https://api.transip.nl/rest/docs.html#introduction
# https://github.com/VeDicium/dehydrated-transip-dns01-go

# Set keys
TRANSIP_ACCOUNT_NAME="transip-account-name"
TRANSIP_KEY_PATH="/full/path/to/awesome-key-pair.key"

# Test mode
# When set to 1, test mode is used
TEST_MODE="1"

# Check if keys are set, just to be sure
if [[ -z "${TRANSIP_ACCOUNT_NAME}" ]] || [[ -z "${TRANSIP_KEY_PATH}" ]]; then
  echo "TransIP credentials not set. Make sure TRANSIP_ACCOUNT_NAME and TRANSIP_KEY_PATH environment variables are set"
fi

# Export variables
export TRANSIP_ACCOUNT_NAME
export TRANSIP_KEY_PATH
export TEST_MODE

# Run script (via go)
go run main.go "$@"
