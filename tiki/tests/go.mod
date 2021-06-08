module testt

go 1.14

replace hthngoc/tikitiki => ../modules/

replace hthngoc/logging => ../logging/

require (
	hthngoc/logging v0.0.0-00010101000000-000000000000 // indirect
	hthngoc/tikitiki v0.0.0-00010101000000-000000000000
)
