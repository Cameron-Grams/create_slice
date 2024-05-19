package create_slice

import(
	"fmt"
	"bufio"
        "os"
	"strconv"
	"strings"
	"log"
)


func EnterDigits() []float64 {

	var numbers []float64
        entering_nums := "Yes"
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(" Enter the numbers for the slice.  Enter 'No' to stop.")

	out:
	for entering_nums != "No" {
            entry, err := reader.ReadString('\n')
	    if err != nil {
		    log.Fatal(err)
	    }

	    if entry == "No\n"{
		   break out 
	    }

	    number_string := strings.TrimSpace(entry)
	    number, err := strconv.ParseFloat(number_string, 64) 
            if err != nil {
		    log.Fatal(err)
	    }

	    numbers = append(numbers, number)
	}
        return numbers
}
