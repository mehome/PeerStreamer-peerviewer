// Code generated by "stringer -type configStreamKind config.go"; DO NOT EDIT

package main

import "fmt"

const configStreamKindName = "configStreamKindVideoWebMconfigStreamKindAudioOpus"

var configStreamKindIndex = [...]uint8{0, 25, 50}

func (i configStreamKind) String() string {
	if i < 0 || i >= configStreamKind(len(configStreamKindIndex)-1) {
		return fmt.Sprintf("configStreamKind(%d)", i)
	}
	return configStreamKindName[configStreamKindIndex[i]:configStreamKindIndex[i+1]]
}