package main

import (
    // "os"
    "log"
    "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func main() {
    pdfg, err := wkhtmltopdf.NewPDFGenerator()
    if err != nil {
        log.Fatal(err)
    }

	// C.21.1. Konversi File HTML ke PDF
		// f, err := os.Open("file/input.html")
		// if f != nil {
		// 	defer f.Close()
		// }
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// pdfg.AddPage(wkhtmltopdf.NewPageReader(f))

		// pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
		// pdfg.Dpi.Set(300)
		
		// err = pdfg.Create()
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// err = pdfg.WriteFile("file/output.pdf")
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// log.Println("Done")


	// C.21.2. Konversi HTML dari URL Menjadi PDF
		pdfg.Orientation.Set(wkhtmltopdf.OrientationLandscape)

		page := wkhtmltopdf.NewPage("https://vuejs.org/")
		page.FooterRight.Set("[page]")
		page.FooterFontSize.Set(10)
		pdfg.AddPage(page)

		err = pdfg.Create()
		if err != nil {
			log.Fatal(err)
		}

		err = pdfg.WriteFile("file/output.pdf")
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Done")
}