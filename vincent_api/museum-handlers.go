package vincent_api

import (

    // "github.com/aws/aws-sdk-go/aws/session"
    "github.com/tmaiaroto/aegis/lambda"
    // "github.com/aws/aws-sdk-go/aws"
    // "github.com/guregu/dynamo"
    
    "encoding/json"
    "net/url"
)

// PostMuseum
func PostMuseum(ctx *lambda.Context, evt *lambda.Event, res *lambda.ProxyResponse, params url.Values) {
    
    museumId := params.Get("museum")
    
    //db := dynamo.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
    //table := db.Table("testTable")
    
    
    ret := struct {
            MuseumId  string  `json:"museum_id"`
        }{museumId}
    
    retJson, _ := json.Marshal(ret)
	res.Body = string(retJson)
    
	res.Headers = map[string]string{"Content-Type": "application/json"}
    
}


// GetMuseum
func GetMuseum(ctx *lambda.Context, evt *lambda.Event, res *lambda.ProxyResponse, params url.Values) {
    
    museumId := params.Get("museum")
    
    ret := struct {
            MuseumId  string  `json:"museum_id"`
        }{museumId}
    
    retJson, _ := json.Marshal(ret)
	res.Body = string(retJson)
    
	res.Headers = map[string]string{"Content-Type": "application/json"}
    
}