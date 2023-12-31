// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateJobListing struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	URL         string `json:"url"`
}

type DeleteJobRespone struct {
	DeleteJobID string `json:"deleteJobId"`
}

type JobListing struct {
	ID          string `json:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Company     string `json:"company"`
	URL         string `json:"url"`
}

type UpdateJobListing struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Company     *string `json:"company,omitempty"`
	URL         *string `json:"url,omitempty"`
}
