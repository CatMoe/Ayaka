package utils

import (
	"github.com/CatMoe/Ayaka/config"
	"github.com/CatMoe/Ayaka/define"
	"github.com/gofiber/fiber/v2"
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

var reCaptchaUrl = "https://www.recaptcha.net/recaptcha/api/siteverify?secret=" + config.RECAPTCHA_SECRET + "&response="

func Verify(responseToken string) bool {
	agent := fiber.AcquireAgent()
	request := agent.Request()
	request.SetRequestURI(reCaptchaUrl + responseToken)
	if err := agent.Parse(); err != nil {
		logrus.Error(err)
		return false
	}
	_, body, errs := agent.Bytes()
	if len(errs) != 0 {
		for _, err := range errs {
			logrus.Error(err)
			return false
		}
	}

	response := define.ReCaptchaResponse{}
	jsoniter.Unmarshal(body, &response)
	if response.Success {
		return true
	} else {
		return false
	}
}
