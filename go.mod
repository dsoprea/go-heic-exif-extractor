module github.com/dsoprea/go-heic-exif-extractor

go 1.13

// Development only
// replace github.com/dsoprea/go-utility => ../go-utility
// replace github.com/dsoprea/go-exif/v2 => ../go-exif/v2

require (
	github.com/dsoprea/go-exif/v2 v2.0.0-20200502203340-6aea10b45f4c
	github.com/dsoprea/go-logging v0.0.0-20200502201358-170ff607885f
	github.com/dsoprea/go-utility v0.0.0-20200424085841-d6691864fa10
	go4.org v0.0.0-20200411211856-f5505b9728dd
	gopkg.in/yaml.v2 v2.2.8 // indirect
)
