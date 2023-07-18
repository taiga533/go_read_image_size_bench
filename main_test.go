package main

import (
	"os"
	"testing"
)

// 画像をメモリ上に展開する
func TestLoadImage(t *testing.T) {
	f, err := os.Open("test.png")
	if err != nil {
		t.Errorf("ReadPngHeader() error = %v", err)
	}
	defer f.Close()

	img, err := LoadImage(f)
	if err != nil {
		t.Errorf("LoadImage() error = %v", err)
	}

	if img.Bounds().Size().X != 7000 {
		t.Errorf("LoadImage() X = %v", img.Bounds().Size().X)
	}
	if img.Bounds().Size().Y != 7000 {
		t.Errorf("LoadImage() Y = %v", img.Bounds().Size().Y)
	}
}

func TestReadPngHeader(t *testing.T) {
	f, err := os.Open("test.png")
	if err != nil {
		t.Errorf("ReadPngHeader() error = %v", err)
	}
	defer f.Close()

	pngHeader, err := ReadPngHeader(f)
	if err != nil {
		t.Errorf("ReadPngHeader() error = %v", err)
	}
	if pngHeader.Width != 7000 {
		t.Errorf("ReadPngHeader() pngHeader.Width = %v", pngHeader.Width)
	}
	if pngHeader.Height != 7000 {
		t.Errorf("ReadPngHeader() pngHeader.Height = %v", pngHeader.Height)
	}
}

func BenchmarkLoadImage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.Open("test.png")
		if err != nil {
			b.Errorf("ReadPngHeader() error = %v", err)
		}
		defer f.Close()
		LoadImage(f)
	}
}

func BenchmarkReadPngHeader(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.Open("test.png")
		if err != nil {
			b.Errorf("ReadPngHeader() error = %v", err)
		}
		defer f.Close()
		ReadPngHeader(f)
	}
}
