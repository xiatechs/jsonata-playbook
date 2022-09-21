package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"text/template"

	jsonata "github.com/xiatechs/jsonata-go"
)

var endpoint = ":8050"

// SetEndpoint - set the endpoint for the app
func SetEndpoint(input string) {
	endpoint = input
}

type PageVariables struct {
	Input   string
	Jsonata string
	Output  string
}

var globalVariables = PageVariables{}

func mainpage(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("mainpage").Parse(mapage)
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
	err2 := t.Execute(w, globalVariables)
	if err2 != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func validate(r *http.Request, item string) (string, bool) {
	if len(r.Form[item]) != 0 && r.Form[item][0] != "" {
		return r.Form[item][0], true
	}
	return "", false
}

func start(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("mainpage").Parse(mapage)
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}

	_ = r.ParseForm()

	submitName := r.FormValue("submit")
	log.Println(submitName)

	var ok1, ok2 bool

	globalVariables.Input, ok1 = validate(r, "inputdata")

	globalVariables.Jsonata, ok2 = validate(r, "jsonatadata")

	if ok1 && ok2 {
		if submitName == "submitquery" {
			globalVariables.Output = processJsonata(globalVariables.Input, globalVariables.Jsonata)
		}
	}

	if ok2 {
		if submitName == "escapequery" {
			globalVariables.Output = jsonEscape(globalVariables.Jsonata)
		}
	}

	err = t.Execute(w, globalVariables)
	if err != nil { // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

func Start() {
	http.HandleFunc("/", mainpage)
	http.HandleFunc("/ui/process", start)
	launch()
}

func launch() {
	fmt.Println("Booting up server... - ", endpoint)
	err := http.ListenAndServe(endpoint, nil) // setting listening port
	if err != nil {
		fmt.Println(err)
	}
}

func generateCSS() string {
	return `
	body {
		background: #e5e5e5;
	}
	
	#title {
		color: #14213d;
		font-size: 30px;
		text-align: center;
		text-decoration: underline;
		text-decoration-color: #fca311;

	}

	#text-boxes {
		display: flex;
		justify-content: space-around;
	}
	
	textarea {
		height: 500px;
		width: 600px;
		border: solid 3px #14213d;
	}
	
	#btns {
		display: flex;
		justify-content: space-around;
		margin-top: 50px;
		
	}

	button {
		padding: 20px;
		background: #14213d;
		color: #fca311;
		font-weight: bold;
		border: solid 4px #fca311;
		border-radius: 15px;
	}
	`
}

var mapage = fmt.Sprintf(`<title>GO QL - Front-end SQL</title>
</head>

<body>
    <style>
     %s   
    </style>
    <p id="title"><b>Go Jsonata Frontend</b></p>
    <form action="/ui/process" method="POST">
		<div id="text-boxes">
				<textarea name="inputdata" >{{.Input}}</textarea>
				<textarea name="jsonatadata" >{{.Jsonata}}</textarea>
				<textarea name="outputdata" >{{.Output}}</textarea>
		</div>

		<div id="btns">
				<button type="submit" value="submitquery" name="submit">Submit Jsonata Query</button>
				<button type="submit" value="escapequery" name="submit">Escape Jsonata Query</button>
		</div>
	</form>
	 <br/>
    <br>
</body>`, generateCSS())

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

func jsonEscape(i string) string {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	// Trim the beginning and trailing " character
	return string(b[1 : len(b)-1])
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
