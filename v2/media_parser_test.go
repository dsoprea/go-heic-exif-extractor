package heicexif

import (
	"fmt"
	"io"
	"os"
	"path"
	"testing"

	"io/ioutil"

	"github.com/dsoprea/go-exif/v3"
	"github.com/dsoprea/go-logging"
)

const (
	testImageRelFilepathWithExif    = "image1-exif.heic"
	testImageRelFilepathWithoutExif = "cheers_1440x960-noexif.heic"
)

func TestNewHeicExifMediaParser(t *testing.T) {
	NewHeicExifMediaParser()
}

func TestHeicExifMediaParser_LooksLikeFormat__Hit(t *testing.T) {
	filepath := path.Join(assetsPath, testImageRelFilepathWithExif)

	f, err := os.Open(filepath)
	log.PanicIf(err)

	defer f.Close()

	data := make([]byte, MinimumHeicStreamLengthForDetection)

	_, err = io.ReadFull(f, data)
	log.PanicIf(err)

	hemp := new(HeicExifMediaParser)

	doesLookLike := hemp.LooksLikeFormat(data)
	if doesLookLike != true {
		t.Fatalf("HEIC image was not detected correctly.")
	}
}

func TestHeicExifMediaParser_LooksLikeFormat__Miss__EmptyData(t *testing.T) {
	data := make([]byte, MinimumHeicStreamLengthForDetection)

	hemp := new(HeicExifMediaParser)

	doesLookLike := hemp.LooksLikeFormat(data)
	if doesLookLike != false {
		t.Fatalf("Empty image data was misidentified as an HEIC image.")
	}
}

func TestHeicExifMediaParser_Parse_Hit(t *testing.T) {
	filepath := path.Join(assetsPath, testImageRelFilepathWithExif)

	f, err := os.Open(filepath)
	log.PanicIf(err)

	defer f.Close()

	hemp := new(HeicExifMediaParser)

	mc, err := hemp.Parse(f, 0)
	log.PanicIf(err)

	_, _, err = mc.Exif()
	log.PanicIf(err)
}

func TestHeicExifMediaParser_Parse_Miss(t *testing.T) {
	filepath := path.Join(assetsPath, testImageRelFilepathWithoutExif)

	f, err := os.Open(filepath)
	log.PanicIf(err)

	defer f.Close()

	hemp := new(HeicExifMediaParser)

	_, err = hemp.Parse(f, 0)
	if err == nil {
		t.Fatalf("Expected no-EXIF error.")
	} else if err != exif.ErrNoExif {
		log.Panic(err)
	}
}

func TestHeicExifMediaParser_ParseFile_Hit(t *testing.T) {
	filepath := path.Join(assetsPath, testImageRelFilepathWithExif)

	hemp := new(HeicExifMediaParser)

	mc, err := hemp.ParseFile(filepath)
	log.PanicIf(err)

	_, _, err = mc.Exif()
	log.PanicIf(err)
}

func TestHeicExifMediaParser_ParseFile_Miss(t *testing.T) {
	filepath := path.Join(assetsPath, testImageRelFilepathWithoutExif)

	hemp := new(HeicExifMediaParser)

	_, err := hemp.ParseFile(filepath)
	if err == nil {
		t.Fatalf("Expected no-EXIF error.")
	} else if err != exif.ErrNoExif {
		log.Panic(err)
	}
}

func TestHeicExifMediaParser_ParseBytes_Hit(t *testing.T) {
	filepath := path.Join(assetsPath, testImageRelFilepathWithExif)

	b, err := ioutil.ReadFile(filepath)
	log.PanicIf(err)

	hemp := new(HeicExifMediaParser)

	mc, err := hemp.ParseBytes(b)
	log.PanicIf(err)

	_, _, err = mc.Exif()
	log.PanicIf(err)
}

func TestHeicExifMediaParser_ParseBytes_Miss(t *testing.T) {
	filepath := path.Join(assetsPath, testImageRelFilepathWithoutExif)

	b, err := ioutil.ReadFile(filepath)
	log.PanicIf(err)

	hemp := new(HeicExifMediaParser)

	_, err = hemp.ParseBytes(b)
	if err == nil {
		t.Fatalf("Expected no-EXIF error.")
	} else if err != exif.ErrNoExif {
		log.Panic(err)
	}
}

func ExampleHeicExifMediaParser_LooksLikeFormat() {
	filepath := path.Join(assetsPath, testImageRelFilepathWithExif)

	f, err := os.Open(filepath)
	log.PanicIf(err)

	defer f.Close()

	data := make([]byte, MinimumHeicStreamLengthForDetection)

	_, err = io.ReadFull(f, data)
	log.PanicIf(err)

	hemp := new(HeicExifMediaParser)

	doesLookLike := hemp.LooksLikeFormat(data)
	fmt.Printf("%v\n", doesLookLike)

	// Output:
	// true
}

func ExampleHeicExifMediaParser_Parse() {
	filepath := path.Join(assetsPath, testImageRelFilepathWithExif)

	f, err := os.Open(filepath)
	log.PanicIf(err)

	defer f.Close()

	hemp := new(HeicExifMediaParser)

	mc, err := hemp.Parse(f, 0)
	log.PanicIf(err)

	rootIfd, _, err := mc.Exif()
	log.PanicIf(err)

	fmt.Println(rootIfd)

	// Output:
	// Ifd<ID=(0) IFD-PATH=[IFD] INDEX=(0) COUNT=(4) OFF=(0x0008) CHILDREN=(0) PARENT=(0x0000) NEXT-IFD=(0x0000)>
}

func ExampleHeicExifMediaParser_ParseFile() {
	filepath := path.Join(assetsPath, testImageRelFilepathWithExif)

	hemp := new(HeicExifMediaParser)

	mc, err := hemp.ParseFile(filepath)
	log.PanicIf(err)

	rootIfd, _, err := mc.Exif()
	log.PanicIf(err)

	fmt.Println(rootIfd)

	// Output:
	// Ifd<ID=(0) IFD-PATH=[IFD] INDEX=(0) COUNT=(4) OFF=(0x0008) CHILDREN=(0) PARENT=(0x0000) NEXT-IFD=(0x0000)>
}

func ExampleHeicExifMediaParser_ParseBytes() {
	filepath := path.Join(assetsPath, testImageRelFilepathWithExif)

	b, err := ioutil.ReadFile(filepath)
	log.PanicIf(err)

	hemp := new(HeicExifMediaParser)

	mc, err := hemp.ParseBytes(b)
	log.PanicIf(err)

	rootIfd, _, err := mc.Exif()
	log.PanicIf(err)

	fmt.Println(rootIfd)

	// Output:
	// Ifd<ID=(0) IFD-PATH=[IFD] INDEX=(0) COUNT=(4) OFF=(0x0008) CHILDREN=(0) PARENT=(0x0000) NEXT-IFD=(0x0000)>
}
