package coffeebean

import "fmt"

func main() {
	initLog(false)
	err := runCmd()
	if err != nil {
		fmt.Println(err.Error())
	}
}
