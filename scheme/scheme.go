// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// go-aah/security source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package scheme

import (
	"aahframework.org/ahttp.v0"
	"aahframework.org/config.v0"
	"aahframework.org/security.v0-unstable/authc"
	"aahframework.org/security.v0-unstable/authz"
)

type (
	// Schemer interface is implemented by aah framework's authentication scheme.
	Schemer interface {
		Init(appCfg *config.Config, keyName string) error
		Scheme() string
		SetAuthenticator(authenticator authc.Authenticator) error
		SetAuthorizer(authenticator authz.Authorizer) error
		DoAuthenticate(authcToken *authc.AuthenticationToken) (*authc.AuthenticationInfo, error)
		DoAuthorizationInfo(authcInfo *authc.AuthenticationInfo) *authz.AuthorizationInfo
		ExtractAuthenticationToken(r *ahttp.Request) *authc.AuthenticationToken
	}
)