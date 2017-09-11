package unifi

import "fmt"

// Devices returns all of the Devices for a specified site name.
func (c *Client) Devices(siteName string) ([]*Device, error) {
	var v struct {
		Devices []*Device `json:"data"`
	}

	req, err := c.newRequest(
		"GET",
		fmt.Sprintf("/api/s/%s/stat/device", siteName),
		nil,
	)
	if err != nil {
		return nil, err
	}

	_, err = c.do(req, &v)
	return v.Devices, err
}

// A Device is the raw structure of a Device returned from the UniFi Controller
// API.
type Device struct {
	ID                   string           `json:"_id"`
	Uptime               int              `json:"_uptime"`
	Adopted              bool             `json:"adopted"`
	AtfEnabled           bool             `json:"atf_enabled"`
	BandsteeringMode     string           `json:"bandsteering_mode"`
	BoardRev             int              `json:"board_rev"`
	Bytes                int64            `json:"bytes"`
	BytesD               int              `json:"bytes-d"`
	BytesR               int              `json:"bytes-r"`
	Cfgversion           string           `json:"cfgversion"`
	ConfigNetwork        []*ConfigNetwork `json:"config_network"`
	ConnectRequestIP     string           `json:"connect_request_ip"`
	ConnectRequestPort   string           `json:"connect_request_port"`
	DeviceID             string           `json:"device_id"`
	DownlinkTable        []*DownlinkTable `json:"downlink_table"`
	EthernetTable        []*EthernetTable `json:"ethernet_table"`
	FwCaps               int              `json:"fw_caps"`
	GuestNumSta          int              `json:"guest-num_sta"`
	GuestToken           string           `json:"guest_token"`
	HasEth1              bool             `json:"has_eth1"`
	HasSpeaker           bool             `json:"has_speaker"`
	InformAuthkey        string           `json:"inform_authkey"`
	InformIP             string           `json:"inform_ip"`
	InformURL            string           `json:"inform_url"`
	IP                   string           `json:"ip"`
	Isolated             bool             `json:"isolated"`
	KnownCfgversion      string           `json:"known_cfgversion"`
	LastSeen             int              `json:"last_seen"`
	LedOverride          string           `json:"led_override"`
	Locating             bool             `json:"locating"`
	Mac                  string           `json:"mac"`
	MapID                string           `json:"map_id"`
	Model                string           `json:"model"`
	NaChannel            int              `json:"na-channel"`
	NaEirp               int              `json:"na-eirp"`
	NaExtchannel         int              `json:"na-extchannel"`
	NaGain               int              `json:"na-gain"`
	NaGuestNumSta        int              `json:"na-guest-num_sta"`
	NaNumSta             int              `json:"na-num_sta"`
	NaState              string           `json:"na-state"`
	NaTxPower            int              `json:"na-tx_power"`
	NaUserNumSta         int              `json:"na-user-num_sta"`
	NaAstBeXmit          int              `json:"na_ast_be_xmit"`
	NaAstCst             *string          `json:"na_ast_cst"`
	NaAstTxto            *string          `json:"na_ast_txto"`
	NaCuSelfRx           int              `json:"na_cu_self_rx"`
	NaCuSelfTx           int              `json:"na_cu_self_tx"`
	NaCuTotal            int              `json:"na_cu_total"`
	NaTxPackets          int              `json:"na_tx_packets"`
	NaTxRetries          int              `json:"na_tx_retries"`
	Name                 string           `json:"name"`
	NgChannel            int              `json:"ng-channel"`
	NgEirp               int              `json:"ng-eirp"`
	NgExtchannel         int              `json:"ng-extchannel"`
	NgGain               int              `json:"ng-gain"`
	NgGuestNumSta        int              `json:"ng-guest-num_sta"`
	NgNumSta             int              `json:"ng-num_sta"`
	NgState              string           `json:"ng-state"`
	NgTxPower            int              `json:"ng-tx_power"`
	NgUserNumSta         int              `json:"ng-user-num_sta"`
	NgAstBeXmit          int              `json:"ng_ast_be_xmit"`
	NgAstCst             *string          `json:"ng_ast_cst"`
	NgAstTxto            *string          `json:"ng_ast_txto"`
	NgCuSelfRx           int              `json:"ng_cu_self_rx"`
	NgCuSelfTx           int              `json:"ng_cu_self_tx"`
	NgCuTotal            int              `json:"ng_cu_total"`
	NgTxPackets          int              `json:"ng_tx_packets"`
	NgTxRetries          int              `json:"ng_tx_retries"`
	NumSta               int              `json:"num_sta"`
	PortTable            []*string        `json:"port_table"`
	RadioNa              []*RadioNa       `json:"radio_na"`
	RadioNg              []*RadioNg       `json:"radio_ng"`
	RadioTable           []*RadioTable    `json:"radio_table"`
	RxBytes              int              `json:"rx_bytes"`
	RxBytesD             int              `json:"rx_bytes-d"`
	Scanning             bool             `json:"scanning"`
	Serial               string           `json:"serial"`
	SiteID               string           `json:"site_id"`
	SpectrumScanning     bool             `json:"spectrum_scanning"`
	SSHSessionTable      []*string        `json:"ssh_session_table"`
	Stat                 []*Stat          `json:"stat"`
	State                int              `json:"state"`
	SysStats             []*SysStats      `json:"sys_stats"`
	TxBytes              int64            `json:"tx_bytes"`
	TxBytesD             int              `json:"tx_bytes-d"`
	Type                 string           `json:"type"`
	Uplink               []*Uplink        `json:"uplink"`
	UplinkTable          []*UplinkTable   `json:"uplink_table"`
	Uptime               int              `json:"uptime"`
	UserNumSta           int              `json:"user-num_sta"`
	VapTable             []*VapTable      `json:"vap_table"`
	Version              string           `json:"version"`
	VwireEnabled         bool             `json:"vwireEnabled"`
	VwireTable           []*VwireTable    `json:"vwire_table"`
	WifiCaps             int              `json:"wifi_caps"`
	WlangroupIDNa        string           `json:"wlangroup_id_na"`
	WlangroupIDNg        string           `json:"wlangroup_id_ng"`
	X                    float64          `json:"x"`
	XAuthkey             string           `json:"x_authkey"`
	XFingerprint         string           `json:"x_fingerprint"`
	XHasSSHHostkey       bool             `json:"x_has_ssh_hostkey"`
	XVwirekey            string           `json:"x_vwirekey"`
	Y                    float64          `json:"y"`
	NgLastInterferenceAt int              `json:"ng_last_interference_at,omitempty"`
	PortStats            []*string        `json:"port_stats,omitempty"`
	Upgradable           bool             `json:"upgradable,omitempty"`
	UpgradeToFirmware    string           `json:"upgrade_to_firmware,omitempty"`
	DiscoveredVia        string           `json:"discovered_via,omitempty"`
	LastUplink           struct {
		UplinkMac string `json:"uplink_mac"`
	} `json:"last_uplink,omitempty"`
	UplinkApMac string `json:"uplink_ap_mac,omitempty"`
}

