// TEMPORARY AUTOGENERATED FILE: easyjson stub code to make the package
// compilable during generation.

package  main

import (
  "github.com/mailru/easyjson/jwriter"
  "github.com/mailru/easyjson/jlexer"
)

func ( Circle ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* Circle ) UnmarshalJSON([]byte) error { return nil }
func ( Circle ) MarshalEasyJSON(w *jwriter.Writer) {}
func (* Circle ) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_Circle *Circle

func ( Point ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* Point ) UnmarshalJSON([]byte) error { return nil }
func ( Point ) MarshalEasyJSON(w *jwriter.Writer) {}
func (* Point ) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_Point *Point

func ( Rectangle ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* Rectangle ) UnmarshalJSON([]byte) error { return nil }
func ( Rectangle ) MarshalEasyJSON(w *jwriter.Writer) {}
func (* Rectangle ) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_Rectangle *Rectangle
