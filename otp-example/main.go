package main

import (
	"context"
	"fmt"
	mails "otp-example/mail"
	utils "otp-example/utils"
)

func main() {

	ctx := context.TODO()
	sendMail := mails.SendEmail(ctx, "sametavc05@gmail.com")
	fmt.Println("result:", sendMail)

	input := utils.GetInput()
	fmt.Println(mails.VerifyEmail(ctx, "sametavc05@gmail.com", input))

}
