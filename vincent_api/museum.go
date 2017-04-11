package vincent_api

import (
    "time"
)

type MuseumConfig struct {
    MuseumId        int             `json:"museum_id" dynamo:"museum_id"`
    MuseumName      string          `json:"museum_name" dynamo:"museum_name"`
    Category        string          `json:"category" dynamo:"category"`
    Description     string          `json:"description" dynamo:"description"`            
    Created         struct{
                        At          time.Time       `json:"at" bson:"at"`
                        By          int   `json:"by" bson:"by"`
                    } `json:"created" bson:"created"`
                    
    Updated         []UpdatedStruct `json:"updated" bson:"updated"`

}


type SummaryMuseumConfig struct {
    Id              int             `json:"id"`
    Facility        int             `json:"facility"`
    Category        string          `json:"category"`
    Name            string          `json:"name"`
    Description     string          `json:"description"`   
    Recommendation  string          `json:"recommendation"`
    LastCount       float32         `json:"last_count"`
    LastRunDate     time.Time       `json:"last_run_date"`
    AlarmCategory   string          `json:"alarm_category"`
    Scheduled       bool            `json:"scheduled"`
}



