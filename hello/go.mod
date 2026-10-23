module example.com/hello

go 1.27.2

require (
	example.com/encodeJSON v0.0.0-00010101000000-000000000000
	example.com/greetings v0.0.0-00010101000000-000000000000
)

replace example.com/greetings => ../greetings

replace example.com/encodeJSON => ../encodeJSON
