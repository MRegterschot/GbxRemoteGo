package structs

type TMSystemInfo struct {
	PublishedIp            string `json:"PublishedIp"`
	Port                   int    `json:"Port"`
	P2PPort                int    `json:"P2PPort"`
	TitleId                string `json:"TitleId"`
	ServerLogin            string `json:"ServerLogin"`
	ServerPlayerId         int    `json:"ServerPlayerId"`
	ConnectionDownloadRate int    `json:"ConnectionDownloadRate"`
	ConnectionUploadRate   int    `json:"ConnectionUploadRate"`
	IsServer               bool   `json:"IsServer"`
	IsDedicated            bool   `json:"IsDedicated"`
}

type TMStatus struct {
	Code int    `json:"Code"`
	Name string `json:"Name"`
}

type TMVersion struct {
	Name       string `json:"Name"`
	TitleId    string `json:"TitleId"`
	Version    string `json:"Version"`
	Build      string `json:"Build"`
	ApiVersion string `json:"ApiVersion"`
}

type TMServerTag struct {
	Name  string `json:"Name" xmlrpc:"Name"`
	Value string `json:"Value" xmlrpc:"Value"`
}

type TMServerVisibility int

const (
	TMServerVisible TMServerVisibility = iota
	TMServerHidden
	TMServerHiddenFromNations
)

type TMMaxPlayers struct {
	CurrentValue int `json:"CurrentValue" xmlrpc:"CurrentValue"`
	NextValue    int `json:"NextValue" xmlrpc:"NextValue"`
}

type TMMaxSpectators struct {
	CurrentValue int `json:"CurrentValue" xmlrpc:"CurrentValue"`
	NextValue    int `json:"NextValue" xmlrpc:"NextValue"`
}

type TMServerOptions struct {
	Name                 string  `json:"Name,omitempty" xmlrpc:"Name"`
	Comment              string  `json:"Comment,omitempty" xmlrpc:"Comment"`
	Password             string  `json:"Password,omitempty" xmlrpc:"Password"`
	PasswordForSpectator string  `json:"PasswordForSpectator,omitempty" xmlrpc:"PasswordForSpectator"`
	NextCallVoteTimeOut  int     `json:"NextCallVoteTimeOut,omitempty" xmlrpc:"NextCallVoteTimeOut"`
	CallVoteRatio        float32 `json:"CallVoteRatio,omitempty" xmlrpc:"CallVoteRatio"`

	// Optional fields
	HideServer                                 int  `json:"HideServer,omitempty" xmlrpc:"HideServer"`
	CurrentMaxPlayers                          int  `json:"CurrentMaxPlayers,omitempty" xmlrpc:"CurrentMaxPlayers"`
	NextMaxPlayers                             int  `json:"NextMaxPlayers,omitempty" xmlrpc:"NextMaxPlayers"`
	CurrentMaxSpecators                        int  `json:"CurrentMaxSpecators,omitempty" xmlrpc:"CurrentMaxSpecators"`
	NextMaxSpectators                          int  `json:"NextMaxSpectators,omitempty" xmlrpc:"NextMaxSpectators"`
	KeepPlayerSlots                            bool `json:"KeepPlayerSlots,omitempty" xmlrpc:"KeepPlayerSlots"`
	CurrentCallVoteTimeOut                     int  `json:"CurrentCallVoteTimeOut,omitempty" xmlrpc:"CurrentCallVoteTimeOut"`
	AllowMapDownload                           bool `json:"AllowMapDownload,omitempty" xmlrpc:"AllowMapDownload"`
	AutoSaveReplays                            bool `json:"AutoSaveReplays,omitempty" xmlrpc:"AutoSaveReplays"`
	ClientInputsMaxLatency                     int  `json:"ClientInputsMaxLatency,omitempty" xmlrpc:"ClientInputsMaxLatency"`
	DisableHorns                               bool `json:"DisableHorns,omitempty" xmlrpc:"DisableHorns"`
	DisableServiceAnnounces                    bool `json:"DisableServiceAnnounces,omitempty" xmlrpc:"DisableServiceAnnounces"`
	PacketAssembly_PacketsPerFrame             int  `json:"PacketAssembly_PacketsPerFrame,omitempty" xmlrpc:"PacketAssembly_PacketsPerFrame"`
	PacketAssembly_FullPacketsPerFrame         int  `json:"PacketAssembly_FullPacketsPerFrame,omitempty" xmlrpc:"PacketAssembly_FullPacketsPerFrame"`
	TrustClientSimu_ClientToServer_SendingRate int  `json:"TrustClientSimu_ClientToServer_SendingRate,omitempty" xmlrpc:"TrustClientSimu_ClientToServer_SendingRate"`
	DelayedVisuals_ServerToClient_SendingRate  int  `json:"DelayedVisuals_ServerToClient_SendingRate,omitempty" xmlrpc:"DelayedVisuals_ServerToClient_SendingRate"`
	DelayedVisuals_RestrictToSpectators        bool `json:"DelayedVisuals_RestrictToSpectators,omitempty" xmlrpc:"DelayedVisuals_RestrictToSpectators"`
}

