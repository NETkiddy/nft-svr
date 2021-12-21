package jsn

//=============================================================================
type CreateTokenClassesRequest struct {
	RequestId     string
	Name          string
	Description   string
	Total         string
	Renderer      string
	CoverImageUrl string
}
type CreateTokenClassesResponse struct {
	Name        string
	Description string
	Issued      string
	Renderer    string
	Uuid        string
	Total       string
	Tags        []string
}

//=============================================================================
type GetTokenClassesRequest struct {
	RequestId string
	TokenUuid string
}
type GetTokenClassesResponse struct {
	Name         string
	Description  string
	Issued       string
	Renderer     string
	Uuid         string
	Total        string
	VerifiedInfo VerifiedInfo
	Tags         []string
}
type VerifiedInfo struct {
	IsVerified     bool
	VerifiedTitle  string
	VerifiedSource string
}
