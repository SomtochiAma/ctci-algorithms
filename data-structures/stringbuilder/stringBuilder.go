package stringbuilder

import "strings"

type StringBuilder struct {
	strings []string
}

func NewSB() *StringBuilder{
	return &StringBuilder{}
}

func(s *StringBuilder) Append(str string) {
	s.strings = append(s.strings, str)
}

func(s *StringBuilder) String() string{
	return strings.Join(s.strings, "")
}