package main

import "fmt"

type Encoder interface {
	Encode(value string)
}

type XEncoder struct {
}

type YEncoder struct {
}

func (xEncoder *XEncoder) Encode(value string) {
	fmt.Println("Encoded with XEncoder.")
}

func (xEncoder *XEncoder) Decode(value string) {
	fmt.Println("Decoded with XEncoder.")
}

func (yEncoder *YEncoder) Encode(value string) {
	fmt.Println("Encoded with YEncoder.")
}

func (yEncoder *YEncoder) Decode(value string) {
	fmt.Println("Decoded with YEncoder.")
}

func main() {
	/*var xEncoder *XEncoder
	var yEncoder *YEncoder

	xEncoder = &XEncoder{}
	yEncoder = &YEncoder{}

	xEncoder.Encode("123456")
	xEncoder.Decode("123456")

	yEncoder.Encode("123456")
	yEncoder.Decode("123456")*/

	var encoder Encoder

	encoder = &YEncoder{}
	encoder.Encode("123456")
}
