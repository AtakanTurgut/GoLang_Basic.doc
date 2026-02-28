// custom exception
package error_handling

import "fmt"

type BorderException struct {
	parameter int
	message   string
}

func (e BorderException) Error() string {
	return fmt.Sprintf("%d: %s", e.parameter, e.message)
}

func TahminEtException(tahmin int) (string, error) {
	if tahmin < 0 || tahmin > 100 {
		return "", &BorderException{parameter: tahmin, message: "tahmin 0-100 aralığında olmalı"}
	}

	return "tahmin başarılı", nil
}
