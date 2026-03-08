package aggregator

const (
	SUM = "sum"
	MIN = "min"
 MAX = "max"
)

func Aggregator(operation string, data []int) (result int){

	switch(operation) {
		case SUM:
			return sum(data)			  
		case MIN:
			return min(data)
		case MAX:
			return max(data)
		default :
			return 0
	}
}

	func sum(data []int) (sum int) {
		  for _, v := range data {
				sum = sum + v
			}
			return
	}


	func min(data []int) (min int) {
		  min = data[0]
		  for _,v :=range data {
				if v < min {
					min = v
				}
			}
			return
	}


	func max(data []int) (max int) {
		  for _,v := range data {
				  if  v > max {
					max = v
				}
			}
			return
	}