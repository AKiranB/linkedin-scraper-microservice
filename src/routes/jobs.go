package routes

import (
	"net/http"
	"strings"

	"github.com/AKiranB/linkedin-scraper-microservice/src/utils"
)

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

type Body struct {
	Keywords        string `json:"keywords"`
	Location        string `json:"location"`
	DateSincePosted string `json:"date_since_posted,omitempty"`
	Salary          string `json:"salary,omitempty"`
	ExperienceLevel string `json:"experience_level,omitempty"`
	RemoteType      string `json:"remote_type,omitempty"`
	JobType         string `json:"job_type,omitempty"`
	Start           int    `json:"start,omitempty"`
	SortBy          string `json:"sort_by,omitempty"`
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

func getDateSincePosted(dateSincePosted string) string {
	dateRange := map[string]string{
		"past month": "r2592000",
		"past week":  "r604800",
		"24hr":       "r86400",
		"8hr":        "r28800",
		"1hr":        "r3600",
	}
	return dateRange[format(dateSincePosted)]
}

func getSalary(salary string) string {
	salaryRange := map[string]string{
		"40000": "1",
		"60000": "2",
		"80000": "3",
		"100000": "4",
		"120000": "5",
	}
	return salaryRange[format(salary)]
}


func JobsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := utils.Decode[Body](r)

		if err != nil {
			utils.Encode(w, http.StatusBadRequest, "Invalid request body")
		}

		queryParams := QueryParams{
			Keywords: body.Keywords,
			Location: body.Location,
			Start:    body.Start,
			SortBy:   body.SortBy,
			Salary:  getSalary(body.Salary),
			ExperienceLevel: getExperienceLevel(body.ExperienceLevel),
			RemoteFilter: getRemoteFilter(body.RemoteType),
			JobType: getJobType(body.JobType),
			DateSincePosted: getDateSincePosted(body.DateSincePosted),
		}

	}
}

func 
