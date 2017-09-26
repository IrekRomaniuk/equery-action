package action

import (
	"testing"
	"fmt"
	"net/url"
) 

func TestTicket(t *testing.T) {	
	//https://stackoverflow.com/questions/19253469/make-a-url-encoded-post-request-using-http-newrequest
	apiURL := "https://commonwealthdev.service-now.com/"
	resource := "api/now/table/u_security_engineering_request"
	data := url.Values{} 
	data.Add("short_description", "Test with CURL")
	u, _ := url.ParseRequestURI(apiURL)
	u.Path = resource
	u.RawQuery = data.Encode()
	urlStr := fmt.Sprintf("%v", u)  // "https://api.com/user/?name=foo&surname=bar"
	
	htmlData, err :=Ticket(urlStr, "netpro:u4-~B;3nd4CmzYB*")  //bytes.NewBufferString(data.Encode())
	if err != nil {
		fmt.Println(err)
	} else {
		//str := fmt.Sprintf("%s", htmlData)
		fmt.Println(string(htmlData))
	}
		
}