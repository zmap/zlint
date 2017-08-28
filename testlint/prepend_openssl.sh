#!/bin/bash
for f in testCerts/*; do
    openssl x509 -in $f -text -noout | cat - $f > /tmp/out && mv /tmp/out $f
done
