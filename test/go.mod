module test.com/test

go 1.16

require (
	rsc.io/quote v1.5.2
	test.com/greetings v0.0.0-00010101000000-000000000000
)

replace test.com/greetings => ../greetings
