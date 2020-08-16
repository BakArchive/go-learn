package others

import "fmt"

const (
	// days
	monday    = "星期一"
	tuesday   = "星期二"
	wednesday = "星期三"
	thursday  = "星期四"
	friday    = "星期五"
	saturday  = "星期六"
	sunday    = "星期天"

	// blank
	blank = "|\t\t|\t\t|"
	// border
	border = "|\t---\t|\t---\t|\t---\t|"

	// start & end
	start = 7
	end   = 24
)

func main() {
	// print title
	//days := []string{monday, tuesday, wednesday, thursday, friday, saturday, sunday}
	days := []string{"周一 ~ 周五", "周六、周日"}
	title := "|		|"
	for i := 0; i < len(days); i++ {
		title += "\t" + days[i] + "\t|"
	}
	fmt.Println(title)

	// print line
	fmt.Println(border)

	// print body
	for i, u := start, 0; i < end; {
		if u%2 == 0 {
			fmt.Printf("|\t%02d:00 - %02d:30\t%s\n", i, i, blank)
		} else {
			fmt.Printf("|\t%02d:30 - %02d:00\t%s\n", i, i+1, blank)
			i++
		}
		u++
	}
}
