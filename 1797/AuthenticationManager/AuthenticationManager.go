package AuthenticationManager

type AuthenticationManager struct {
	verificationCode map[string]int
	timeToLive       int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		verificationCode: map[string]int{},
		timeToLive:       timeToLive,
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.verificationCode[tokenId] = currentTime
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if _, ok := this.verificationCode[tokenId]; !ok {
		return
	} else if currentTime-this.verificationCode[tokenId] >= this.timeToLive {
		delete(this.verificationCode, tokenId)
	} else {
		this.verificationCode[tokenId] = currentTime
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	var count int
	for authCode, logTime := range this.verificationCode {
		if currentTime-logTime < this.timeToLive {
			count++
		} else {
			delete(this.verificationCode, authCode)
		}
	}
	return count
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
