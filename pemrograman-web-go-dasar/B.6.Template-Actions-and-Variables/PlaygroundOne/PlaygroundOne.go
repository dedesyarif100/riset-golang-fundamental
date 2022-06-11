package main

type Info struct {
	Affiliation string
	Address		string
}

type Person struct {
	Name	string
	Gender	string
	Hobbies	[]string
	Info	Info
	IsTrue	bool
	Abjad	string
}

func (t *Info) GetAffiliationDetailInfo() string {
	return "HAVE 31 DIVISIONS"
}

func main() {

}