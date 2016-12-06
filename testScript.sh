# Script to recompile both certlint and testlint then run testlint
go install github.com/teamnsrg/zlint/zlint
cd testlint
testlint -in-file=UnitTests.csv
cd ..
