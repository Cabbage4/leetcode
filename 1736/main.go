package main

func main() {
}

func maximumTime(time string) string {
	r := []byte(time)
	if time[0] == '?' {
		if time[1] == '?' {
			r[0] = '2'
		} else if time[1] > '3' {
			r[0] = '1'
		} else {
			r[0] = '2'
		}
	}

	if time[1] == '?' {
		if time[0] == '?' {
			r[1] = '3'
		} else if time[0] == '2' {
			r[1] = '3'
		} else {
			r[1] = '9'
		}
	}

	if time[3] == '?' {
		r[3] = '5'
	}

	if time[4] == '?' {
		r[4] = '9'
	}

	return string(r)
}
