package main

import (
	"fmt"
	"time"
)

func main() {
	// A.40.1. Penggunaan time.Time
	fmt.Println("# - A.40.1. Penggunaan time.Time");
		var time1 = time.Now()
		fmt.Printf("time1 %v\n", time1)
		// time1 2015-09-01 17:59:31.73600891 +0700 WIB
		var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
		fmt.Printf("time2 %v\n", time2)
		// time2 2011-12-24 10:20:00 +0000 UTC
		fmt.Println();

	// A.40.2. Method Milik time.Time
	fmt.Println("# - A.40.2. Method Milik time.Time");
		var now = time.Now()
		fmt.Println("year		: ", now.Year(), "month:", now.Month())
		fmt.Println("YEAR DAY 	: ", now.YearDay());
		fmt.Println("WEEK DAY 	: ", now.Weekday());
		// fmt.Println("ISO WEEK : ", now.ISOWeek());
		fmt.Println("DAY 		: ", now.Day());
		fmt.Println("HOUR 		: ", now.Hour());
		fmt.Println("MINUTE 	: ", now.Minute());
		fmt.Println("SECOND 	: ", now.Second());
		fmt.Println("Nanosecond : ", now.Nanosecond());
		fmt.Println("LOCAL 		: ", now.Local());
		fmt.Println("LOCATION 	: ", now.Location());
		// fmt.Println("ZONE 		: ", now.Zone());
		fmt.Println("IS ZERO 	: ", now.IsZero());
		fmt.Println("UTC 		: ", now.UTC());
		fmt.Println("Unix 		: ", now.Unix());
		fmt.Println("UNIXNANO 	: ", now.UnixNano());
		fmt.Println("STRING 	: ", now.String());
		fmt.Println();

	// A.40.3. Parsing dari string ke time.Time
	fmt.Println("# - A.40.3. Parsing dari string ke time.Time");
		var layoutFormat, value string
		var date time.Time

		layoutFormat = "2006-01-02 15:04:05"
		value = "2015-09-02 08:04:00"
		date, _ = time.Parse(layoutFormat, value)
		fmt.Println(value, "\t->", date.String())
		// 2015-09-02 08:04:00 +0000 UTC

		layoutFormat = "02/01/2006 MST"
		value = "02/09/2015 WIB"
		date, _ = time.Parse(layoutFormat, value)
		fmt.Println(value, "\t\t->", date.String())
		fmt.Println();

	// A.40.4. Predefined Layout Format Untuk Keperluan Parsing Time
	fmt.Println("# - A.40.4. Predefined Layout Format Untuk Keperluan Parsing Time");
		fmt.Println();
	
	fmt.Println("# - A.40.5. Format dari time.Time ke string");
		var dates, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
		fmt.Println(dates.String())
		
		var d, _ = time.Parse(time.RFC822, "02 Jan 22 08:00 WIB");
		fmt.Println(d)

		var dateS1 = dates.Format("Monday 02, January 2006 15:04 MST");
		fmt.Println("dateS1 :", dateS1)

		var dateS2 = date.Format(time.RFC3339)
		fmt.Println("dateS2 :", dateS2)
		fmt.Println();

	// A.40.6. Handle Error Parsing time.Time
	fmt.Println("# - A.40.6. Handle Error Parsing time.Time");
		var exampleDateSix, err = time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")
		if err != nil {
			fmt.Println("error", err.Error())
			return
		}
		fmt.Println(exampleDateSix)
}