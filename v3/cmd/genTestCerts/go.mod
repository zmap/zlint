module github.com/zmap/zlint/v3/cmd/genTestCerts

go 1.18

require (
	github.com/zmap/zcrypto v0.0.0-20210811211718-6f9bc4aff20f
	github.com/zmap/zlint/v3 v3.0.0
)

require (
	github.com/weppos/publicsuffix-go v0.15.1-0.20210807195340-dc689ff0bb59 // indirect
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110 // indirect
	golang.org/x/text v0.3.5 // indirect
)

replace github.com/zmap/zlint/v3 => ../../
