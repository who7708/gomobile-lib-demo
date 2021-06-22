module Hello

go 1.16

require (
	github.com/go-resty/resty/v2 v2.6.0
	golang.org/x/mobile v0.0.0-20210614202936-7c8f154d1008 // indirect
)

replace golang.org/x/mobile v0.0.0-20210614202936-7c8f154d1008 => github.com/ydnar/gomobile v0.0.0-20210301201239-fb6ffafc9ef9
