package mark

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/DATA-DOG/godog/gherkin"

	"github.com/DATA-DOG/godog"
	"github.com/mrricksan/mark/pages"
	"github.com/mrricksan/mark/support"
)

var page pages.Page

func queEuAcesseiAPaginaPrincipal() error {
	page.Visit("https://hack.me")
	return nil
}

func facoLoginComE(email, senha string) (err error) {
	homePage := pages.HomePage{Page: page}
	homePage.GoToLoginPage().LogIn(email, senha)
	return nil
}

func souAtenticadoComSucesso() (err error) {
	authenticated := pages.AccountPage{Page: page}
	if !authenticated.IsAuthenticated() {
		return fmt.Errorf("Erro ao validar usuário autenticado")
	}
	return nil
}

func euDevoVerASeguinteMensagem(msg string) (err error) {
	loginPage := pages.LoginPage{Page: page}
	saida := loginPage.GetErrorMsg()
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
		page.Driver = support.WDInit()
	})

	s.AfterScenario(func(i interface{}, e error) {

		// Pega o nome do scenario para usar no nome do arquivo da screenshot
		// e subistitui espaços em branco e caracteres especiais por "-"
		sc := i.(*gherkin.Scenario)
		rgex := regexp.MustCompile("[^0-9a-zA-Z]+")
		fileName := strings.ToLower(rgex.ReplaceAllString(sc.Name, "-"))

		shot, _ := page.Driver.Screenshot()
		support.SaveImage(shot, fileName)

		err := page.Driver.Quit()
		if err != nil {
			fmt.Println(err)
		}
	})
}
