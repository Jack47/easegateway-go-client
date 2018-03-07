package pdu

type PluginsRetrieveClusterResponse struct {
	Plugins    []PluginSpec   `json:"plugins"`
	Pagination paginationInfo `json:"pagination"`
}
