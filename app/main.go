package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"runtime"

	jsonata "github.com/blues/jsonata-go"
)

//go:embed public
var embeddedFiles embed.FS

func main() {
	fsys, err := fs.Sub(embeddedFiles, "public")
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.FS(fsys)))

	http.Handle("/jsonata", http.HandlerFunc(jsonataRequest))

	log.Println("booting up server...")

	if len(os.Args) > 1 {
		log.Println(os.Args[1])
		if os.Args[1] == "web" {
			openBrowser("http://127.0.0.1:8050")
		}
	}

	http.ListenAndServe(":8050", nil)
}

//open up browser/tab dependent on your OS.
func openBrowser(url string) (err error) {
	switch runtime.GOOS { //open browser/tab dependent on what OS you are on
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform????")
	}

	return
}

type request struct {
	Input   string
	Jsonata string
	Output  string
}

func jsonataRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // for CORS
	data := request{}

	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			json.NewEncoder(w).Encode(data)
			return
		}

		_ = r.Body.Close()

		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			json.NewEncoder(w).Encode(data)
			return
		}

		response := processJsonata(data.Input, data.Jsonata)

		data.Output = response

		json.NewEncoder(w).Encode(data)

	default:
		fmt.Fprintf(w, "Request type other than POST not supported")
		json.NewEncoder(w).Encode(data)
	}
}

func processJsonata(input, jsonataString string) (output string) {
	defer func() {
		if r := recover(); r != nil {
			output = fmt.Sprintf("jsonata error: %v", r)
		}
	}()

	var dataToInterface interface{}

	err := json.Unmarshal([]byte(input), &dataToInterface)
	if err != nil {
		output = fmt.Sprintf("input json error: %v", err)
		return
	}

	jsnt := replaceQuotesAndCommentsInPaths(jsonataString)

	e := jsonata.MustCompile(jsnt)

	res, err := e.Eval(dataToInterface)
	if err != nil {
		return "jsonata error: " + err.Error()
	}

	str, _ := json.MarshalIndent(res, "", " ")

	return string(str)
}

/*
	enables:
	- comments in jsonata code
	- fields with any character in their name
*/

var (
	reQuotedPath      = regexp.MustCompile(`([A-Za-z\$\\*\` + "`" + `])\.[\"']([\s\S]+?)[\"']`)
	reQuotedPathStart = regexp.MustCompile(`^[\"']([ .0-9A-Za-z]+?)[\"']\.([A-Za-z\$\*\"\'])`)
	commentsPath      = regexp.MustCompile(`/\*([\s\S]*?)\*/`)
)

func replaceQuotesAndCommentsInPaths(s string) string {
	if reQuotedPathStart.MatchString(s) {
		s = reQuotedPathStart.ReplaceAllString(s, "`$1`.$2")
	}

	for reQuotedPath.MatchString(s) {
		s = reQuotedPath.ReplaceAllString(s, "$1.`$2`")
	}

	for commentsPath.MatchString(s) {
		s = commentsPath.ReplaceAllString(s, "")
	}

	return s
}