type TMServerOptionsRequest struct {
	Name                 string  `json:"Name,omitempty" xmlrpc:"Name"`
	Comment              string  `json:"Comment,omitempty" xmlrpc:"Comment"`
	Password             string  `json:"Password,omitempty" xmlrpc:"Password"`
	PasswordForSpectator string  `json:"PasswordForSpectator,omitempty" xmlrpc:"PasswordForSpectator"`
	NextCallVoteTimeOut  int     `json:"NextCallVoteTimeOut" xmlrpc:"NextCallVoteTimeOut"`
	CallVoteRatio        float32 `json:"CallVoteRatio" xmlrpc:"CallVoteRatio"`

	// Optional fields now use pointers
	HideServer                                 *int  `json:"HideServer,omitempty" xmlrpc:"HideServer"`
	CurrentMaxPlayers                          *int  `json:"CurrentMaxPlayers,omitempty" xmlrpc:"CurrentMaxPlayers"`
	NextMaxPlayers                             *int  `json:"NextMaxPlayers,omitempty" xmlrpc:"NextMaxPlayers"`
	CurrentMaxSpecators                        *int  `json:"CurrentMaxSpecators,omitempty" xmlrpc:"CurrentMaxSpecators"`
	NextMaxSpectators                          *int  `json:"NextMaxSpectators,omitempty" xmlrpc:"NextMaxSpectators"`
	KeepPlayerSlots                            *bool `json:"KeepPlayerSlots,omitempty" xmlrpc:"KeepPlayerSlots"`
	CurrentCallVoteTimeOut                     *int  `json:"CurrentCallVoteTimeOut,omitempty" xmlrpc:"CurrentCallVoteTimeOut"`
	AllowMapDownload                           *bool `json:"AllowMapDownload,omitempty" xmlrpc:"AllowMapDownload"`
	AutoSaveReplays                            *bool `json:"AutoSaveReplays,omitempty" xmlrpc:"AutoSaveReplays"`
	ClientInputsMaxLatency                     *int  `json:"ClientInputsMaxLatency,omitempty" xmlrpc:"ClientInputsMaxLatency"`
	DisableHorns                               *bool `json:"DisableHorns,omitempty" xmlrpc:"DisableHorns"`
	DisableServiceAnnounces                    *bool `json:"DisableServiceAnnounces,omitempty" xmlrpc:"DisableServiceAnnounces"`
	PacketAssembly_PacketsPerFrame             *int  `json:"PacketAssembly_PacketsPerFrame,omitempty" xmlrpc:"PacketAssembly_PacketsPerFrame"`
	PacketAssembly_FullPacketsPerFrame         *int  `json:"PacketAssembly_FullPacketsPerFrame,omitempty" xmlrpc:"PacketAssembly_FullPacketsPerFrame"`
	TrustClientSimu_ClientToServer_SendingRate *int  `json:"TrustClientSimu_ClientToServer_SendingRate,omitempty" xmlrpc:"TrustClientSimu_ClientToServer_SendingRate"`
	DelayedVisuals_ServerToClient_SendingRate  *int  `json:"DelayedVisuals_ServerToClient_SendingRate,omitempty" xmlrpc:"DelayedVisuals_ServerToClient_SendingRate"`
	DelayedVisuals_RestrictToSpectators        *bool `json:"DelayedVisuals_RestrictToSpectators,omitempty" xmlrpc:"DelayedVisuals_RestrictToSpectators"`
}
