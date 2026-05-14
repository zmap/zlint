module github.com/zmap/zlint/v3/cmd/genTestCerts

go 1.25.0

replace github.com/zmap/zlint/v3 => ../../

require (
	github.com/zmap/zcrypto v0.0.0-20260514033604-a1159eb3cad9
	github.com/zmap/zlint/v3 v3.6.8
)

require (
	github.com/weppos/publicsuffix-go v0.50.4-0.20260507075217-1bd47f85b3da // indirect
	golang.org/x/crypto v0.51.0 // indirect
	golang.org/x/net v0.54.0 // indirect
	golang.org/x/text v0.37.0 // indirect
)
