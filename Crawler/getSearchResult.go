package Crawler

import "github.com/tebeka/selenium"

func getSearchResult(driver selenium.WebDriver, searchResultTR string, notFoundTR string) (bool, error) {
	srchReturn, err := searchReturn(driver, searchResultTR)
	if err != nil {
		return false, err
	}

	if srchReturn {
		resultElem, err := driver.FindElements(selenium.ByXPATH, searchResultTR)
		if err != nil {
			return false, err
		}

		foundBooks, err := bookWasFound(resultElem, notFoundTR)
		if err != nil {
			return false, err
		}

		if foundBooks {
			return true, nil
		}
	}

	return false, nil
}

func searchReturn(driver selenium.WebDriver, searchResultTR string) (bool, error) {
	resultElem, err := driver.FindElements(selenium.ByXPATH, searchResultTR)
	if err != nil {
		return false, err
	}

	if len(resultElem) != 0 {
		return true, nil
	}

	return false, nil
}

func bookWasFound(resultElem []selenium.WebElement, notFoundTR string) (bool, error) {
	resultText, err := resultElem[0].Text()
	if err != nil {
		return false, err
	}

	if len(resultText) >= 16 {
		if resultText[0:16] != notFoundTR {
			return true, nil
		}
	}

	return false, nil
}
