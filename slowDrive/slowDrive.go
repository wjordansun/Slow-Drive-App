package main
​
import (
	"fmt"
	"io/ioutil"
    "log"
    "strconv"
    "os"
//    "math"
)
​
var (
	str string
	filename string
	textFiller = "Text Filler for file"
	x float64 = 2
	extraText string
	num = 1000
)
​
func deleteFile() {
​
	errr := os.Remove(filename)
   	if errr != nil {
    log.Fatal(errr)
  	}
}
​
func main() {
	for i := 1; i <= num; i++ {
​
		str = strconv.Itoa(i)
		filename = "file" + str + ".txt"
​
		// x = math.Pow(4, x)
		// extraText = strconv.Itoa(int(x))
		// textFiller += extraText
​
		for j := 0; j < 10000; j++ {
			extraText = "more text to fill file"
			textFiller += extraText 
		}
​
		err := ioutil.WriteFile(filename, []byte(textFiller), 0666)
    	if err != nil {
        log.Fatal(err)
    	}
​
    	fmt.Printf("%v\n", i)
​
    	deleteFile()
        
    }
​
	
}