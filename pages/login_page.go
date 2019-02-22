package pages

// LoginPage encapsulates login page information and actions
type LoginPage struct {
	Page Page
}

var (
	campoEmail = "username"
	campoSenha = "password"
	botaoLogin = "btnlogin"
	alertError = "alert-error"
)

// LogIn performs a login with the given email and password
func (s *LoginPage) LogIn(email, senha string) *AccountPage {

	s.Page.InputText(s.Page.FindElementByID(campoEmail), email)
	s.Page.InputText(s.Page.FindElementByID(campoSenha), senha)
	s.Page.FindElementByID(botaoLogin).Click()
	return &AccountPage{Page: s.Page}
}

// GetErrorMsg returns the error message displayed in login page
func (s *LoginPage) GetErrorMsg() string {
	text, _ := s.Page.FindElementByClass(alertError).Text()
	return text
}
