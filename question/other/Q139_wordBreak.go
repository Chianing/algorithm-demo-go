package other

// 单词拆分
// https://leetcode.cn/problems/word-break/
func wordBreak(s string, wordDict []string) bool {
	if len(wordDict) == 0 && s != "" {
		return false
	}

	dictMap := make(map[string]bool)
	for _, v := range wordDict {
		dictMap[v] = true
	}

	ret := false

	wordCacheMap := make(map[int]bool)
	var backtrace func(int) bool
	backtrace = func(index int) bool {
		if index == len(s) {
			return true
		}
		if cacheRet, ok := wordCacheMap[index]; ok {
			return cacheRet
		}

		for i := index + 1; i <= len(s); i++ {
			subStr := s[index:i]
			if dictMap[subStr] && backtrace(i) {
				wordCacheMap[index] = true
				return true
			}
		}

		wordCacheMap[index] = false
		return false

	}
	ret = backtrace(0)

	return ret

}
