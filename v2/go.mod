module github.com/dsoprea/go-heic-exif-extractor/v2

go 1.12

// Development only
// replace github.com/dsoprea/go-utility/v2 => ../../go-utility/v2
// replace github.com/dsoprea/go-exif/v3 => ../../go-exif/v3

require (
	github.com/dsoprea/go-exif/v3 v3.0.0-20200717071058-9393e7afd446
	github.com/dsoprea/go-logging v0.0.0-20200517223158-a10564966e9d
	github.com/dsoprea/go-utility/v2 v2.0.0-20200717064901-2fccff4aa15e
	github.com/jdeng/goheif v0.0.0-20200323230657-a0d6a8b3e68f
	go4.org v0.0.0-20200411211856-f5505b9728dd
	golang.org/x/net v0.0.0-20200520182314-0ba52f642ac2 // indirect
)
