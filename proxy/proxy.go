/**
* Copyright © 2017, ACM@UIUC
*
* This file is part of the Groot Project.  
* 
* The Groot Project is open source software, released under the University of
* Illinois/NCSA Open Source License. You should have received a copy of
* this license in a file with the distribution.
**/

package proxy

import (
	"github.com/acm-uiuc/groot/config"
)

var JSONHeader string = "application/json; charset=UTF-8"
var TEXTHeader string = "text/plain; charset=utf-8"
var HTMLHeader string = "text/html;charset=utf-8"
var XMLHeader  string = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>"
var AccessControlPolicy = config.AccessControlPolicy
