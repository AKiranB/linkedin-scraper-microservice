package routes

type QueryParams struct {
	Keywords        string `url:"keywords"`
	Location        string `url:"location"`
	DatePostedRange string `url:"f_TPR,omitempty"`
	Salary          string `url:"f_SB2,omitempty"`
	ExperienceLevel string `url:"f_E,omitempty"`
	RemoteFilter    string `url:"f_WT,omitempty"`
	JobType         string `url:"f_JT,omitempty"`
	EasyApply       string `url:"f_AL,omitempty"`
	FewApplicants   string `url:"f_JIYN,omitempty"`
	PostalPlaceID   string `url:"f_PP,omitempty"`
	CompanyID       string `url:"f_C,omitempty"`

	Start  int    `url:"start,omitempty"`
	SortBy string `url:"sortBy,omitempty"`
}

type Body struct {
	Keywords string `json:"keywords"` // required
	Location string `json:"location"` // required

	DateSincePosted string `json:"date_since_posted,omitempty"` // user-friendly (mapped to f_TPR)
	Salary          string `json:"salary,omitempty"`            // user-friendly (mapped to f_SB2)
	ExperienceLevel string `json:"experience_level,omitempty"`  // (mapped to f_E)
	RemoteType      string `json:"remote_type,omitempty"`       // (mapped to f_WT)
	JobType         string `json:"job_type,omitempty"`          // (mapped to f_JT)
	Start           int    `json:"start,omitempty"`             // (mapped to start)
	SortBy          string `json:"sort_by,omitempty"`           // (mapped to sortBy)

	EasyApply     bool `json:"easy_apply,omitempty"`      // (mapped to f_AL)
	FewApplicants bool `json:"few_applicants,omitempty"`  // (mapped to f_JIYN)
	PostalPlaceID int  `json:"postal_place_id,omitempty"` // (mapped to f_PP)
	CompanyID     int  `json:"company_id,omitempty"`      // (mapped to f_C)
}
