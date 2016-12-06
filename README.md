# zlint
To install this code:

If you have not already set up a GOPATH, do the following after installing golang to temporarily set one under bash:

`$ mkdir $HOME/godir`

`$ export GOPATH=$HOME/godir`

`$ export PATH=$PATH:$GOPATH/bin`

or for Windows under cmd:

`> mkdir %HOMEPATH%\godir`

`> set GOPATH=%HOMEPATH%\godir`

`> set PATH=%PATH%;%GOPATH%\bin`

Then, simply install certlint:

`$ go get github.com/zmap/zlint`
