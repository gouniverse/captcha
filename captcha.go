package captcha

import "github.com/mojocn/base64Captcha"

// DigitsBase64ImageWithDigitsOnly generates a Base64 encoded image
func NewBase64ImageWithDigitsOnly(numericString string, height int, width int) string {
	b64driver := base64Captcha.NewDriverDigit(height, width, 5, 0.7, 80)
	b64c, _ := b64driver.DrawCaptcha(numericString)
	b64s := b64c.EncodeB64string()
	return b64s
}
