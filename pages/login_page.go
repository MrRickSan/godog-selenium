package pages

// LoginPage encapsulates login page information and actions
type LoginPage struct {
	Page Page
}

var (
	campoEmail = "username"
	campoSenha = "password"
	botaoLogin = "btnlogin"
)

func (s *LoginPage) logIn(email, senha string) {

	s.Page.InputText(s.Page.FindElementByID(campoEmail), email)
	s.Page.InputText(s.Page.FindElementByID(campoSenha), senha)
	s.Page.FindElementByID(botaoLogin).Click()
}