type ConfigNetwork struct {
	IP   string `json:"ip"`
	Type string `json:"type"`
}

type DownlinkTable struct {
	ApMac         string `json:"ap_mac"`
	AuthTime      int64  `json:"auth_time"`
	Authorized    bool   `json:"authorized"`
	Ccq           int    `json:"ccq"`
	Channel       int    `json:"channel"`
	DhcpendTime   int    `json:"dhcpend_time"`
	DhcpstartTime int    `json:"dhcpstart_time"`
	Hostname      string `json:"hostname"`
	Idletime      int    `json:"idletime"`
	Is11A         bool   `json:"is_11a"`
	Is11Ac        bool   `json:"is_11ac"`
	Is11N         bool   `json:"is_11n"`
	Mac           string `json:"mac"`
	Name          string `json:"name"`
	Noise         int    `json:"noise"`
	Radio         string `json:"radio"`
	Rssi          int    `json:"rssi"`
	RxBytes       int64  `json:"rx_bytes"`
	RxBytesR      int    `json:"rx_bytes-r"`
	RxMcast       int    `json:"rx_mcast"`
	RxPackets     int    `json:"rx_packets"`
	RxRate        int    `json:"rx_rate"`
	RxRetries     int    `json:"rx_retries"`
	Signal        int    `json:"signal"`
	State         int    `json:"state"`
	StateHt       bool   `json:"state_ht"`
	StatePwrmgt   bool   `json:"state_pwrmgt"`
	TxBytes       int64  `json:"tx_bytes"`
	TxBytesR      int    `json:"tx_bytes-r"`
	TxPackets     int    `json:"tx_packets"`
	TxPower       int    `json:"tx_power"`
	TxRate        int    `json:"tx_rate"`
	TxRetries     int    `json:"tx_retries"`
	Uptime        int    `json:"uptime"`
	VlanID        int    `json:"vlan_id"`
}

