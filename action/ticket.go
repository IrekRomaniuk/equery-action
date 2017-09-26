package action 

import (
	"strings"
	"io/ioutil"
	"net/http"
	"crypto/tls"
)
// Ticket to post
func Ticket(url, cred string) ([]byte, error) {  //data io.Reader
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return []byte{}, err
	}
	credentials := strings.Split(cred, ":")
	req.SetBasicAuth(credentials[0], credentials[1])
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	//req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	resp.Body.Close()
	return htmlData, nil
}

