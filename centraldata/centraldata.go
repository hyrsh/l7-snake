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
	selfData.Terminator = configstruct.CurrentConfig.Data.Routing.Terminator
	selfData.LastUpdated = time.Now().Format("2006-01-02 15:04:05")
	selfData.Health = 1                                                                  //init value with "Offline"
	selfData.Targets = int32(len(configstruct.CurrentConfig.Data.Communication.Targets)) //the value never changes after start so we can use the target counter

	//append routes to empty data
	for i := 0; i < len(configstruct.CurrentConfig.Data.Routing.Routes); i++ {
		selfData.Routes = append(selfData.Routes, configstruct.CurrentConfig.Data.Routing.Routes[i])
	}

	return &selfData
}
