package zeendoc

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadDoc(wsdlURL, sessionCookie, collID, fileName, hash, sourceID, indexation string, pdfPath string) (string, error) {
	pdfBytes, err := os.ReadFile(pdfPath)
	if err != nil {
		return "", err
	}
	base64Doc := base64.StdEncoding.EncodeToString(pdfBytes)

	soapBody := fmt.Sprintf(`
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:urn="urn:Zeendoc">
  <soapenv:Header/>
  <soapenv:Body>
    <urn:uploadDoc>
      <Coll_Id>%s</Coll_Id>
      <Source_Id>%s</Source_Id>
      <Base64_Document>%s</Base64_Document>
      <Filename>%s</Filename>
      <Hash>%s</Hash>
      <Indexation>%s</Indexation>
      <Folder></Folder>
      <Access_token></Access_token>
    </urn:uploadDoc>
  </soapenv:Body>
</soapenv:Envelope>`, collID, sourceID, base64Doc, fileName, hash, indexation)

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
