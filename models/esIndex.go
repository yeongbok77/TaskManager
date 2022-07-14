package models

type IssueInfo struct {
	IssueId      int64    `json:"issue_id""`
	IssueContent string   `json:"issue_content"`
	Tags         []string `json:"tags,omitempty"`
	Comments     []string `json:"comments,omitempty"`
}