type EthernetTable struct {
	Mac     string `json:"mac"`
	Name    string `json:"name"`
	NumPort int    `json:"num_port"`
}

type RadioNa struct {
	AntennaGain    int    `json:"antenna_gain"`
	BuiltinAntGain int    `json:"builtin_ant_gain"`
	BuiltinAntenna bool   `json:"builtin_antenna"`
	Channel        string `json:"channel"`
	HasDfs         bool   `json:"has_dfs"`
	HasFccdfs      bool   `json:"has_fccdfs"`
	Ht             string `json:"ht"`
	Is11Ac         bool   `json:"is_11ac"`
	MaxTxpower     int    `json:"max_txpower"`
	MinRssi        int    `json:"min_rssi"`
	MinRssiEnabled bool   `json:"min_rssi_enabled"`
	MinTxpower     int    `json:"min_txpower"`
	Name           string `json:"name"`
	Nss            int    `json:"nss"`
	Radio          string `json:"radio"`
	TxPowerMode    string `json:"tx_power_mode"`
}

type RadioNg struct {
	AntennaGain    int    `json:"antenna_gain"`
	BuiltinAntGain int    `json:"builtin_ant_gain"`
	BuiltinAntenna bool   `json:"builtin_antenna"`
	Channel        int    `json:"channel"`
	Ht             string `json:"ht"`
	MaxTxpower     int    `json:"max_txpower"`
	MinRssi        int    `json:"min_rssi"`
	MinRssiEnabled bool   `json:"min_rssi_enabled"`
	MinTxpower     int    `json:"min_txpower"`
	Name           string `json:"name"`
	Nss            int    `json:"nss"`
	Radio          string `json:"radio"`
	TxPowerMode    string `json:"tx_power_mode"`
}

type RadioTable struct {
	AntennaGain    int    `json:"antenna_gain"`
	BuiltinAntGain int    `json:"builtin_ant_gain"`
	BuiltinAntenna bool   `json:"builtin_antenna"`
	Channel        int    `json:"channel"`
	Ht             string `json:"ht"`
	MaxTxpower     int    `json:"max_txpower"`
	MinRssi        int    `json:"min_rssi"`
	MinRssiEnabled bool   `json:"min_rssi_enabled"`
	MinTxpower     int    `json:"min_txpower"`
	Name           string `json:"name"`
	Nss            int    `json:"nss"`
	Radio          string `json:"radio"`
	TxPowerMode    string `json:"tx_power_mode"`
	HasDfs         bool   `json:"has_dfs,omitempty"`
	HasFccdfs      bool   `json:"has_fccdfs,omitempty"`
	Is11Ac         bool   `json:"is_11ac,omitempty"`
}

