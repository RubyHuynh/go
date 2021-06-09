module hthngoc/tiki

go 1.14

replace hthngoc/tikitiki => ./modules/

replace hthngoc/logging => ./logging/

replace tests => ./tests/

require (
	hthngoc/logging v0.0.0-00010101000000-000000000000
	hthngoc/tikitiki v0.0.0-00010101000000-000000000000
	tests v0.0.0-00010101000000-000000000000 // indirect
)
