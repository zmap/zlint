module github.com/zmap/zlint/v3/cmd/genTestCerts

go 1.23.0

toolchain go1.23.6

replace github.com/zmap/zlint/v3 => ../../

require (
	github.com/zmap/zcrypto v0.0.0-20250129210703-03c45d0bae98
	github.com/zmap/zlint/v3 v3.6.5
)

require (
	github.com/weppos/publicsuffix-go v0.40.3-0.20250127173806-e489a31678ca // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/text v0.23.0 // indirect
)
