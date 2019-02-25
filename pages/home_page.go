package pages

// HomePage encapsulates home page information and actions
type HomePage struct {
	Page Page
}

var (
	signInToggle = "//a[@class='dropdown-toggle']"
	logInLink    = "a[href='https://me.hack.me/login']"
)

// GoToLoginPage follow the signin option in the home page
func (s *HomePage) GoToLoginPage() *LoginPage {
	element := s.Page.FindElementByXPATH(signInToggle)
	element.Click()
	s.Page.FindElementByCSS(logInLink).Click()
	return &LoginPage{Page: s.Page}
}
