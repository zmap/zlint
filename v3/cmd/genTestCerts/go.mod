module github.com/zmap/zlint/v3/cmd/genTestCerts

go 1.18

replace github.com/zmap/zlint/v3 => ../../

require (
	github.com/zmap/zcrypto v0.0.0-20220402174210-599ec18ecbac
	github.com/zmap/zlint/v3 v3.0.0-00010101000000-000000000000
)

require (
	github.com/weppos/publicsuffix-go v0.15.1-0.20220329081811-9a40b608a236 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5 // indirect
	golang.org/x/text v0.3.7 // indirect
)
