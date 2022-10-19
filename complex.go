package cute

// class of rune
type rune_class int

// enum
const (
	ignore rune_class = iota // don't count it
	normal                   // count it as 1
	double                   // count it as 2
)

// Arabic Special Characters
func isArabicTashkeel(r rune) (ok bool, class rune_class) {
	switch r {
	// tashkeel
	case '\u064b', '\u064c', '\u064d', '\u064e', '\u064f', '\u0650', '\u0651', '\u0652':
		ok = true
		class = ignore
		return
	}
	return
}

/*
	CJK Chinese Characters:
	Hànzì in Chinese, Kanji and Kana in Japanese, Hanja and Hangul in Korean.
*/
func isCJK(r rune) (ok bool, class rune_class) {
	// every rune is bold
	for i := '\u4e00'; i <= '\u9FFF'; i++ {
		if r == i {
			ok = true
			class = double
			return
		}
	}
	return
}

// Japanese Hiragana Characters
func isJapanese(r rune) (ok bool, class rune_class) {
	// special runes
	if r == '\u30FC' || r == '\u3001' || r == '\u3002' {
		ok = true
		class = double
		return
	}
	// all runes
	for i := '\u3040'; i <= '\u309f'; i++ {
		if r == i {
			ok = true
			// ignore runes
			if r == '\u3040' || r == '\u3097' || r == '\u3098' || r == '\u3099' || r == '\u309a' {
				class = ignore
				return
			}
			// bold rune
			class = double
			return
		}
	}
	return
}

// Hindi Devanagari Characters
func isHindi(r rune) (ok bool, class rune_class) {
	for i := '\u0900'; i <= '\u097F'; i++ {
		if r == i {
			ok = true
			switch r {
			case '\u0900', '\u0901', '\u0902', '\u0903':
				// ignore
				class = ignore
				return
			case '\u093a', '\u093c':
				// ignore
				class = ignore
				return
			case '\u0941', '\u0942', '\u0943', '\u0944', '\u0945', '\u0946', '\u0947', '\u0948', '\u094d':
				// ignore
				class = ignore
				return
			case '\u0951', '\u0952', '\u0953', '\u0954', '\u0955', '\u0956', '\u0957':
				// ignore
				class = ignore
				return
			case '\u0962', '\u0963':
				// ignore
				class = ignore
				return
			case '\u0970', '\u0971':
				// ignore
				class = ignore
				return
			default:
				// normal rune
				class = normal
				return
			}
		}
	}
	return
}

// Emoji Characters (include korean lang)
func isEmoji(r rune) (ok bool, class rune_class) {
	switch r {
	//﴾ ﴿ ﷺ ﷻ
	case '\uFD3E', '\uFD3F', '\uFDFA', '\uFDFB':
		ok = true
		class = normal
		return
	}
	// Emoticons
	for i := 0x1f601; i <= 0x1f64f; i++ {
		if r == rune(i) {
			ok = true
			class = double
			return
		}
	}
	// Dingbats
	for i := 0x2702; i <= 0x27B0; i++ {
		if r == rune(i) {
			ok = true
			class = double
			return
		}
	}
	// Transport and map
	for i := 0x1F680; i <= 0x1F6C0; i++ {
		if r == rune(i) {
			ok = true
			class = double
			return
		}
	}
	// Enclosed
	for i := 0x24C2; i <= 0x1F251; i++ {
		if r == rune(i) {
			ok = true
			class = double
			return
		}
	}
	// Others ..
	for i := 0x1F30D; i <= 0x1F567; i++ {
		if r == rune(i) {
			ok = true
			class = double
			return
		}
	}
	return
}

// calculate how many characters in a text
func calculateLength(text string) (length int) {
	for _, r := range text {
		// chinese cjk characters
		ok, class := isCJK(r)
		if ok {
			switch class {
			case ignore:
				continue
			case normal:
				length += 1
				continue
			case double:
				length += 2
				continue
			}
		}
		// emoji characters
		ok, class = isEmoji(r)
		if ok {
			switch class {
			case ignore:
				continue
			case normal:
				length += 1
				continue
			case double:
				length += 2
				continue
			}
		}
		// japanese characters
		ok, class = isJapanese(r)
		if ok {
			switch class {
			case ignore:
				continue
			case normal:
				length += 1
				continue
			case double:
				length += 2
				continue
			}
		}
		// hindi characters
		ok, class = isHindi(r)
		if ok {
			switch class {
			case ignore:
				continue
			case normal:
				length += 1
				continue
			case double:
				length += 2
				continue
			}
		}
		// arabic tashkeel characters
		ok, class = isArabicTashkeel(r)
		if ok {
			switch class {
			case ignore:
				continue
			case normal:
				length += 1
				continue
			case double:
				length += 2
				continue
			}
		}
		// other normal characters
		length++
	}
	return
}
