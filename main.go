package main

import (
	"encoding/binary"
	"image"
	_ "image/png"
	"io"
)

// 画像をメモリ上に展開する
func LoadImage(r io.Reader) (image.Image, error) {
	img, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// PngHeader はPNGファイルのヘッダーが持つ情報を構造体にまとめたもの
type PngHeader struct {
	Width  uint32
	Height uint32
}

// ReadPngHeader はPNGファイルのヘッダーを読み込む
func ReadPngHeader(r io.Reader) (PngHeader, error) {
	rawHeader := make([]byte, 24)
	_, err := r.Read(rawHeader)
	if err != nil {
		return PngHeader{}, err
	}
	pngHeader := PngHeader{
		Width:  binary.BigEndian.Uint32(rawHeader[16:20]),
		Height: binary.BigEndian.Uint32(rawHeader[20:24]),
	}
	return pngHeader, nil
}

func main() {
	// メインの処理
}
