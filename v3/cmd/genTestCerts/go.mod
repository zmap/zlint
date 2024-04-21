module github.com/zmap/zlint/v3/cmd/genTestCerts

go 1.18

replace github.com/zmap/zlint/v3 => ../../

require (
	github.com/zmap/zcrypto v0.0.0-20230310154051-c8b263fd8300
	github.com/zmap/zlint/v3 v3.0.0
)

require (
	github.com/weppos/publicsuffix-go v0.30.0 // indirect
	golang.org/x/crypto v0.21.0 // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
