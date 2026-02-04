module github.com/zmap/zlint/v3/cmd/gen_test_crl

go 1.24.0

replace github.com/zmap/zlint/v3 => ../../

require (
	github.com/sirupsen/logrus v1.9.3
	github.com/zmap/zcrypto v0.0.0-20251227215108-5ca1211d486b
	github.com/zmap/zlint/v3 v3.6.8
)

require (
	github.com/weppos/publicsuffix-go v0.50.1 // indirect
	golang.org/x/crypto v0.46.0 // indirect
	golang.org/x/net v0.48.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/text v0.32.0 // indirect
)
