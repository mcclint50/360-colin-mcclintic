module val

go 1.15

replace codingvalidator.com/val1 => ../val1

require (
	codingvalidator.com/val1 v0.0.0-00010101000000-000000000000
	codingvalidator.com/val2 v0.0.0-00010101000000-000000000000
)

replace codingvalidator.com/val2 => ../val2
