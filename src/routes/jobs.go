package routes

import "strings"

type LinkedInClient struct {
	BaseUrl string
	Headers map[string]string
}

func NewLinkedInClient() *LinkedInClient {
	return &LinkedInClient{
		BaseUrl: "https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search",
		Headers: map[string]string{
			"Accept":           "application/json, text/javascript, */*; q=0.01",
			"Accept-Language":  "en-US,en;q=0.9",
			"Accept-Encoding":  "gzip, deflate, br",
			"Referer":          "https://www.linkedin.com/jobs",
			"X-Requested-With": "XMLHttpRequest",
			"Connection":       "keep-alive",
			"Sec-Fetch-Dest":   "empty",
			"Sec-Fetch-Mode":   "cors",
			"Sec-Fetch-Site":   "same-origin",
			"Cache-Control":    "no-cache",
			"Pragma":           "no-cache",
		},
	}
}

type QueryParams struct {
	Keywords        string `url:"keywords"`
	Location        string `url:"location"`
	DateSincePosted string `url:"f_TPR,omitempty"`
	Salary          string `url:"f_SB2,omitempty"`
	ExperienceLevel string `url:"f_E,omitempty"`
	RemoteFilter    string `url:"f_WT,omitempty"`
	JobType         string `url:"f_JT,omitempty"`
	Start           int    `url:"start"`
	SortBy          string `url:"sortBy,omitempty"`
}

func format(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}

func getExperienceLevel(experienceLevel string) string {
	experienceRange := map[string]string{
		"internship":  "1",
		"entry level": "2",
		"associate":   "3",
		"senior":      "4",
		"director":    "5",
		"executive":   "6",
	}
	key := format(experienceLevel)
	return experienceRange[key]
}

func getJobType(jobType string) string {
	jobTypeRange := map[string]string{
		"full-time":  "F",
		"part-time":  "P",
		"contract":   "C",
		"temporary":  "T",
		"volunteer":  "V",
		"internship": "I",
	}
	key := format(jobType)
	return jobTypeRange[key]
}

func getRemoteFilter(remoteType string) string {
	remoteFilterRange := map[string]string{
		"onsite": "1",
		"remote": "2",
		"hybrid": "3",
	}
	return remoteFilterRange[format(remoteType)]
}
