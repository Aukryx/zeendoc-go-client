package zeendoc

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func TestZeendocLogin(wsdlURL, login, password string) (string, error) {
	soapBody := fmt.Sprintf(`
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:urn="urn:Zeendoc">
  <soapenv:Header/>
  <soapenv:Body>
    <urn:login>
      <Login>%s</Login>
      <Password></Password>
      <CPassword>%s</CPassword>
    </urn:login>
  </soapenv:Body>
</soapenv:Envelope>`, login, password)

	req, err := http.NewRequest("POST", wsdlURL, bytes.NewBuffer([]byte(soapBody)))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	// Récupérer le cookie de session ZeenDoc
	cookies := resp.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "ZeenDoc" {
			return fmt.Sprintf("ZeenDoc=%s", cookie.Value), nil
		}
	}
	return "", fmt.Errorf("cookie ZeenDoc non trouvé")
}
