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

# Run command
HANDLER="$1"; shift

# Fetch deploy command, because Golang doesn't do anything with it
# Use this for reload nginx for example
if [[ "${HANDLER}" =~ ^(deploy_cert|deploy_ocsp)$ ]]; then

  # deploy_cert command
  if [[ "${HANDLER}" == "deploy_cert" ]]; then
    # Add all variables
    DOMAIN="${1}"
    KEYFILE="${2}"
    CERTFILE="${3}"
    FULLCHAINFILE="${4}"
    CHAINFILE="${5}"
    TIMESTAMP="${6}"
  fi

  # deploy_ocsp command
  if [[ "${HANDLER}" == "deploy_ocsp" ]]; then
    # Add all variables
    DOMAIN="${1}"
    OCSPFILE="${2}"
    TIMESTAMP="${3}"
  fi

  # Run this command, for both command
  # Like reloading nginx
  # systemctl reload nginx

else
  # Run go when not deploy_cert or deploy_ocsp

  # Using Go
  go run main.go "$HANDLER" "$@"

  # Using built binary
  # /full/path/to/executable "$HANDLER" "$@"
fi
