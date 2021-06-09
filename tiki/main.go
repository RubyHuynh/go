package main

import(
	"hthngoc/logging"
//	"os"
	"hthngoc/tikitiki"
)

func init() {
	/*f, err := os.OpenFile("./full.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("Starting service...")*/
}
func main() {
	logger := log.GetInstance()
	logger.Println("Starting.....")
	tiki.Get(123)
}
