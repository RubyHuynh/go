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
	log.Println(" checking tiki api connection...")
	log.Println(tiki.Access())
	log.Printf(" rs=%v\n", tiki.Get(123))
}
