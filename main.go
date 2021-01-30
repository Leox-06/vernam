package main

func splitSting(msg string) []string {
	var msgsplitted []string
	for a := 0; a < len(msg); a++ {
		msgsplitted = append(msgsplitted, string(msg[a]))
	}
	return msgsplitted
}

func mergeString(msg []string) string {
	var msgmerged string
	for a := 0; a < len(msg); a++ {
		msgmerged += msg[a]
	}
	return msgmerged
}

func lettersToIndex(msg []string) []int {
	char := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var msgindex []int
	for _, c := range msg {
		for i, v := range char {
			if c == v {
				msgindex = append(msgindex, i)
			}
		}
	}
	return msgindex
}

func indexToLetters(index []int) []string {
	char := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var indexvalue []string
	for _, c := range index {
		for i, v := range char {
			if c == i {
				indexvalue = append(indexvalue, v)
			}
		}
	}
	return indexvalue
}

func encrypt(msg string, key string) string {
	var msgindex []int = lettersToIndex(splitSting(msg))
	var keyindex []int = lettersToIndex(splitSting(key))
	var msgcrypted []int
	for a := 0; a < len(msgindex); a++ {
		if msgindex[a]+keyindex[a] < 26 {
			msgcrypted = append(msgcrypted, msgindex[a]+keyindex[a])
		} else {
			msgcrypted = append(msgcrypted, msgindex[a]+keyindex[a]-26)
		}
	}
	return mergeString(indexToLetters(msgcrypted))
}

func decrypt(msg string, key string) string {
	var msgindex []int = lettersToIndex(splitSting(msg))
	var keyindex []int = lettersToIndex(splitSting(key))
	var msgdecrypted []int
	for a := 0; a < len(msg); a++ {
		if msgindex[a]-keyindex[a] >= 0 {
			msgdecrypted = append(msgdecrypted, msgindex[a]-keyindex[a])
		} else {
			msgdecrypted = append(msgdecrypted, 26-keyindex[a]+msgindex[a])
		}
	}
	return mergeString(indexToLetters(msgdecrypted))
}

func main() {

}
