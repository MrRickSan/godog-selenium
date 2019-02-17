package mark

import (
	"github.com/DATA-DOG/godog"
	"github.com/mrricksan/mark/support"
	"github.com/tebeka/selenium"
)

var driver selenium.WebDriver

func queEuAcesseiAPaginaPrincipal() error {
	return godog.ErrPending
}

func facoLoginComE(arg1, arg2 string) error {
	return godog.ErrPending
}

func souAtenticadoComSucesso() error {
	return godog.ErrPending
}

func euDevoVerASeguinteMensagem(arg1 string) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^que eu acessei a pagina principal$`, queEuAcesseiAPaginaPrincipal)
	s.Step(`^faco login com "([^"]*)" e "([^"]*)"$`, facoLoginComE)
	s.Step(`^sou atenticado com sucesso$`, souAtenticadoComSucesso)
	s.Step(`^eu devo ver a seguinte mensagem "([^"]*)"$`, euDevoVerASeguinteMensagem)

	s.BeforeScenario(func(interface{}) {
		support.WDInit()
	})
}
