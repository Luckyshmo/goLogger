package log //todo should be imported from gitlab repo

import (
	"io"
	"log"
)

//Custom loggers to file
var (
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
	Trace   *log.Logger
)

//Init initialize loggers
func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {
	// file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func exampleUsage() {
	Info.Println("Starting the application...")
	Info.Println("Something noteworthy happened")
	Warning.Println("There is something you should know about")
	Error.Fatal("Something went wrong")
}
