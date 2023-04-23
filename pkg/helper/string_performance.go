package helper

import "strings"

func StringBuildString(ss ...string) string {
	var build strings.Builder

	for _, v := range ss {
		build.WriteString(v)
	}
	return build.String()
}

func StringBuildByte(bs ...byte) string {
	var build strings.Builder

	for _, v := range bs {
		build.WriteByte(v)
	}
	return build.String()
}
