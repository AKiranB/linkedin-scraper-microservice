package routes

import (
	"fmt"
	"net/http"
	"net/url"

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

func CreateQueryParams(body Body) QueryParams {
	return QueryParams{
		Keywords:        body.Keywords,
		Location:        body.Location,
		DatePostedRange: GetDatePostedRange(body.DateSincePosted),
		Salary:          GetSalary(body.Salary),
		ExperienceLevel: GetExperienceLevel(body.ExperienceLevel),
		RemoteFilter:    GetRemoteFilter(body.RemoteType),
		JobType:         GetJobType(body.JobType),
		EasyApply:       GetEasyApply(body.EasyApply),
		FewApplicants:   GetFewApplicants(body.FewApplicants),
		PostalPlaceID:   GetPostalPlaceID(body.PostalPlaceID),
		CompanyID:       GetCompanyID(body.CompanyID),
		Start:           body.Start,
		SortBy:          body.SortBy,
	}
}

func JobsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := utils.Decode[Body](r)

		if err != nil {
			utils.Encode(w, http.StatusBadRequest, "Invalid request body")
		}

		queryParams := CreateQueryParams(body)

		client := NewLinkedInClient()

	}
}

func (c *LinkedInClient) CreateUrl(queryparams QueryParams) string {
	base, err := url.Parse(c.BaseUrl)

	if err != nil {
		fmt.Println("Error parsing base URL:", err)
	}

	params := url.Values{}
	params.Add("q", "this will get encoded as well")
	base.RawQuery = params.Encode()

	fmt.Printf("Encoded URL is %q\n", base.String())
}
