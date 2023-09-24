// Copyright (C) 2014 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at https://mozilla.org/MPL/2.0/.

package api

import (
	"testing"

	"github.com/syncthing/syncthing/lib/config"
)

var guiCfg config.GUIConfiguration

func init() {
	guiCfg.User = "user"
	guiCfg.SetPassword("pass")
}

func TestStaticAuthOK(t *testing.T) {
	t.Parallel()

	ok := authStatic("user", "pass", guiCfg)
	if !ok {
		t.Fatalf("should pass auth")
	}
}

func TestSimpleAuthUsernameFail(t *testing.T) {
	t.Parallel()

	ok := authStatic("userWRONG", "pass", guiCfg)
	if ok {
		t.Fatalf("should fail auth")
	}
}

func TestStaticAuthPasswordFail(t *testing.T) {
	t.Parallel()

	ok := authStatic("user", "passWRONG", guiCfg)
	if ok {
		t.Fatalf("should fail auth")
	}
}

func TestAuthLDAPSendsCorrectBindDNWithTemplate(t *testing.T) {
	t.Parallel()

	templatedDn := ldapTemplateBindDN("cn=%s,dc=some,dc=example,dc=com", "username")
	expectedDn := "cn=username,dc=some,dc=example,dc=com"
	if expectedDn != templatedDn {
		t.Fatalf("ldapTemplateBindDN should be %s != %s", expectedDn, templatedDn)
	}
}

func TestAuthLDAPSendsCorrectBindDNWithNoTemplate(t *testing.T) {
	t.Parallel()

	templatedDn := ldapTemplateBindDN("cn=fixedusername,dc=some,dc=example,dc=com", "username")
	expectedDn := "cn=fixedusername,dc=some,dc=example,dc=com"
	if expectedDn != templatedDn {
		t.Fatalf("ldapTemplateBindDN should be %s != %s", expectedDn, templatedDn)
	}
}
