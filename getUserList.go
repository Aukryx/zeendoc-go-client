package zeendoc

import (
	"bytes"
	"io"
	"net/http"
)

func GetUserList(wsdlURL, sessionCookie string) (string, error) {
	soapBody := `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:urn="urn:Zeendoc">
  <soapenv:Header/>
  <soapenv:Body>
    <urn:getUserList>
      <User_Id></User_Id>
      <Access_token></Access_token>
    </urn:getUserList>
  </soapenv:Body>
</soapenv:Envelope>`

	req, err := http.NewRequest("POST", wsdlURL, bytes.NewBuffer([]byte(soapBody)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	if sessionCookie != "" {
		req.Header.Set("Cookie", sessionCookie)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return string(body), nil
}
