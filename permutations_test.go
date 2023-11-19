package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestPermsEmptyInput(t *testing.T) {
	result := getPermutation("")
	expected := []string{}
	sort.Strings(result)
	sort.Strings(expected)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("getPermutation('') expected %v, but got %v", expected, result)
	}
}

func TestPerms1Input(t *testing.T) {
	result := getPermutation("a")
	expected := []string{"a"}
	sort.Strings(result)
	sort.Strings(expected)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("getPermutation('a') expected %v, but got %v", expected, result)
	}
}

func TestPerms2Input(t *testing.T) {
	result := getPermutation("ab")
	expected := []string{"ab", "ba"}
	sort.Strings(result)
	sort.Strings(expected)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("getPermutation('ab') expected %v, but got %v", expected, result)
	}
}

func TestPerms3Input(t *testing.T) {
	result := getPermutation("abc")
	expected := []string{"abc", "acb", "bac", "bca", "cab", "cba"}
	sort.Strings(result)
	sort.Strings(expected)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("getPermutation('abc') expected %v, but got %v", expected, result)
	}
}

func TestPerms4InputWithDup(t *testing.T) {
	result := getPermutation("aabb")
	expected := []string{"aabb", "abab", "abba", "baab", "baba", "bbaa"}
	sort.Strings(result)
	sort.Strings(expected)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("getPermutation('abc') expected %v, but got %v", expected, result)
	}
}
