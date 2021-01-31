package vernam

import (
	"math/rand"
	"time"
)

//genera una chiave randomica
func RandomKey(leng int) string {
	rand.Seed(time.Now().UnixNano())
	var char = []rune("abcdefghijklmnopqrstuvwxyz")
	key := make([]rune, leng)
	for i := range key {
		key[i] = char[rand.Intn(len(char))]
	}
	return string(key)
}

//separa una stringa in un array si stringhe
func splitSting(msg string) []string {
	var msgsplitted []string
	for a := 0; a < len(msg); a++ {
		msgsplitted = append(msgsplitted, string(msg[a]))
	}
	return msgsplitted
}

//unisce un array di stringhe in una stringa
func mergeString(msg []string) string {
	var msgmerged string
	for a := 0; a < len(msg); a++ {
		msgmerged += msg[a]
	}
	return msgmerged
}

//trasforma il messaggio in numeri --> a=0, b=1, c=2, d=3...
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

// trasforma numeri nel messaggio --> 0=a, 1=b, 2=c, 3=c...
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

// cripta il messaggio attraverso una chiave
func Encrypt(msg string, key string) string {
	var msgindex []int = lettersToIndex(splitSting(msg))
	var keyindex []int = lettersToIndex(splitSting(key))
	var msgcrypted []int
	for a := range msg {
		if msgindex[a]+keyindex[a] < 26 {
			msgcrypted = append(msgcrypted, msgindex[a]+keyindex[a])
		} else {
			msgcrypted = append(msgcrypted, msgindex[a]+keyindex[a]-26)
		}
	}
	return mergeString(indexToLetters(msgcrypted))
}

// decripta il messaggio criptato attraverso la chiave
func Decrypt(msg string, key string) string {
	var msgindex []int = lettersToIndex(splitSting(msg))
	var keyindex []int = lettersToIndex(splitSting(key))
	var msgdecrypted []int
	for a := range msg {
		if msgindex[a]-keyindex[a] >= 0 {
			msgdecrypted = append(msgdecrypted, msgindex[a]-keyindex[a])
		} else {
			msgdecrypted = append(msgdecrypted, 26-keyindex[a]+msgindex[a])
		}
	}
	return mergeString(indexToLetters(msgdecrypted))
}

