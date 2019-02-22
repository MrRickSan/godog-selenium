package mark

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/DATA-DOG/godog/gherkin"

	"github.com/DATA-DOG/godog"
	"github.com/mrricksan/mark/pages"
	"github.com/mrricksan/mark/support"
	"github.com/tebeka/selenium"
)

var page pages.Page

func queEuAcesseiAPaginaPrincipal() error {
	page.Visit("https://me.hack.me/login")
	// Driver.Get("https://me.hack.me/login")
	return nil
}

func facoLoginComE(email, senha string) (err error) {
	login := pages.HomePage{Page: page}
	login.
	return nil
}

func souAtenticadoComSucesso() (err error) {
	myAccount, err := Driver.FindElement(selenium.ByXPATH, "//a[text()='My Account']")
	if err != nil {
		return
	}

	saida, _ := myAccount.Text()

	if saida != "My Account" {
		return fmt.Errorf("Erro ao validar usuário autenticado")
	}
	return nil
}

func euDevoVerASeguinteMensagem(msg string) (err error) {
	divAlerta, err := Driver.FindElement(selenium.ByClassName, "alert-error")
	if err != nil {
		return
	}

	saida, _ := divAlerta.Text()

	if !strings.Contains(saida, msg) {
		return fmt.Errorf("Esperava: %v Obtido: %v", msg, saida)
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^que eu acessei a pagina principal$`, queEuAcesseiAPaginaPrincipal)
	s.Step(`^faco login com "([^"]*)" e "([^"]*)"$`, facoLoginComE)
	s.Step(`^sou atenticado com sucesso$`, souAtenticadoComSucesso)
	s.Step(`^eu devo ver a seguinte mensagem "([^"]*)"$`, euDevoVerASeguinteMensagem)

	s.BeforeScenario(func(interface{}) {
		Driver = support.WDInit()
	})

	s.AfterScenario(func(i interface{}, e error) {

		// Pega o nome do scenario para usar no nome do arquivo da screenshot
		// e subistitui espaços em branco e caracteres especiais por "-"
		sc := i.(*gherkin.Scenario)
		rgex := regexp.MustCompile("[^0-9a-zA-Z]+")
		fileName := strings.ToLower(rgex.ReplaceAllString(sc.Name, "-"))

		shot, _ := Driver.Screenshot()
		support.SaveImage(shot, fileName)

		Driver.Quit()
	})
}
