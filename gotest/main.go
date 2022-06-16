package main

/* import (
	trieMap "com.grpk.utils/trie/triemap"
) */

func main() {
	var p Person
	p = &Member{}
}

type MemberBasicInfo struct {
}

type Member struct {
}

type Person interface {
	BasicInfo() MemberBasicInfo
}

func (member *Member) BasicInfo() MemberBasicInfo {

}
