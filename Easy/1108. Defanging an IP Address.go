package main

import "strings"

func defangIPaddr(address string) string {
	return (strings.Replace(address, ".", "[.]", -1))
}

// func defangIPaddr(address string) string {
//     var builder strings.Builder
//     for _, s := range address {
//         if s == '.' {
//             builder.WriteString("[.]")
//         } else {
//             builder.WriteByte(byte(s))
//         }
//     }
//     return builder.String()
// }
