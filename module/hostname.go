package module

type ShowHostname struct {
	FQDN          string             `json:"fqdn"`
	Hostname  		string  					 `json:"hostname"`
}

func (b *ShowHostname) GetCmd() string {
	return "show hostname"
}