package jobs

import (
	"log"

	"github.com/robfig/cron/v3"
)

// InitialJobs is function
// khong can su dung go routine khi goi ham nay
// vi ban than c.Start() da tao ra 1 go rotine ben trong

func InitialJobs() {

	c := cron.New()

	// add job with check error
	jobID, err := c.AddFunc("@every 10s", jobEvery10Seconds)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Job Every 10 Seconds Ok => ID: %v", jobID)
	}

	// add job normal
	c.AddFunc("@every 1m3s", jobEveryMinute)

	// start job
	c.Start()

	// kill job
	// c.Remove(jobID)
}

// Example of job definition:
// .---------------- minute (0 - 59)
// |  .------------- hour (0 - 23)
// |  |  .---------- day of month (1 - 31)
// |  |  |  .------- month (1 - 12) OR jan,feb,mar,apr ...
// |  |  |  |  .---- day of week (0 - 6) (Sunday=0 or 7) OR sun,mon,tue,wed,thu,fri,sat
// |  |  |  |  |
// *  *  *  *  * user-name  command to be executed

// document link: https://pkg.go.dev/github.com/robfig/cron
// Entry                  | Description                                | Equivalent To
// -----                  | -----------                                | -------------
// @yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *
// @monthly               | Run once a month, midnight, first of month | 0 0 0 1 * *
// @weekly                | Run once a week, midnight between Sat/Sun  | 0 0 0 * * 0
// @daily (or @midnight)  | Run once a day, midnight                   | 0 0 0 * * *
// @hourly                | Run once an hour, beginning of hour        | 0 0 * * * *

// vi du 1: chay job tao bao cao moi quy
// syntax	: 02 03 1 1,4,7,10 *
// desc		: At 03:02 on day-of-month 1 in January, April, July, and October
// job nay cho phep chung ta chay vao ngay dau tien cua cac thang 1, 4, 7, 10 luc 03:02
// vi khong co mau cho ngay cuoi cua thang, nen chung ta su dung ngay dau tien
// cua thang ke tiep de chay job

// vi du 2:
// syntax	: 01 09-17 * * *
// desc		: At minute 1 past every hour from 9 through 17
// job chạy trong vòng một phút, mỗi giờ một lần, từ 9:01 a.m đến 5:01 p.m

// vi du 3:
// syntax	: */5 08-18/2 * * *
// desc		: At every 5th minute past every 2nd hour from 8 through 18
// Đôi khi, có những job cần được thực hiện sau mỗi 2, 3 hay 4 giờ. Khi đó,
// ta có thể lấy thương số của giờ và khoảng thời gian mong muốn.
// Chẳng hạn như */3, tương đương với job sau mỗi ba giờ. Hay 6-18/3 để chạy mỗi ba tiếng,
// từ 6 a.m đến 6 p.m. Các khoảng thời gian khác cũng có thể được chia tương tự.
// Lấy ví dụ, biểu thức */15 ở vị trí phút có nghĩa là “chạy job sau mỗi 15 phút”.

// vi du 4:
// syntax	: 00 15 * * Thu
// desc		: At 15:00 on Thursday
// chạy một job nào đó vào 3 giờ chiều, mỗi thứ Ba
