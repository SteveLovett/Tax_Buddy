package main

import (
	"fmt"
	"os"

	huh "github.com/charmbracelet/huh"
)

func (person personalInformation) PersonalInfo() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("First name:").
				Value(&person.FirstName).
				Description("Let's start with the basics, enter the person's personal information."),
			huh.NewInput().
				Title("Last name:").
				Value(&person.LastName),
			huh.NewInput().
				Placeholder("555-55-5555").
				Title("SSN:").
				Value(&person.SSN),
			huh.NewInput().
				Placeholder("01/01/2000").
				Title("DoB:").
				Value(&person.DOB),
			huh.NewInput().
				Placeholder("example@email.com").
				Title("Email:").
				Value(&person.Email),
			huh.NewInput().
				Placeholder("(555)555-5555").
				Title("Phone:").
				Value(&person.Phone),
			huh.NewConfirm().
				Title("Foreign address?").
				Value(&person.Foreign).
				Affirmative("Yes").
				Negative("No"),
			huh.NewConfirm().
				Title("2nd address line needed?").
				Value(&person.Add2Bool).
				Affirmative("Yes").
				Negative("No"),
		),
	)

	err := form.Run()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	if person.Foreign {
		if person.Add2Bool {
			form = huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("Address:").
						Value(&person.Address1),
					huh.NewInput().
						Title("Address line 2:").
						Value(&person.Address2),
					huh.NewInput().
						Title("City:").
						Value(&person.Address2),
					huh.NewInput().
						Title("State:").
						Value(&person.State),
					huh.NewInput().
						Title("Zip:").
						Value(&person.Zip),
					huh.NewInput().
						Title("Country:").
						Value(&person.Country),
				),
			)
			err := form.Run()
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}
		} else {
			form = huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("Address:").
						Value(&person.Address1),
					huh.NewInput().
						Title("City:").
						Value(&person.Address2),
					huh.NewInput().
						Title("State:").
						Value(&person.State),
					huh.NewInput().
						Title("Zip:").
						Value(&person.Zip),
					huh.NewInput().
						Title("Country:").
						Value(&person.Country),
				),
			)
			err := form.Run()
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}
		}
	} else {
		if person.Add2Bool {
			form = huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("Address:").
						Value(&person.Address1),
					huh.NewInput().
						Title("Address line 2:").
						Value(&person.Address2),
					huh.NewInput().
						Title("City:").
						Value(&person.Address2),
					huh.NewInput().
						Title("State:").
						Value(&person.State),
					huh.NewInput().
						Title("Zip:").
						Value(&person.Zip),
				),
			)
			err := form.Run()
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
	}
		} else {
			form = huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("Address:").
						Value(&person.Address1),
					huh.NewInput().
						Title("City:").
						Value(&person.Address2),
					huh.NewInput().
						Title("State:").
						Value(&person.State),
					huh.NewInput().
						Title("Zip:").
						Value(&person.Zip),
				),
			)
			err := form.Run()
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}
		}
	}
	

	// p := ""
	// fmt.Println("Let's start with the basics, enter the person's personal information.")
	// p = "First name:"
	// textEntry(p, &person.FirstName)
	// p = "Last name:"
	// textEntry(p, &person.LastName)
	// p = "SSN:"
	// textEntry(p, &person.SSN)
	// p = "DOB:"
	// textEntry(p, &person.DOB)
	// p = "Email:"
	// textEntry(p, &person.Email)
	// p = "Phone:"
	// textEntry(p, &person.Phone)
	// p = "Foreign address? Enter 'true' or 'false'."
	// questionTF(p, &person.Foreign)
	// if person.Foreign {
	// 	p = "Country:"
	// 	textEntry(p, &person.Country)
	// }
	// p = "Address:"
	// textEntry(p, &person.SSN)
	// p = "2nd address line needed? Enter 'true' or 'false'."
	// questionTF(p, &person.Add2Bool)
	// if person.Add2Bool {
	// 	p = "Address line 2:"
	// 	textEntry(p, &person.Address2)
	// }
	// p = "City:"
	// textEntry(p, &person.City)
	// p = "State:"
	// textEntry(p, &person.State)
	// p = "Zip:"
	// textEntry(p, &person.Zip)
}
