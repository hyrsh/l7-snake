package centraldata

import (
	"l7-snake/configstruct"
	"l7-snake/pt3status"
	"time"
)

//we store our runtime data in here
var CentralData []*pt3status.Status_Chain

//returns self-information in format to append to CentralData
func OwnData() *pt3status.Status_Chain {
	//set own data
	var selfData pt3status.Status_Chain
	selfData.Id = configstruct.CurrentConfig.Data.Communication.Id
	selfData.Terminator = Bool2String(configstruct.CurrentConfig.Data.Routing.Terminator)
	selfData.LastUpdated = time.Now().Format("2006-01-02 15:04:05")
	selfData.Health = "INIT"                                                             //init value
	selfData.Targets = int32(len(configstruct.CurrentConfig.Data.Communication.Targets)) //the value never changes after start so we can use the target counter

	//append routes to empty data
	for i := 0; i < len(configstruct.CurrentConfig.Data.Routing.Routes); i++ {
		selfData.Routes = append(selfData.Routes, configstruct.CurrentConfig.Data.Routing.Routes[i])
	}

	return &selfData
}

//I needed this, don't judge
func Bool2String(b bool) string {
	if b {
		return "yes"
	} else {
		return "no"
	}
}
