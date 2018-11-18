package format

import (
	"errors"
	"github.com/nyaruka/phonenumbers"
	"regexp"
	"strconv"
	"strings"
)

func Email(email string) (ret string, err error)  {
	if len(email) == 0 {
		return "", errors.New("email is empty")
	}

	if strings.Count(email, "@") != 1 {
		return "", errors.New("bad email")
	}

	email = strings.ToLower(strings.TrimSpace(email))
	if len(email) == 0 {
		return "", errors.New("email is empty")
	}

	return email, nil
}

func Mobile(mobile string) (ret string, err error) {
	// 格式化手机，返回E164格式（其中中国手机号为保持兼容，格式化会去掉国家代码）
	if len(mobile) == 0 {
		return "", errors.New("mobile is empty")
	}

	// 国际手机号，以+开头
	if strings.HasPrefix(mobile, "+") {
		mobile = strings.Replace(mobile, " ", "", -1)
		mobile = strings.Replace(mobile, "-", "", -1)
	}

	phoneNum, err := phonenumbers.Parse(mobile, phonenumbers.UNKNOWN_REGION)
	if err != nil {
		return "", err
	}

	if phoneNum.GetCountryCode() == 86 {
		mobile = strconv.FormatUint(phoneNum.GetNationalNumber(), 10)
	} else {
		mobile = phonenumbers.Format(phoneNum, phonenumbers.E164)
	}
	return mobile, nil
}

var re_password *regexp.Regexp = regexp.MustCompile(`^[\w~!@#$%^&*()+,.:;=<>?/|\-\[\]\\]{6,20}$`)

func Password(password string) (err error) {
	l :=  len(password)
	if l == 0 {
		return errors.New("password is empty")
	}
	if l < 6 || l > 20 {
		return  errors.New("password digit error")
	}

	if !re_password.Match([]byte(password)) {
		return  errors.New("password pattern not match")
	}

	return nil
}

func IsDigit(str string) bool {
	for _, b := range []byte(str) {
		if int(b) < '0' || int(b) > '9' {
			return false
		}
	}
	return true
}