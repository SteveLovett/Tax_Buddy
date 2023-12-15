package TaxPrepAssistant

func main() {

}

type personOnReturn interface {
	personalInformation
}

type personalInformation struct {
	FirstName string
	LastName  string
	SSN       string
	DOB       string
	Email     string
	Phone     string
	Address   string
	City      string
	State     string
	Zip       string
}
