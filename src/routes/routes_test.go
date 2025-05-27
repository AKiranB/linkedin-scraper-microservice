package routes

import (
	"net/url"
	"reflect"
	"testing"
)

func TestCreateQueryParams(t *testing.T) {

	expectedUrl := "https://www.linkedin.com/jobs-guest/jobs/api/seeMoreJobPostings/search?start=0&keywords=Test%20Job&location=United%20States&f_AL=false&f_E=2&f_JT=F&f_WT=1&f_JIYN=false&f_PP=&f_C=&f_TPR="

	body := Body{
		Start:           0,
		Keywords:        "Test Job",
		Location:        "United States",
		EasyApply:       false,
		FewApplicants:   false,
		ExperienceLevel: "entry level",
		JobType:         "full-time",
		RemoteType:      "onsite",
		PostalPlaceID:   0,
		CompanyID:       0,
		DateSincePosted: "",
		Salary:          "",
		SortBy:          "",
	}

	queryParams := CreateQueryParams(body)

	actualURL := NewLinkedInClient().CreateURL(queryParams)

	got, _ := url.Parse(actualURL)
	want, _ := url.Parse(expectedUrl)

	if !reflect.DeepEqual(got.Query(), want.Query()) {
		t.Errorf("Query params = %v; want %v", got.Query(), want.Query())
	}
}
