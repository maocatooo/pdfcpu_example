package main

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/font"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	failed() // with simsun.ttf
	//ok() with unifont-13.0.03.ttf
}

func failed() {

	conf := api.LoadConfiguration()
	fmt.Println(font.UserFontNames(), font.UserFontDir)
	fonts, err := userFonts("font")
	if err != nil {
		panic(err)
	}
	if err := api.InstallFonts(fonts); err != nil {
		panic(err)
	}
	fmt.Println(font.UserFontNames())
	err = api.FillFormFile("template/wrapper.pdf", "template/wrapper.json", "wrapper_out.pdf", conf)
	if err != nil {
		panic(err)
	}
}

func ok() {

	conf := api.LoadConfiguration()
	fmt.Println(font.UserFontNames(), font.UserFontDir)
	fonts, err := userFonts("font")
	if err != nil {
		panic(err)
	}
	if err := api.InstallFonts(fonts); err != nil {
		panic(err)
	}
	fmt.Println(font.UserFontNames())
	err = api.FillFormFile("template/chineseSimple.pdf", "template/chineseSimple.json", "chineseSimple_out.pdf", conf)
	if err != nil {
		panic(err)
	}
}

func isTrueType(filename string) bool {
	s := strings.ToLower(filename)
	return strings.HasSuffix(s, ".ttf") || strings.HasSuffix(s, ".ttc")
}

func userFonts(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	ff := []string(nil)
	for _, f := range files {
		if isTrueType(f.Name()) {
			fn := filepath.Join(dir, f.Name())
			ff = append(ff, fn)
		}
	}
	return ff, nil
}
