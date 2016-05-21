// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package discordapp provides constants for using OAuth2 to access Discord.
package discordapp

import (
	"golang.org/x/oauth2"
)

// Endpoint is Discord's OAuth 2.0 endpoint.
var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://discordapp.com/api/oauth2/authorize",
	TokenURL: "https://discordapp.com/api/oauth2/token",
}
