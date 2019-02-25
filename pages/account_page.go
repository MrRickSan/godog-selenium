package pages

// AccountPage encapsulates account page information and actions
type AccountPage struct {
	Page Page
}

var (
	myAccount = "//a[text()='My Account']"
)

// IsAuthenticated verifies if the user is logged in or not
func (s *AccountPage) IsAuthenticated() bool {
	text, _ := s.Page.FindElementByXPATH(myAccount).Text()

	if text != "My Account" {
		return false
	}
	return true
}
