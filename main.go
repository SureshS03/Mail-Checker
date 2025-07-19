package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/mail"
	"net/smtp"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <email-address>")
		return
	}

	email := os.Args[1]

	if err := SyntaxChecker(email); err != nil {
		fmt.Println(err)
		return
	}

	host, err := CheckDomain(email)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := CheckPublicData(email); err != nil {
		fmt.Println(err)
		return
	}

	if err := SmtpPing(host, email); err != nil {
		fmt.Println(err)
		return
	}
}

func SyntaxChecker(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email syntax looks good")
	return nil
}

func CheckDomain(mail string) (string, error) {
	domain := strings.Split(mail, "@")[1]
	res, err := net.LookupMX(domain)
	if err != nil || len(res) == 0 {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Domain is exits")
	return res[0].Host, nil
}

func CheckPublicData(email string) error {
	url := "https://breachdirectory.p.rapidapi.com/?func=auto&term=" + email
	fmt.Println("Checking any data breach for:", email)

	api := os.Getenv("API")
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-key", api)
	req.Header.Add("x-rapidapi-host", "breachdirectory.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var response map[string]interface{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("JSON parsing error:", err)
		return err
	}

	foundVal, ok := response["found"].(float64)
	if !ok || foundVal == 0 {
		fmt.Println("No data breach found for this email.")
		return nil
	}

	fmt.Printf("Email found in %v public breaches\n", int(foundVal))

	rawResults, ok := response["result"]
	if !ok {
		return fmt.Errorf("'result' field missing in response")
	}

	results, ok := rawResults.([]interface{})
	if !ok {
		return fmt.Errorf("'result' is not a list")
	}

	for i, item := range results {
		record, ok := item.(map[string]interface{})
		if !ok {
			fmt.Printf("Record %d is not a map\n", i)
			continue
		}

		source, exists := record["sources"]
		if !exists {
			fmt.Printf("No 'sources' field in record %d\n", i)
			continue
		}

		fmt.Printf("Mail data found in: %v\n", source)
	}

	fmt.Printf("Since email is found in %v breaches, it likely exists: %s\n", int(foundVal), email)
	fmt.Printf("Email exists in data breaches: %s\n", email)
	return nil
}

func SmtpPing(host, email string) error {

	conn, err := smtp.Dial(host + ":25")
	if err != nil {
		return fmt.Errorf("SMTP dial failed: %v", err)
	}
	defer conn.Quit()

	if err = conn.Hello("example.com"); err != nil {
		return fmt.Errorf("HELO failed: %v", err)
	}

	var validator string = "validator@" + strings.Split(email, "@")[1]
	fmt.Println(validator)

	if err = conn.Mail(validator); err != nil {
		return fmt.Errorf("MAIL FROM failed: %v", err)
	}

	err = conn.Rcpt(email)
	if err != nil {
		return fmt.Errorf("Code := %s:%s, Mail is not Exits", strings.Split(string(err.Error()), " ")[0], strings.Split(string(err.Error()), " ")[1])
	}

	fmt.Printf("Email exists: %s\n", email)
	return nil
}
