package main

import (
    // "os"
    "github.com/jcgarciaram/vincent/vincent_api"
    api "github.com/jcgarciaram/general-api"
    "github.com/jcgarciaram/general-api/routes"
)

var (
    
    //Get MySQL and Mongo configuration from environment variables
    // clientdb_server         = os.Getenv("MYSQL_SERVER")
    // clientdb_port           = os.Getenv("MYSQL_PORT")
    // clientdb_user           = os.Getenv("MYSQL_USER")
    // clientdb_pass           = os.Getenv("MYSQL_PASS")
    // clientdb_schema_db      = os.Getenv("MYSQL_DB")
    // clientdb_dbtype         = "mysql"
    
    // mongoConnStr = os.Getenv("MONGO_CONNECT_STRING")
    
    // secret = os.Getenv("AUTH0_CLIENT_SECRET")
    
    
    //NILESH specific variables
    clientdb_server         = "devdb1.cxw6llilc50s.ap-south-1.rds.amazonaws.com"
    clientdb_port           = "3306"
    clientdb_user           = "agilliticsdev1"
    clientdb_pass           = "agilliticsdev1"
    clientdb_schema_db      = ""
    clientdb_dbtype         = "mysql"

    mongoConnStr = "mongodb://localhost/wmsight-devdb-copy"
    
    secret = "empiricalfriedtreespracticalcarnies"
    
)


func main() {
    
    r := routes.Routes{}
    
    // Append general_museum routes
    r.AppendRoutes(vincent_api.GetRoutes())
    

    router := api.NewRouter(r)

    router.Listen()
    // router.Gateway()
}