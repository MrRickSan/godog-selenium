package pages

// HomePage encapsulates home page information and actions
type HomePage struct {
	Page Page
}

var (
	signInToggle  = "dropdown-toggle"
	logInLinkText = "https://me.hack.me/login"
)

func (s *HomePage) GoToLoginPage() *LoginPage {
	s.Page.MouseHoverToElement(signInToggle)
	s.Page.FindElementByLinkText(logInLinkText).Click()
	return &LoginPage{Page: s.Page}
}
