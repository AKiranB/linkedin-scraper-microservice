package routes

import "testing"

func TestCreateQueryParams(t *testing.T) {

	expectedUrl := "https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?start=0&keywords=Test%20Job&location=United%20States&f_AL=false&f_E=2&f_JT=F&f_WT=1&f_JIYN=false&f_PP=&f_C=&f_TPR="

	body := Body{
		Start:           0,
		Keywords:        "Test Job",
		Location:        "United States",
		EasyApply:       true,
		ExperienceLevel: "4",
		JobType:         "C",
		RemoteType:      "3",
		FewApplicants:   true,
		PostalPlaceID:   123456,    // placeholder city ID
		CompanyID:       789012,    // placeholder company ID
		DateSincePosted: "r604800", // past week
	}

	queryParams := CreateQueryParams(body)

	if queryParams.Start != body.Start ||
		queryParams.Keywords != body.Keywords ||
		queryParams.Location != body.Location ||
		queryParams.EasyApply != "false" || // EasyApply is false in the expected URL
		queryParams.ExperienceLevel != "2" || // ExperienceLevel is set to 2 (Entry level)
		queryParams.JobType != "F" || // JobType is set to Full-time
		queryParams.RemoteFilter != "1" || // RemoteFilter is set to Onsite
		queryParams.FewApplicants != "true" ||
		queryParams.PostalPlaceID != "" || // PostalPlaceID is empty
		queryParams.CompanyID != "" {
		t.Errorf("Expected query params to match, got: %+v", queryParams)
	}

	url := NewLinkedInClient().CreateUrl(queryParams)

	if url != expectedUrl {
		t.Errorf("Expected URL to be %s, got %s", expectedUrl, url)
	}
}
