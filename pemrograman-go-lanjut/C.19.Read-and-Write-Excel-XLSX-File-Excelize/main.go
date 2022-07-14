package main

import (
    "fmt"
    "github.com/360EntSecGroup-Skylar/excelize"
    "log"
)

type M map[string]interface{}

var data = []M{
    M{"Name": "DEDE", "Gender": "MALE", "Age": 22},
    M{"Name": "HENDRO", "Gender": "FEMALE", "Age": 23},
    M{"Name": "RIAND", "Gender": "MALE", "Age": 23},
    M{"Name": "TRISULI", "Gender": "MALE", "Age": 23},
    M{"Name": "MIRDAS", "Gender": "MALE", "Age": 23},
    M{"Name": "ALI", "Gender": "MALE", "Age": 23},
    M{"Name": "MA'ARIF", "Gender": "MALE", "Age": 23},
    M{"Name": "TAMA", "Gender": "MALE", "Age": 23},
}

func main() {
    // magic here
	// C.19.1. Membuat File Excel .xlsx
		// xlsx := excelize.NewFile()

		// sheet1Name := "Sheet One"
		// xlsx.SetSheetName(xlsx.GetSheetName(1), sheet1Name)
		// xlsx.SetCellValue(sheet1Name, "A1", "Name")
		// xlsx.SetCellValue(sheet1Name, "B1", "Gender")
		// xlsx.SetCellValue(sheet1Name, "C1", "Age")

		// err := xlsx.AutoFilter(sheet1Name, "A1", "C1", "")
		// if err != nil {
		// 	log.Fatal("ERROR", err.Error())
		// }

		// for i, each := range data {
		// 	xlsx.SetCellValue(sheet1Name, fmt.Sprintf("A%d", i+2), each["Name"])
		// 	xlsx.SetCellValue(sheet1Name, fmt.Sprintf("B%d", i+2), each["Gender"])
		// 	xlsx.SetCellValue(sheet1Name, fmt.Sprintf("C%d", i+2), each["Age"])
		// }


	// C.19.2. Pembuatan Sheet, Merge Cell, dan Cell 
		// sheet2Name := "Sheet two"
		// sheetIndex := xlsx.NewSheet(sheet2Name)
		// xlsx.SetActiveSheet(sheetIndex)

		// xlsx.SetCellValue(sheet2Name, "A1", "Hello")
		// xlsx.MergeCell(sheet2Name, "A1", "B1")

		// style, err := xlsx.NewStyle(`{
		// 	"font": {
		// 		"bold": true,
		// 		"size": 36
		// 	},
		// 	"fill": {
		// 		"type": "pattern",
		// 		"color": ["#E0EBF5"],
		// 		"pattern": 1
		// 	}
		// }`)
		// if err != nil {
		// 	log.Fatal("ERROR", err.Error())
		// }
		// xlsx.SetCellStyle(sheet2Name, "A1", "A1", style)

		// err = xlsx.SaveAs("file/file2.xlsx")
		// if err != nil {
		// 	fmt.Println(err)
		// }

	// C.19.3. Membaca File Excel .xlsx
		xlsx, err := excelize.OpenFile("file/file1.xlsx")
		if err != nil {
			log.Fatal("ERROR", err.Error())
		}
		
		sheet1Name := "Sheet One"
		
		rows := make([]M, 0)
		for i := 2; i < 5; i++ {
			row := M{
				"Name":   xlsx.GetCellValue(sheet1Name, fmt.Sprintf("A%d", i)),
				"Gender": xlsx.GetCellValue(sheet1Name, fmt.Sprintf("B%d", i)),
				"Age":    xlsx.GetCellValue(sheet1Name, fmt.Sprintf("C%d", i)),
			}
			rows = append(rows, row)
		}
		
		fmt.Printf("%v \n", rows)
}