type Stat struct {
	Bytes            int64  `json:"bytes"`
	GuestNaRxBytes   int    `json:"guest-na-rx_bytes"`
	GuestNaRxPackets int    `json:"guest-na-rx_packets"`
	GuestNaTxBytes   int    `json:"guest-na-tx_bytes"`
	GuestNaTxDropped int    `json:"guest-na-tx_dropped"`
	GuestNaTxErrors  int    `json:"guest-na-tx_errors"`
	GuestNaTxPackets int    `json:"guest-na-tx_packets"`
	GuestNgRxBytes   int    `json:"guest-ng-rx_bytes"`
	GuestNgRxPackets int    `json:"guest-ng-rx_packets"`
	GuestNgTxBytes   int64  `json:"guest-ng-tx_bytes"`
	GuestNgTxDropped int    `json:"guest-ng-tx_dropped"`
	GuestNgTxPackets int    `json:"guest-ng-tx_packets"`
	GuestNgTxRetries int    `json:"guest-ng-tx_retries"`
	GuestRxBytes     int    `json:"guest-rx_bytes"`
	GuestRxPackets   int    `json:"guest-rx_packets"`
	GuestTxBytes     int64  `json:"guest-tx_bytes"`
	GuestTxDropped   int    `json:"guest-tx_dropped"`
	GuestTxErrors    int    `json:"guest-tx_errors"`
	GuestTxPackets   int    `json:"guest-tx_packets"`
	GuestTxRetries   int    `json:"guest-tx_retries"`
	Mac              string `json:"mac"`
	NaRxBytes        int    `json:"na-rx_bytes"`
	NaRxCrypts       int    `json:"na-rx_crypts"`
	NaRxDropped      int    `json:"na-rx_dropped"`
	NaRxErrors       int    `json:"na-rx_errors"`
	NaRxPackets      int    `json:"na-rx_packets"`
	NaTxBytes        int    `json:"na-tx_bytes"`
	NaTxDropped      int    `json:"na-tx_dropped"`
	NaTxErrors       int    `json:"na-tx_errors"`
	NaTxPackets      int    `json:"na-tx_packets"`
	NgRxBytes        int    `json:"ng-rx_bytes"`
	NgRxCrypts       int    `json:"ng-rx_crypts"`
	NgRxDropped      int    `json:"ng-rx_dropped"`
	NgRxErrors       int    `json:"ng-rx_errors"`
	NgRxPackets      int    `json:"ng-rx_packets"`
	NgTxBytes        int64  `json:"ng-tx_bytes"`
	NgTxDropped      int    `json:"ng-tx_dropped"`
	NgTxPackets      int    `json:"ng-tx_packets"`
	NgTxRetries      int    `json:"ng-tx_retries"`
	RxBytes          int    `json:"rx_bytes"`
	RxCrypts         int    `json:"rx_crypts"`
	RxDropped        int    `json:"rx_dropped"`
	RxErrors         int    `json:"rx_errors"`
	RxPackets        int    `json:"rx_packets"`
	TxBytes          int64  `json:"tx_bytes"`
	TxDropped        int    `json:"tx_dropped"`
	TxErrors         int    `json:"tx_errors"`
	TxPackets        int    `json:"tx_packets"`
	TxRetries        int    `json:"tx_retries"`
	UplinkRxBytes    int64  `json:"uplink-rx_bytes"`
	UplinkRxDropped  int    `json:"uplink-rx_dropped"`
	UplinkRxPackets  int    `json:"uplink-rx_packets"`
	UplinkTxBytes    int    `json:"uplink-tx_bytes"`
	UplinkTxPackets  int    `json:"uplink-tx_packets"`
	UserNaRxBytes    int    `json:"user-na-rx_bytes"`
	UserNaRxCrypts   int    `json:"user-na-rx_crypts"`
	UserNaRxDropped  int    `json:"user-na-rx_dropped"`
	UserNaRxErrors   int    `json:"user-na-rx_errors"`
	UserNaRxPackets  int    `json:"user-na-rx_packets"`
	UserNaTxBytes    int    `json:"user-na-tx_bytes"`
	UserNaTxDropped  int    `json:"user-na-tx_dropped"`
	UserNaTxErrors   int    `json:"user-na-tx_errors"`
	UserNaTxPackets  int    `json:"user-na-tx_packets"`
	UserNgRxBytes    int    `json:"user-ng-rx_bytes"`
	UserNgRxCrypts   int    `json:"user-ng-rx_crypts"`
	UserNgRxDropped  int    `json:"user-ng-rx_dropped"`
	UserNgRxErrors   int    `json:"user-ng-rx_errors"`
	UserNgRxPackets  int    `json:"user-ng-rx_packets"`
	UserNgTxBytes    int    `json:"user-ng-tx_bytes"`
	UserNgTxDropped  int    `json:"user-ng-tx_dropped"`
	UserNgTxPackets  int    `json:"user-ng-tx_packets"`
	UserNgTxRetries  int    `json:"user-ng-tx_retries"`
	UserRxBytes      int    `json:"user-rx_bytes"`
	UserRxCrypts     int    `json:"user-rx_crypts"`
	UserRxDropped    int    `json:"user-rx_dropped"`
	UserRxErrors     int    `json:"user-rx_errors"`
	UserRxPackets    int    `json:"user-rx_packets"`
	UserTxBytes      int64  `json:"user-tx_bytes"`
	UserTxDropped    int    `json:"user-tx_dropped"`
	UserTxErrors     int    `json:"user-tx_errors"`
	UserTxPackets    int    `json:"user-tx_packets"`
	UserTxRetries    int    `json:"user-tx_retries"`
}

type SysStats struct {
	Loadavg1  string `json:"loadavg_1"`
	Loadavg15 string `json:"loadavg_15"`
	Loadavg5  string `json:"loadavg_5"`
	MemBuffer int    `json:"mem_buffer"`
	MemTotal  int    `json:"mem_total"`
	MemUsed   int    `json:"mem_used"`
}

