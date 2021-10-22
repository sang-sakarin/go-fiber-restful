package news

import "go-fiber-restful/apps/commons"

type News struct {
	commons.SoftControlModel `pg:"override"`
	Title                    string `json:"title"`
}
