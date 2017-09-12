package unifi

import "fmt"

// Wlans returns all of the wireless networks for a specified site name.
func (c *Client) Wlans(siteName string) ([]*Wlan, error) {
	var v struct {
		Wlans []*Wlan `json:"data"`
	}

	req, err := c.newRequest(
		"GET",
		fmt.Sprintf("api/s/%s/rest/wlanconf", siteName),
		nil,
	)
	if err != nil {
		return nil, err
	}

	_, err = c.do(req, &v)
	return v.Wlans, err
}

// A Wlan is the raw structure of a wireless network returned from the UniFi Controller
// API.
type Wlan struct {
	ID               string        `json:"_id"`
	BcFilterEnabled  bool          `json:"bc_filter_enabled"`
	BcFilterList     []interface{} `json:"bc_filter_list"`
	DtimMode         string        `json:"dtim_mode"`
	DtimNa           int           `json:"dtim_na"`
	DtimNg           int           `json:"dtim_ng"`
	Enabled          bool          `json:"enabled"`
	MacFilterEnabled bool          `json:"mac_filter_enabled"`
	MacFilterList    []interface{} `json:"mac_filter_list"`
	MacFilterPolicy  string        `json:"mac_filter_policy"`
	Name             string        `json:"name"`
	RatectrlNa12     string        `json:"ratectrl_na_12"`
	RatectrlNa18     string        `json:"ratectrl_na_18"`
	RatectrlNa24     string        `json:"ratectrl_na_24"`
	RatectrlNa36     string        `json:"ratectrl_na_36"`
	RatectrlNa48     string        `json:"ratectrl_na_48"`
	RatectrlNa54     string        `json:"ratectrl_na_54"`
	RatectrlNa6      string        `json:"ratectrl_na_6"`
	RatectrlNa9      string        `json:"ratectrl_na_9"`
	RatectrlNaMode   string        `json:"ratectrl_na_mode"`
	RatectrlNg12     string        `json:"ratectrl_ng_12"`
	RatectrlNg18     string        `json:"ratectrl_ng_18"`
	RatectrlNg24     string        `json:"ratectrl_ng_24"`
	RatectrlNg36     string        `json:"ratectrl_ng_36"`
	RatectrlNg48     string        `json:"ratectrl_ng_48"`
	RatectrlNg54     string        `json:"ratectrl_ng_54"`
	RatectrlNg6      string        `json:"ratectrl_ng_6"`
	RatectrlNg9      string        `json:"ratectrl_ng_9"`
	RatectrlNgCck1   string        `json:"ratectrl_ng_cck_1"`
	RatectrlNgCck11  string        `json:"ratectrl_ng_cck_11"`
	RatectrlNgCck2   string        `json:"ratectrl_ng_cck_2"`
	RatectrlNgCck55  string        `json:"ratectrl_ng_cck_5_5"`
	RatectrlNgMode   string        `json:"ratectrl_ng_mode"`
	Schedule         []string      `json:"schedule"`
	Security         string        `json:"security"`
	SiteID           string        `json:"site_id"`
	UsergroupID      string        `json:"usergroup_id"`
	Vlan             string        `json:"vlan"`
	VlanEnabled      bool          `json:"vlan_enabled"`
	WepIdx           int           `json:"wep_idx"`
	WlangroupID      string        `json:"wlangroup_id"`
	WpaEnc           string        `json:"wpa_enc"`
	WpaMode          string        `json:"wpa_mode"`
	XIappKey         string        `json:"x_iapp_key"`
	XPassphrase      string        `json:"x_passphrase"`
	ScheduleEnabled  bool          `json:"schedule_enabled"`
}
