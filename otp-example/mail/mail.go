package mail

import (
	"context"
	"otp-example/cache"
	utils "otp-example/utils"
	"strconv"
)

var client = cache.Client{}
var newClient = cache.NewClient()

func SendEmail(ctx context.Context, mail string) string {

	client.Client = newClient

	code := utils.RandNumber(100000, 999999)

	err := utils.SendEmail(mail, code)

	message := "E-Posta başarıyla gönderildi."

	key := "user-email:" + mail + "code:" + strconv.Itoa(code)
	err = client.Set(ctx, key, code, 180)
	if err != nil {
		return "Oluşturulan kod kaydedilirken hata meydana geldi lütfen tekrar deneyiniz"
	}
	return message
}
func VerifyEmail(ctx context.Context, mail, code string) string {

	client.Client = newClient
	key := "user-email:" + mail + "code:" + code
	isExist := client.Get(ctx, key)
	if isExist {
		return "kullanıcı doğrulaması başarılı"
	}
	return "kod eksik veya yanlış lütfen tekrar deneyiniz."

}
