package job

import (
	"github.com/imroc/req/v3"
	"golang.org/x/exp/rand"
	"regexp"
	"strings"
	"time"
)

func (j *Job) getHashForm(r *req.Response) string {
	if strings.Contains(r.String(), "formhash") == true {
		rule := regexp.MustCompile(` name="formhash" value="([\da-zA-z]+)"`)
		result := rule.FindStringSubmatch(r.String())
		if len(result) <= 1 {
			return ""
		}
		return result[len(result)-1]
	}

	return ""
}

func (j *Job) getCoin(r *req.Response) string {
	rule := regexp.MustCompile(`金钱: <span id="hcredit_2">([0-9]+)`)
	result := rule.FindStringSubmatch(r.String())
	return result[len(result)-1]
}

func (j *Job) randomUID(count, max int) []int {
	rand.Seed(uint64(time.Now().UnixNano())) // 用当前时间作为种子
	numbers := make(map[int]struct{})
	var result []int

	for len(result) < count {
		num := rand.Intn(max)
		if _, exists := numbers[num]; !exists {
			numbers[num] = struct{}{}
			result = append(result, num)
		}
	}

	return result
}
