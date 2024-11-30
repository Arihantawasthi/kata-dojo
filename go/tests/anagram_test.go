package tests

import (
	"kata-dojo/go/structures"
	"testing"
)

func TestAnagram(t *testing.T) {
    strs := []string{"eat", "tea", "ate", "bat", "nat", "tan"}
    result := structures.GroupAnagrams(strs)
    t.Log(result)
}
