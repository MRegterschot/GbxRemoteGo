package structs

import "strings"

type TMForceShowAllOpponents struct {
	CurrentValue int `json:"CurrentValue" xmlrpc:"CurrentValue"`
	NextValue    int `json:"NextValue" xmlrpc:"NextValue"`
}

type TMScriptName struct {
	CurrentValue string `json:"CurrentValue" xmlrpc:"CurrentValue"`
	NextValue    string `json:"NextValue" xmlrpc:"NextValue"`
}

// TMScriptMode represents a type for TrackMania Script Modes.
type TMScriptMode string

// Define all possible script modes.
const (
	ScriptTimeAttack TMScriptMode = "Trackmania/TM_TimeAttack_Online.Script.txt"
	ScriptLaps       TMScriptMode = "Trackmania/TM_Laps_Online.Script.txt"
	ScriptRounds     TMScriptMode = "Trackmania/TM_Rounds_Online.Script.txt"
	ScriptCup        TMScriptMode = "Trackmania/TM_Cup_Online.Script.txt"
	ScriptTeams      TMScriptMode = "Trackmania/TM_Teams_Online.Script.txt"
	ScriptKnockout   TMScriptMode = "Trackmania/TM_Knockout_Online.Script.txt"
	ScriptChampion   TMScriptMode = "Trackmania/Deprecated/TM_Champion_Online.Script.txt"
	ScriptRoyal      TMScriptMode = "Trackmania/TM_RoyalTimeAttack_Online.Script.txt"
	ScriptStunt      TMScriptMode = "Trackmania/TM_StuntMulti_Online.Script.txt"
	ScriptPlatform   TMScriptMode = "Trackmania/TM_Platform_Online.Script.txt"
	ScriptTMWC       TMScriptMode = "TrackMania/TM_TMWC2023_Online.Script.txt"
	ScriptTMWT       TMScriptMode = "TrackMania/TM_TMWTTeams_Online.Script.txt"
)

// Map of user input to script filenames
var scriptMap = map[string]TMScriptMode{
	"ta":         ScriptTimeAttack,
	"timeattack": ScriptTimeAttack,
	"laps":       ScriptLaps,
	"rounds":     ScriptRounds,
	"cup":        ScriptCup,
	"team":       ScriptTeams,
	"teams":      ScriptTeams,
	"ko":         ScriptKnockout,
	"knockout":   ScriptKnockout,
	"champion":   ScriptChampion,
	"royal":      ScriptRoyal,
	"stunt":      ScriptStunt,
	"platform":   ScriptPlatform,
	"tmwc":       ScriptTMWC,
	"tmwt":       ScriptTMWT,
}

// Convert user input into a valid script mode (or return empty if not found)
func GetScriptByName(userInput string) string {
	lowerInput := strings.ToLower(userInput)
	if script, exists := scriptMap[lowerInput]; exists {
		return string(script)
	}
	return userInput
}
