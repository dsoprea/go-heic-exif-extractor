[![Build Status](https://travis-ci.org/dsoprea/go-heic-exif-extractor.svg?branch=master)](https://travis-ci.org/dsoprea/go-heic-exif-extractor)
[![Coverage Status](https://coveralls.io/repos/github/dsoprea/go-heic-exif-extractor/badge.svg?branch=master)](https://coveralls.io/github/dsoprea/go-heic-exif-extractor?branch=master)
[![Report Card](https://goreportcard.com/badge/github.com/dsoprea/go-heic-exif-extractor)](https://goreportcard.com/report/github.com/dsoprea/go-heic-exif-extractor)
[![GoDoc](https://godoc.org/github.com/dsoprea/go-heic-exif-extractor?status.svg)](https://godoc.org/github.com/dsoprea/go-heic-exif-extractor)

# Overview

This project invokes a third-party project to parse HEIC/HEIF ([ISO 23008-12](https://www.iso.org/standard/66067.html)) content and extract an EXIF blob and then invokes [go-exif](https://github.com/dsoprea/go-exif) to parse that EXIF blob. It satisfies the [riimage.MediaParser](https://github.com/dsoprea/go-utility/blob/master/image/media_parser_type.go) interface and exists to provide HEIC/HEIF support to the [go-exif-knife](https://github.com/dsoprea/go-exif-knife) tool.

# Examples

See the [GoDoc page](https://godoc.org/github.com/dsoprea/go-heic-exif-extractor) for [usage examples](https://godoc.org/github.com/dsoprea/go-heic-exif-extractor#pkg-examples).
