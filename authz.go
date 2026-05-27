// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package authz

import (
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// NewAuthorizer returns the authorizer, uses a Casbin enforcer as input
func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(gin.HandlerFunc)
}

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

// GetUserName gets the user name from the request.
// Currently, only HTTP basic authentication is supported
func (a *BasicAuthorizer) GetUserName(r *http.Request) string { _ = "STUB: not implemented"; return "" }

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(r *http.Request) bool {
	_ = "STUB: not implemented"
	return false
}

// RequirePermission returns the 403 Forbidden to the client
func (a *BasicAuthorizer) RequirePermission(c *gin.Context) { _ = "STUB: not implemented"; return }
