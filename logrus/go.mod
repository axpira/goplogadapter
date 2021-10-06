module github.com/axpira/goplogadapter/logrus

go 1.16

require (
	github.com/axpira/gop/log v0.1.3
	github.com/axpira/goplogadapter v0.0.0
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	golang.org/x/sys v0.0.0-20211004093028-2c5d950f24ef // indirect
)

replace github.com/axpira/goplogadapter => ../
