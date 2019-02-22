package pages

// AccountPage encapsulates account page information and actions
type AccountPage struct {
	Page Page
}

var (
	myAccount = "//a[text()='My Account']"
)

func (s *LoginPage) isAuthenticated(email, senha string) bool {
	text, _ := s.Page.FindElementByXPATH(myAccount).Text()

	if text != "My Account" {
		return false
	}
	return true
}
