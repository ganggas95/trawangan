package job

import (
	"fmt"
	"net/http"
	"time"
)

func ReadHttpBody(response *http.Response) string {

	fmt.Println("Reading body")

	bodyBuffer := make([]byte, 5000)
	var str string

	count, err := response.Body.Read(bodyBuffer)

	for ; count > 0; count, err = response.Body.Read(bodyBuffer) {
		if err != nil {
			panic(err)
		}
		str += string(bodyBuffer[:count])
	}
	return str

}

func SendToken(email, token string) error {

	url := "https://trawangan.herokuapp.com/verifyaccount?code=" + token
	isi := "<html><body><h1>Hello World!</h1><a href='" + url + "'>Click Me</a></body></html>"
	err := Send(email, isi)

	if err != nil {
		return err
	}
	fmt.Println("Message sent!")
	return nil
}
func TokenExpiry() time.Time {
	now := time.Now()
	ed, _ := time.Parse(time.UnixDate, now.Format(time.UnixDate))
	threeDays := time.Hour * 24 * 1

	diff := ed.Add(threeDays)
	timess, err := time.Parse(time.UnixDate, diff.Format(time.UnixDate))
	if err != nil {
		fmt.Println(err)
	}

	return timess
}
