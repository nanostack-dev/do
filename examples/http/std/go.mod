module github.com/nanostack-dev/do/examples/http/std

go 1.18

replace github.com/nanostack-dev/do => ../../../

replace github.com/samber/do/http/std/v2 => ../../../http/std

require (
	github.com/samber/do/http/std/v2 v2.0.0-00010101000000-000000000000
	github.com/nanostack-dev/do v2.0.0-00010101000000-000000000000
)

require github.com/samber/go-type-to-string v1.6.1 // indirect
