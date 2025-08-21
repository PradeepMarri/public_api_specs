package main

import (
	"github.com/community-api/mcp-server/config"
	"github.com/community-api/mcp-server/models"
	tools_user_content "github.com/community-api/mcp-server/tools/user_content"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_user_content.CreateGet_user_content_url_jsonTool(cfg),
		tools_user_content.CreateGet_user_content_user_jsonTool(cfg),
		tools_user_content.CreateGet_user_content_by_date_jsonTool(cfg),
		tools_user_content.CreateGet_user_content_recent_jsonTool(cfg),
	}
}
