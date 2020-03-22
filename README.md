[![GoDoc](https://godoc.org/github.com/dsoprea/go-exif-heic-wrapper?status.svg)](https://godoc.org/github.com/dsoprea/go-exif-heic-wrapper)

# Overview

This project invokes a third-party project to parse HEIC/HEIF ([ISO 23008-12](https://www.iso.org/standard/66067.html)) content and extract an EXIF blob and then invokes [go-exif](https://github.com/dsoprea/go-exif) to parse that EXIF blob. It satisfies the [riimage.MediaParser](https://github.com/dsoprea/go-utility/blob/master/image/media_parser_type.go) interface and exists to provide HEIC/HEIF support to the [go-exif-knife](https://github.com/dsoprea/go-exif-knife) tool.
