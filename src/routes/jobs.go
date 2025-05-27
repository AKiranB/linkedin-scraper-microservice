package routes

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/AKiranB/linkedin-scraper-microservice/src/utils"
)

type LinkedInClient struct {
	BaseURL string
	Headers map[string]string
}

func NewLinkedInClient() *LinkedInClient {
	return &LinkedInClient{
		BaseURL: "https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search",
		Headers: map[string]string{
			"Accept":           "application/json, text/javascript, */*; q=0.01",
			"Accept-Language":  "en-US,en;q=0.9",
			"Accept-Encoding":  "gzip, deflate, br",
			"Referer":          "https://www.linkedin.com/jobs",
			"X-Requested-With": "XMLHttpRequest",
			"Connection":       "keep-alive",
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
		EasyApply:       strconv.FormatBool(body.EasyApply),
		FewApplicants:   strconv.FormatBool(body.FewApplicants),
		PostalPlaceID:   GetPostalPlaceID(body.PostalPlaceID),
		CompanyID:       GetCompanyID(body.CompanyID),
		Start:           body.Start,
		SortBy:          body.SortBy,
	}
}

func buildQueryValues(qp QueryParams) url.Values {
	values := url.Values{}

	values.Add("start", strconv.Itoa(qp.Start))

	if qp.Keywords != "" {
		values.Add("keywords", qp.Keywords)
	}
	if qp.Location != "" {
		values.Add("location", qp.Location)
	}

	values.Add("f_AL", qp.EasyApply)
	values.Add("f_JIYN", qp.FewApplicants)

	if qp.ExperienceLevel != "" {
		values.Add("f_E", qp.ExperienceLevel)
	}
	if qp.JobType != "" {
		values.Add("f_JT", qp.JobType)
	}
	if qp.RemoteFilter != "" {
		values.Add("f_WT", qp.RemoteFilter)
	}

	values.Add("f_PP", qp.PostalPlaceID)
	values.Add("f_C", qp.CompanyID)

	values.Add("f_TPR", qp.DatePostedRange)
	if qp.Salary != "" {
		values.Add("f_SB2", qp.Salary)
	}

	if qp.SortBy != "" {
		values.Add("sortBy", qp.SortBy)
	}

	return values
}

func (c *LinkedInClient) CreateURL(qp QueryParams) string {
	base, err := url.Parse(c.BaseURL)
	if err != nil {
		fmt.Println("Error parsing base URL:", err)
		return c.BaseURL
	}
	values := buildQueryValues(qp)
	base.RawQuery = values.Encode()
	return base.String()
}

func JobsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := utils.Decode[Body](r)
		if err != nil {
			utils.Encode(w, http.StatusBadRequest, "Invalid request body")
			return
		}

		queryParams := CreateQueryParams(body)
		client := NewLinkedInClient()
		requestURL := client.CreateURL(queryParams)

		req, err := http.NewRequest("GET", requestURL, nil)
		if err != nil {
			utils.Encode(w, http.StatusInternalServerError, "Failed to create request")
			return
		}
		for k, v := range client.Headers {
			req.Header.Set(k, v)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			utils.Encode(w, http.StatusBadGateway, "Failed to fetch LinkedIn data")
			return
		}
		defer resp.Body.Close()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
	}
}
