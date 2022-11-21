package models

import "time"

type Vendor struct {
	Name            string
	Dob             time.Time
	Gender          byte
	GovtID          GovernmentID
	Photo           string // link to a blob storage
	ContactsDetails Contacts
	TaxDetails      Tax
}
type GovernmentID struct {
	ID   string
	Type string // Aadhar or voter ID etc
}
type Contacts struct {
	EmailDetails Emails
	PhoneDetails Phones
	Website      string
}
type Emails struct {
	PrimaryEmail   string
	SecondaryEmail string
}
type Phones struct {
	PrimaryPhone   string
	SecondaryPhone string
	OfficePhone    string
}
type Tax struct {
	ID string // PAN etc
}

func (v *Vendor) ValidateVendorData() error {
	return nil // TODO
}