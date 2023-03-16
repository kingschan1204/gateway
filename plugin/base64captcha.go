package plugin

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var store = base64Captcha.DefaultMemStore

// CaptchaGenerate
func CaptchaGenerate() (id, base64 string, err error) {
	driverString := base64Captcha.DriverString{
		Height:          54,
		Width:           145,
		NoiseCount:      0,
		ShowLineOptions: 3,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: nil,
	}
	//captcha := base64Captcha.NewCaptcha(driver, store)
	driver := driverString.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	if err != nil {
		return "", "", err
	}
	return id, b64s, nil
}

// CaptchaValidate
func CaptchaValidate(id string, code string) bool {
	if store.Verify(id, code, true) {
		return true
	}
	return false
}
