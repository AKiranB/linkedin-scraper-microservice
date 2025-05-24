package routes

import (
	"strconv"
	"strings"
)

func GetExperienceLevel(experienceLevel string) string {
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

func GetJobType(jobType string) string {
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

func GetRemoteFilter(remoteType string) string {
	remoteFilterRange := map[string]string{
		"onsite": "1",
		"remote": "2",
		"hybrid": "3",
	}
	return remoteFilterRange[format(remoteType)]
}

func GetEasyApply(easyApply bool) string {
	if easyApply {
		return "true"
	}
	return ""
}

func GetFewApplicants(fewApplicants bool) string {
	if fewApplicants {
		return "true"
	}
	return ""
}

func GetPostalPlaceID(placeID int) string {
	if placeID > 0 {
		return strconv.Itoa(placeID)
	}
	return ""
}

func GetCompanyID(companyID int) string {
	if companyID > 0 {
		return strconv.Itoa(companyID)
	}
	return ""
}

func GetDatePostedRange(datePostedRange string) string {
	if datePostedRange != "" {
		return datePostedRange
	}
	return ""
}

func GetDateSincePosted(dateSincePosted string) string {
	dateRange := map[string]string{
		"past month": "r2592000",
		"past week":  "r604800",
		"24hr":       "r86400",
		"8hr":        "r28800",
		"1hr":        "r3600",
	}
	return dateRange[format(dateSincePosted)]
}

func GetSalary(salary string) string {
	salaryRange := map[string]string{
		"40000":  "1",
		"60000":  "2",
		"80000":  "3",
		"100000": "4",
		"120000": "5",
	}
	return salaryRange[format(salary)]
}

func format(str string) string {
	return strings.ToLower(strings.TrimSpace(str))
}