type Uplink struct {
	FullDuplex  bool   `json:"full_duplex"`
	IP          string `json:"ip"`
	Mac         string `json:"mac"`
	MaxSpeed    int    `json:"max_speed"`
	Name        string `json:"name"`
	Netmask     string `json:"netmask"`
	NumPort     int    `json:"num_port"`
	RxBytes     int    `json:"rx_bytes"`
	RxBytesR    int    `json:"rx_bytes-r"`
	RxDropped   int    `json:"rx_dropped"`
	RxErrors    int    `json:"rx_errors"`
	RxMulticast int    `json:"rx_multicast"`
	RxPackets   int    `json:"rx_packets"`
	Speed       int    `json:"speed"`
	TxBytes     int    `json:"tx_bytes"`
	TxBytesR    int    `json:"tx_bytes-r"`
	TxDropped   int    `json:"tx_dropped"`
	TxErrors    int    `json:"tx_errors"`
	TxPackets   int    `json:"tx_packets"`
	Type        string `json:"type"`
	Up          bool   `json:"up"`
}

type VapTable struct {
	ApMac      string `json:"ap_mac"`
	Bssid      string `json:"bssid"`
	Ccq        int    `json:"ccq"`
	Channel    int    `json:"channel"`
	Essid      string `json:"essid"`
	ID         string `json:"id"`
	IsGuest    bool   `json:"is_guest"`
	IsWep      bool   `json:"is_wep"`
	MapID      string `json:"map_id"`
	Name       string `json:"name"`
	NumSta     int    `json:"num_sta"`
	Radio      string `json:"radio"`
	RxBytes    int    `json:"rx_bytes"`
	RxCrypts   int    `json:"rx_crypts"`
	RxDropped  int    `json:"rx_dropped"`
	RxErrors   int    `json:"rx_errors"`
	RxFrags    int    `json:"rx_frags"`
	RxNwids    int    `json:"rx_nwids"`
	RxPackets  int    `json:"rx_packets"`
	SiteID     string `json:"site_id"`
	State      string `json:"state"`
	T          string `json:"t"`
	TxBytes    int    `json:"tx_bytes"`
	TxDropped  int    `json:"tx_dropped"`
	TxErrors   int    `json:"tx_errors"`
	TxPackets  int    `json:"tx_packets"`
	TxPower    int    `json:"tx_power"`
	TxRetries  int    `json:"tx_retries"`
	Up         bool   `json:"up"`
	Usage      string `json:"usage"`
	WlanconfID string `json:"wlanconf_id"`
	Extchannel int    `json:"extchannel,omitempty"`
}

type UplinkTable struct {
	ApConnected   bool   `json:"ap_connected"`
	ApMac         string `json:"ap_mac"`
	AuthTime      int64  `json:"auth_time"`
	Authorized    bool   `json:"authorized"`
	Ccq           int    `json:"ccq"`
	Channel       int    `json:"channel"`
	Configured    bool   `json:"configured"`
	DhcpendTime   int    `json:"dhcpend_time"`
	DhcpstartTime int    `json:"dhcpstart_time"`
	Hostname      string `json:"hostname"`
	Idletime      int    `json:"idletime"`
	Is11A         bool   `json:"is_11a"`
	Is11Ac        bool   `json:"is_11ac"`
	Is11N         bool   `json:"is_11n"`
	Mac           string `json:"mac"`
	Name          string `json:"name"`
	Noise         int    `json:"noise"`
	Radio         string `json:"radio"`
	Rssi          int    `json:"rssi"`
	RxBytes       int64  `json:"rx_bytes"`
	RxBytesR      int    `json:"rx_bytes-r"`
	RxMcast       int    `json:"rx_mcast"`
	RxPackets     int    `json:"rx_packets"`
	RxRate        int    `json:"rx_rate"`
	RxRetries     int    `json:"rx_retries"`
	Signal        int    `json:"signal"`
	State         int    `json:"state"`
	StateHt       bool   `json:"state_ht"`
	StatePwrmgt   bool   `json:"state_pwrmgt"`
	TxBytes       int64  `json:"tx_bytes"`
	TxBytesR      int    `json:"tx_bytes-r"`
	TxPackets     int    `json:"tx_packets"`
	TxPower       int    `json:"tx_power"`
	TxRate        int    `json:"tx_rate"`
	TxRetries     int    `json:"tx_retries"`
	Type          string `json:"type"`
	Up            bool   `json:"up"`
	UplinkMac     string `json:"uplink_mac"`
	Uptime        int    `json:"uptime"`
	VlanID        int    `json:"vlan_id"`
}

type VwireTable struct {
	ApMac string `json:"ap_mac"`
	Radio string `json:"radio"`
}
