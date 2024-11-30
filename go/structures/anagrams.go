package structures

func GroupAnagrams(strs []string) [][]string {
    var result [][]string
    dict := make(map[[26]int][]string)

    for _, word := range strs {
        var freqArr [26]int
        for _, letter := range word {
            freqArr[int(letter)-97] += 1
        }
        arr := []string{word}
        if _, ok := dict[freqArr]; !ok {
            dict[freqArr] = arr
            continue
        }
        dict[freqArr] = append(dict[freqArr], word)
    }

    for _, val := range dict {
        result = append(result, val)
    }

    return result
}
