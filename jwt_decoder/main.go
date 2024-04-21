package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strings"
	"time"
)

var (
	includeHeaders bool
	headers        = make(map[string]interface{})
	claims         = make(map[string]interface{})
	err            error
)

func main() {
	startTime := time.Now()

	if len(os.Args) < 2 || matchHelp(os.Args) {
		printUsage()
		os.Exit(1)
	}

	token := os.Args[1]
	includeHeaders := false

	if len(os.Args) > 2 && slices.Index(os.Args, "-h") > -1 {
		includeHeaders = true
	}

	jwtSections := strings.Split(token, ".")

	if len(jwtSections) != 3 {
		log.Fatalf("Provided string is not a valid token: %s\n", token)
	}

	var headersIndented string
	if includeHeaders {
		headersIndented = decodeAndIndent(jwtSections[0], "headers")
	}

	payload := decodeAndIndent(jwtSections[1], "payload")

	if includeHeaders {
		fmt.Printf("HEADERS:\n%s\n", headersIndented)
	}
	fmt.Printf("\nPAYLOAD:\n%s\n", payload)

	currTime := time.Now().Unix()

	fmt.Printf("Time to decode JWT: %s\n", time.Since(startTime))

	exp, ok := (claims)["exp"].(float64)
	if ok {
		timeDiff := exp - float64(currTime)
		coloredExpiration := toColorRed(fmt.Sprintf("%f minutes", timeDiff/60))
		fmt.Printf("Token expires in %s", coloredExpiration)
	}

}

func decodeAndIndent(jwtSection, target string) string {
	decoded, err := decodeString(jwtSection)
	if err != nil {
		log.Fatalf("could not decode %s:\n%s", target, err)
	}
	var indented []byte

	switch target {
	case "headers":
		if err := json.Unmarshal(decoded, &headers); err != nil {
			log.Fatalf("could not structure %s:\n%s", target, err)
		}

		indented, err = json.MarshalIndent(headers, "", "  ")
		if err != nil {
			log.Fatalf("could not structure %s:\n%s", target, err)
		}
	case "payload":
		if err := json.Unmarshal(decoded, &claims); err != nil {
			log.Fatalf("could not structure %s:\n%s", target, err)
		}

		indented, err = json.MarshalIndent(claims, "", "  ")
		if err != nil {
			log.Fatalf("could not structure %s:\n%s", target, err)
		}

	}
	return string(indented)
}

func decodeString(s string) ([]byte, error) {
	bytesArr, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return []byte{}, err
	}
	return bytesArr, nil
}

func printUsage() {
	fmt.Printf("Usage: %s <jwt_token_here> [-h]\n", os.Args[0])
	fmt.Println("Decodes the provided JWT token.")
	fmt.Println("Flags:")

	printFlag := func(flag, flag_type, description string) {
		fmt.Printf("\t%s (%s): %s\n", flag, flag_type, description)
	}
	printFlag("-h", "string", "whether to include headers in output")
}

func matchHelp(osArgs []string) bool {
	regex, err := regexp.Compile("(--help)|(--h)|(-h)")
	if err != nil {
		log.Fatalln("Could not compile regex pattern")

	}
	matched := false
	for _, ele := range osArgs {
		if regex.MatchString(ele) {
			matched = true
			break
		}

	}
	return matched
}

func toColorRed(s string) string {
	return fmt.Sprintf("\033[31m%s\033[0m\n", s)
}
