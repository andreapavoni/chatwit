package main

import (
  "html/template"
  "github.com/shaoshing/train"
)

func NewTemplate(partial, layout string) *template.Template {
  tpl := template.New(partial)
  tpl.Funcs(train.HelperFuncs)

  return template.Must(tpl.ParseFiles(partial, layout))
}
