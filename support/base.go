package support

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/tebeka/selenium"
)

var driver selenium.WebDriver

// WDInit retorna uma instancia do WebDriver
func WDInit() selenium.WebDriver {
	var err error
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	driver, err = selenium.NewRemote(caps, "")

	if err != nil {
		log.Fatal("Erro ao instanciar o driver:", err.Error())
	}

	driver.SetImplicitWaitTimeout(time.Second * 10)

	driver.ResizeWindow("note", 1280, 800)

	return driver
}

// SaveImage pega o print do webdriver em bytes, converte para png e savla no projeto
func SaveImage(foto []byte, name string) {
	img, _, _ := image.Decode(bytes.NewReader(foto))

	out, err := os.Create("./log/screenshots/" + name + ".png")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = png.Encode(out, img)

	if err != nil {
		fmt.Println(err)
	}
}
