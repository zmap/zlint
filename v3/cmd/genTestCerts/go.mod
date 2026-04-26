module github.com/zmap/zlint/v3/cmd/genTestCerts

go 1.25.0

replace github.com/zmap/zlint/v3 => ../../

require (
	github.com/zmap/zcrypto v0.0.0-20260426170728-e95752a6dfc1
	github.com/zmap/zlint/v3 v3.6.8
)

require (
	github.com/weppos/publicsuffix-go v0.50.4-0.20260424101603-5ad6bdf70b02 // indirect
	golang.org/x/crypto v0.50.0 // indirect
	golang.org/x/net v0.53.0 // indirect
	golang.org/x/text v0.36.0 // indirect
)
