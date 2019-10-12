package lengthOfLongestSubstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	if lengthOfLongestSubstring("abcabcbb") != 3 {
		t.Error("abcabcbb err")
	}
	if lengthOfLongestSubstring("bbbbb") != 1 {
		t.Error("bbbbb err", lengthOfLongestSubstring("bbbbb"))
	}
	if lengthOfLongestSubstring("pwwkew") != 3 {
		t.Error("pwwkew err", lengthOfLongestSubstring("pwwkew"))
	}

	if lengthOfLongestSubstring("aab") != 2 {
		t.Error("aab err")
	}
	if lengthOfLongestSubstring("dvdf") != 3 {
		t.Error("dvdf err")
	}
	if lengthOfLongestSubstring(" ") != 1 {
		t.Error("  err", lengthOfLongestSubstring(" "))
	}

	if lengthOfLongestSubstring("au") != 2 {
		t.Error("au  err", lengthOfLongestSubstring("au"))
	}
	if lengthOfLongestSubstring("asjrgapa") != 6 {
		t.Error("asjrgapa  err", lengthOfLongestSubstring("asjrgapa"))
	}
}
