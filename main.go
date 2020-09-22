package playlist

import (
	"fmt"
	"playlist/server"
)

func main() {
	s := server.New()
	s.Start()
	fmt.Println("did the thing")
}
