module github.com/zmap/zlint/v3/cmd/gen_test_crl

go 1.23.0

toolchain go1.23.6

replace github.com/zmap/zlint/v3 => ../../

require (
	github.com/sirupsen/logrus v1.9.3
	github.com/zmap/zcrypto v0.0.0-20240803002437-3a861682ac77
	github.com/zmap/zlint/v3 v3.6.4
)

require (
	github.com/weppos/publicsuffix-go v0.40.2 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
)
