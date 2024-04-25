package dto

type UpsertPostRequest struct {
	PostContent string `json:"post_content"`
}

type PostResponse struct {
	ID          string `json:"id"`
	PostContent string `json:"post_content"`
}

type UpdateResponse struct {
	Message string `json:"message"`
}
