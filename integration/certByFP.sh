#!/bin/bash -e

cd "$(dirname "$0")"

DATA="../data/*.csv"

row=$(grep "$1" $DATA)

echo "$row" | \
  awk -F "," '{print $(NF-1)}' | \
  base64 -d | \
  openssl x509 -inform DER -outform PEM -text

echo ""
echo "+ View on Censys: https://censys.io/certificates/$1"
echo ""
