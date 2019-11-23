#!/bin/bash -e

cd "$(dirname "$0")"

DATA="../data/*.csv"

row=$(grep "$1" $DATA)

echo "https://censys.io/certificates/$1"
echo ""
echo "$row" | \
  awk -F "," '{print $(NF-1)}' | \
  base64 -d | \
  openssl x509 -inform DER -outform PEM -text
