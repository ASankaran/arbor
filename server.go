/**
* Copyright © 2017, ACM@UIUC
*
* This file is part of the Groot Project.
*
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package arbor

import (
	"fmt"
	"os"

	"github.com/acm-uiuc/arbor/logger"
	"github.com/acm-uiuc/arbor/security"
	"github.com/acm-uiuc/arbor/server"
)

const help = `Usage: executable [-r | --register-client client_name] [-c | --check-registration token] [-u | --unsecured]
                   -r | --register-client client_name -> registers a client, generates a token
                   -c | --check-registration token    -> checks if a token is valid and returns name of client
                   -u | --unsecured                   -> runs arbor without the security layer
                   without args                       -> runs arbor with the security layer	`

// Boot is a standard server CLI
//
// Provide a set of routes to serve and a port to serve on.
//
// Usage: executable [-r | --register-client client_name] [-c | --check-registration token] [-u | --unsecured]
//
// 	-r | --register-client client_name
// registers a client, generates a token
//
// 	-c | --check-registration token
// checks if a token is valid and returns name of client
//
// 	-u | --unsecured
// runs arbor without the security layer
//
// 	without args
// runs arbor with the security layer
//
// It will start the arbor instance, parsing the command arguments and execute the behavior.
func Boot(routes RouteCollection, port uint16) *server.ArborServer {
	var srv *server.ArborServer
	if len(os.Args) == 3 && (os.Args[1] == "--register-client" || os.Args[1] == "-r") {
		RegisterClient(os.Args[2])
	} else if len(os.Args) == 3 && (os.Args[1] == "--check-registration" || os.Args[1] == "-c") {
		CheckRegistration(os.Args[2])
	} else if len(os.Args) == 2 && (os.Args[1] == "--unsecured" || os.Args[1] == "-u") {
		logger.Log(logger.WARN, "Starting Arbor in unsecured mode")
		srv = server.StartUnsecuredServer(routes.toServiceRoutes(), port)
	} else if len(os.Args) == 2 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
		fmt.Println(help)
	} else if len(os.Args) > 1 {
		logger.Log(logger.ERR, "Unknown Command")
		fmt.Println(help)
	} else {
		srv = server.StartSecuredServer(routes.toServiceRoutes(), port)
	}
	return srv
}

// RegisterClient will generate a access token for a client
//
// Currently uses a db of client names.
func RegisterClient(name string) {
	security.Init()
	token, err := security.AddClient(name)
	if err != nil {
		fmt.Println(err)
		return
	}
	logger.Log(logger.SPEC, "Client "+name+" has been granted authorization token: "+token)
	defer security.Shutdown()
}

// CheckRegistration allows you to check what client was assigned to a particular token
func CheckRegistration(token string) {
	security.Init()
	fmt.Println(security.IsAuthorizedClient(token))
	defer security.Shutdown()
}
