package pages

import (
	"fmt"
	"log"

	"github.com/tebeka/selenium"
)

// Page is an abstraction for selenium webdriver
type Page struct {
	Driver selenium.WebDriver
}

func (s *Page) driver() selenium.WebDriver {
	return s.Driver
}

// Visit opens the given url
func (s *Page) Visit(url string) {
	err := s.Driver.Get(url)
	if err != nil {
		log.Fatal(err)
	}
}

// InputText clears and send text to the given element
func (s *Page) InputText(element selenium.WebElement, text string) {
	var err error
	err = element.Clear()
	if err != nil {
		fmt.Println(err)
	}
	err = element.SendKeys(text)
	if err != nil {
		fmt.Println(err)
	}
}

// FindElementByID finds the given element by id
func (s *Page) FindElementByID(locator string) selenium.WebElement {
	element, err := s.Driver.FindElement(selenium.ByID, locator)
	if err != nil {
		fmt.Println(err)
	}
	return element
}

// FindElementByXPATH finds the given element by XPATH
func (s *Page) FindElementByXPATH(locator string) selenium.WebElement {
	element, err := s.Driver.FindElement(selenium.ByXPATH, locator)
	if err != nil {
		fmt.Println(err)
	}
	return element
}

// FindElementByLinkText finds the given element by LinkText
func (s *Page) FindElementByLinkText(locator string) selenium.WebElement {
	element, err := s.Driver.FindElement(selenium.ByLinkText, locator)
	if err != nil {
		fmt.Println(err)
	}
	return element
}

// FindElementByPartialLink finds the given element by PartialLinkText
func (s *Page) FindElementByPartialLink(locator string) selenium.WebElement {
	element, err := s.Driver.FindElement(selenium.ByPartialLinkText, locator)
	if err != nil {
		fmt.Println(err)
	}
	return element
}

// FindElementByName finds the given element by Name
func (s *Page) FindElementByName(locator string) selenium.WebElement {
	element, err := s.Driver.FindElement(selenium.ByName, locator)
	if err != nil {
		fmt.Println(err)
	}
	return element
}

// FindElementByTag finds the given element by Tag
func (s *Page) FindElementByTag(locator string) selenium.WebElement {
	element, err := s.Driver.FindElement(selenium.ByTagName, locator)
	if err != nil {
		fmt.Println(err)
	}
	return element
}

// FindElementByClass finds the given element by Class
func (s *Page) FindElementByClass(locator string) selenium.WebElement {
	element, err := s.Driver.FindElement(selenium.ByClassName, locator)
	if err != nil {
		fmt.Println(err)
	}
	return element
}

// FindElementByCSS finds the given element by CSS
func (s *Page) FindElementByCSS(locator string) selenium.WebElement {
	element, err := s.Driver.FindElement(selenium.ByCSSSelector, locator)
	if err != nil {
		fmt.Println(err)
	}
	return element
}

// FindElementsByXPATH finds the given elements by XPATH
func (s *Page) FindElementsByXPATH(locator string) []selenium.WebElement {
	element, err := s.Driver.FindElements(selenium.ByXPATH, locator)
	if err != nil {
		fmt.Println(err)
	}
	return element
}

// FindElementsByLinkText finds the given elements by LinkText
func (s *Page) FindElementsByLinkText(locator string) []selenium.WebElement {
	element, err := s.Driver.FindElements(selenium.ByLinkText, locator)
	if err != nil {
		fmt.Println(err)
	}
	return element
}

// FindElementsByPartialLink finds the given elements by PartialLinkText
func (s *Page) FindElementsByPartialLink(locator string) []selenium.WebElement {
	element, err := s.Driver.FindElements(selenium.ByPartialLinkText, locator)
	if err != nil {
		fmt.Println(err)
	}
	return element
}

// FindElementsByName finds the given elements by Name
func (s *Page) FindElementsByName(locator string) []selenium.WebElement {
	element, err := s.Driver.FindElements(selenium.ByName, locator)
	if err != nil {
		fmt.Println(err)
	}
	return element
}

// FindElementsByTag finds the given elements by Tag
func (s *Page) FindElementsByTag(locator string) []selenium.WebElement {
	element, err := s.Driver.FindElements(selenium.ByTagName, locator)
	if err != nil {
		fmt.Println(err)
	}
	return element
}

// FindElementsByClass finds the given elements by Class
func (s *Page) FindElementsByClass(locator string) []selenium.WebElement {
	element, err := s.Driver.FindElements(selenium.ByClassName, locator)
	if err != nil {
		fmt.Println(err)
	}
	return element
}

// FindElementsByCSS finds the given elements by CSS
func (s *Page) FindElementsByCSS(locator string) []selenium.WebElement {
	element, err := s.Driver.FindElements(selenium.ByCSSSelector, locator)
	if err != nil {
		fmt.Println(err)
	}
	return element
}
