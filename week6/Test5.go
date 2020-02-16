
import "fmt"

func main() {
	r := []bool{true, false, true, true, false, true}
	for _, v := range r {
		fmt.Println("%t\n", v)
	}
}
