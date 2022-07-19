package main

import (
	"fmt"

	"k8s.io/kubectl/pkg/util/templates"

	"github.com/Fish-pro/i18n-demo/pkg/util/i18n"
)

var sss = templates.LongDesc(i18n.T(`test sss`))

func main() {
	i18n.LoadTranslations("karmadactl", nil)

	fmt.Println(i18n.T("test haha"))
	fmt.Println(i18n.T("aaaaa"))
	fmt.Println(i18n.T("bbbbb"))
	fmt.Println(i18n.T("ccccc"))
	fmt.Println(sss)
}
