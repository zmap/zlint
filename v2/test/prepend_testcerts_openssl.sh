#!/bin/bash

set -e -o pipefail

BASE_DIR=$(dirname "$0")
CERTS_DIR="$BASE_DIR/../testdata"
TMP_DIR=$(mktemp -d -t zlint-XXXX)

# Trap EXIT to cleanup the TMP_DIR
trap '{ rmdir --ignore-fail-on-non-empty $TMP_DIR; }' EXIT

# For every .pem file in the $CERTS directory, prepend 0penSSL text output if
# required.
for f in "$CERTS_DIR"/*.pem; do
  # Skip any files that don't begin with a PEM header. These are assumed to
  # already have the OpenSSL text output prepended.
  if [[ ! $(head -n1 "$f") =~ "-----BEGIN" ]]; then
    continue
  fi

  # If an argument is provided only consider filenames that match the provided
  # argument. This allows only prepending a specific testcert instead of all
  # unprepended testcerts.
  CERT_NAME=$(basename "$f")
  if [[ -n "$1" && ! $CERT_NAME =~ $1 ]]; then
    continue
  fi

  # If the certificate has errors parsing with OpenSSL print a warning to stderr
  # and continue. Sometimes our test data is too weird to parse and that's OK.
  if ! openssl x509 -in "$f" -noout || false; then
    echo "error parsing $f with OpenSSL" >&2
    continue
  fi

  # Prepend the test cert with its -text OpenSSL output.
  openssl x509 -text -in "$f" -outform PEM -out "$TMP_DIR/$CERT_NAME.new" \
    && mv "$TMP_DIR/$CERT_NAME.new" "$f"
done
