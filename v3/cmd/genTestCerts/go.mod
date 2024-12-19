module github.com/zmap/zlint/v3/cmd/genTestCerts

go 1.18

replace github.com/zmap/zlint/v3 => ../../

require (
	github.com/zmap/zcrypto v0.0.0-20240803002437-3a861682ac77
	github.com/zmap/zlint/v3 v3.0.0
)

require (
	github.com/weppos/publicsuffix-go v0.40.2 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)
