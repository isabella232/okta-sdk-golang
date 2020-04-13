/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import (
	"fmt"
	"github.com/okta/okta-sdk-golang/okta.v2/query"
	"time"
)

type UserResource resource

type User struct {
	Embedded              interface{}      `json:"_embedded,omitempty"`
	Links                 interface{}      `json:"_links,omitempty"`
	Activated             *time.Time       `json:"activated,omitempty"`
	Created               *time.Time       `json:"created,omitempty"`
	Credentials           *UserCredentials `json:"credentials,omitempty"`
	Id                    string           `json:"id,omitempty"`
	LastLogin             *time.Time       `json:"lastLogin,omitempty"`
	LastUpdated           *time.Time       `json:"lastUpdated,omitempty"`
	PasswordChanged       *time.Time       `json:"passwordChanged,omitempty"`
	Profile               *UserProfile     `json:"profile,omitempty"`
	Status                string           `json:"status,omitempty"`
	StatusChanged         *time.Time       `json:"statusChanged,omitempty"`
	TransitioningToStatus string           `json:"transitioningToStatus,omitempty"`
	Type                  *UserType        `json:"type,omitempty"`
}

// Creates a new user in your Okta organization with or without credentials.
func (m *UserResource) CreateUser(body CreateUserRequest, qp *query.Params) (*User, *Response, error) {
	url := fmt.Sprintf("/api/v1/users")
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var user *User

	resp, err := m.client.requestExecutor.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

// Fetches a user from your Okta organization.
func (m *UserResource) GetUser(userId string) (*User, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v", userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var user *User

	resp, err := m.client.requestExecutor.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

// Update a user&#x27;s profile and/or credentials using strict-update semantics.
func (m *UserResource) UpdateUser(userId string, body User, qp *query.Params) (*User, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v", userId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var user *User

	resp, err := m.client.requestExecutor.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

// Deletes a user permanently.  This operation can only be performed on users that have a &#x60;DEPROVISIONED&#x60; status.  **This action cannot be recovered!**
func (m *UserResource) DeactivateOrDeleteUser(userId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v", userId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Lists users in your organization with pagination in most cases.  A subset of users can be returned that match a supported filter expression or search criteria.
func (m *UserResource) ListUsers(qp *query.Params) ([]*User, *Response, error) {
	url := fmt.Sprintf("/api/v1/users")
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var user []*User

	resp, err := m.client.requestExecutor.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

// Clears Okta sessions for the currently logged in user. By default, the current session remains active. Use this method in a browser-based application.
func (m *UserResource) ClearCurrentUserSessions(qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/me/lifecycle/delete_sessions")
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *UserResource) SetLinkedObjectForUser(associatedUserId string, primaryRelationshipName string, primaryUserId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/linkedObjects/%v/%v", associatedUserId, primaryRelationshipName, primaryUserId)

	req, err := m.client.requestExecutor.WithAccept("").WithContentType("application/json").NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Fetch a user by &#x60;id&#x60;, &#x60;login&#x60;, or &#x60;login shortname&#x60; if the short name is unambiguous.
func (m *UserResource) PartialUpdateUser(userId string, body User, qp *query.Params) (*User, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v", userId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var user *User

	resp, err := m.client.requestExecutor.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

// Fetches appLinks for all direct or indirect (via group membership) assigned applications.
func (m *UserResource) ListAppLinks(userId string) ([]*AppLink, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/appLinks", userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var appLink []*AppLink

	resp, err := m.client.requestExecutor.Do(req, &appLink)
	if err != nil {
		return nil, resp, err
	}

	return appLink, resp, nil
}

// Lists all client resources for which the specified user has grants or tokens.
func (m *UserResource) ListUserClients(userId string) ([]*OAuth2Client, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/clients", userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2Client []*OAuth2Client

	resp, err := m.client.requestExecutor.Do(req, &oAuth2Client)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2Client, resp, nil
}

// Revokes all grants for the specified user and client
func (m *UserResource) RevokeGrantsForUserAndClient(userId string, clientId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/clients/%v/grants", userId, clientId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Lists all grants for a specified user and client
func (m *UserResource) ListGrantsForUserAndClient(userId string, clientId string, qp *query.Params) ([]*OAuth2ScopeConsentGrant, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/clients/%v/grants", userId, clientId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2ScopeConsentGrant []*OAuth2ScopeConsentGrant

	resp, err := m.client.requestExecutor.Do(req, &oAuth2ScopeConsentGrant)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2ScopeConsentGrant, resp, nil
}

// Revokes all refresh tokens issued for the specified User and Client.
func (m *UserResource) RevokeTokensForUserAndClient(userId string, clientId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/clients/%v/tokens", userId, clientId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Lists all refresh tokens issued for the specified User and Client.
func (m *UserResource) ListRefreshTokensForUserAndClient(userId string, clientId string, qp *query.Params) ([]*OAuth2RefreshToken, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/clients/%v/tokens", userId, clientId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2RefreshToken []*OAuth2RefreshToken

	resp, err := m.client.requestExecutor.Do(req, &oAuth2RefreshToken)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2RefreshToken, resp, nil
}

// Revokes the specified refresh token.
func (m *UserResource) RevokeTokenForUserAndClient(userId string, clientId string, tokenId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/clients/%v/tokens/%v", userId, clientId, tokenId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Gets a refresh token issued for the specified User and Client.
func (m *UserResource) GetRefreshTokenForUserAndClient(userId string, clientId string, tokenId string, qp *query.Params) (*OAuth2RefreshToken, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/clients/%v/tokens/%v", userId, clientId, tokenId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2RefreshToken *OAuth2RefreshToken

	resp, err := m.client.requestExecutor.Do(req, &oAuth2RefreshToken)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2RefreshToken, resp, nil
}

// Changes a user&#x27;s password by validating the user&#x27;s current password. This operation can only be performed on users in &#x60;STAGED&#x60;, &#x60;ACTIVE&#x60;, &#x60;PASSWORD_EXPIRED&#x60;, or &#x60;RECOVERY&#x60; status that have a valid password credential
func (m *UserResource) ChangePassword(userId string, body ChangePasswordRequest, qp *query.Params) (*UserCredentials, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/credentials/change_password", userId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var userCredentials *UserCredentials

	resp, err := m.client.requestExecutor.Do(req, &userCredentials)
	if err != nil {
		return nil, resp, err
	}

	return userCredentials, resp, nil
}

// Changes a user&#x27;s recovery question &amp; answer credential by validating the user&#x27;s current password.  This operation can only be performed on users in **STAGED**, **ACTIVE** or **RECOVERY** &#x60;status&#x60; that have a valid password credential
func (m *UserResource) ChangeRecoveryQuestion(userId string, body UserCredentials) (*UserCredentials, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/credentials/change_recovery_question", userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var userCredentials *UserCredentials

	resp, err := m.client.requestExecutor.Do(req, &userCredentials)
	if err != nil {
		return nil, resp, err
	}

	return userCredentials, resp, nil
}

// Generates a one-time token (OTT) that can be used to reset a user&#x27;s password
func (m *UserResource) ForgotPasswordGenerateOneTimeToken(userId string, qp *query.Params) (*ForgotPasswordResponse, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/credentials/forgot_password", userId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var forgotPasswordResponse *ForgotPasswordResponse

	resp, err := m.client.requestExecutor.Do(req, &forgotPasswordResponse)
	if err != nil {
		return nil, resp, err
	}

	return forgotPasswordResponse, resp, nil
}

// Sets a new password for a user by validating the user&#x27;s answer to their current recovery question
func (m *UserResource) ForgotPasswordSetNewPassword(userId string, body UserCredentials, qp *query.Params) (*ForgotPasswordResponse, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/credentials/forgot_password", userId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var forgotPasswordResponse *ForgotPasswordResponse

	resp, err := m.client.requestExecutor.Do(req, &forgotPasswordResponse)
	if err != nil {
		return nil, resp, err
	}

	return forgotPasswordResponse, resp, nil
}

// Revokes all grants for a specified user
func (m *UserResource) RevokeUserGrants(userId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/grants", userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Lists all grants for the specified user
func (m *UserResource) ListUserGrants(userId string, qp *query.Params) ([]*OAuth2ScopeConsentGrant, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/grants", userId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2ScopeConsentGrant []*OAuth2ScopeConsentGrant

	resp, err := m.client.requestExecutor.Do(req, &oAuth2ScopeConsentGrant)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2ScopeConsentGrant, resp, nil
}

// Revokes one grant for a specified user
func (m *UserResource) RevokeUserGrant(userId string, grantId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/grants/%v", userId, grantId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Gets a grant for the specified user
func (m *UserResource) GetUserGrant(userId string, grantId string, qp *query.Params) (*OAuth2ScopeConsentGrant, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/grants/%v", userId, grantId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var oAuth2ScopeConsentGrant *OAuth2ScopeConsentGrant

	resp, err := m.client.requestExecutor.Do(req, &oAuth2ScopeConsentGrant)
	if err != nil {
		return nil, resp, err
	}

	return oAuth2ScopeConsentGrant, resp, nil
}

// Fetches the groups of which the user is a member.
func (m *UserResource) ListUserGroups(userId string) ([]*Group, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/groups", userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var group []*Group

	resp, err := m.client.requestExecutor.Do(req, &group)
	if err != nil {
		return nil, resp, err
	}

	return group, resp, nil
}

// Lists the IdPs associated with the user.
func (m *UserResource) ListUserIdentityProviders(userId string) ([]*IdentityProvider, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/idps", userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var identityProvider []*IdentityProvider

	resp, err := m.client.requestExecutor.Do(req, &identityProvider)
	if err != nil {
		return nil, resp, err
	}

	return identityProvider, resp, nil
}

// Activates a user.  This operation can only be performed on users with a &#x60;STAGED&#x60; status.  Activation of a user is an asynchronous operation.  The user will have the &#x60;transitioningToStatus&#x60; property with a value of &#x60;ACTIVE&#x60; during activation to indicate that the user hasn&#x27;t completed the asynchronous operation.  The user will have a status of &#x60;ACTIVE&#x60; when the activation process is complete.
func (m *UserResource) ActivateUser(userId string, qp *query.Params) (*UserActivationToken, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/lifecycle/activate", userId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var userActivationToken *UserActivationToken

	resp, err := m.client.requestExecutor.Do(req, &userActivationToken)
	if err != nil {
		return nil, resp, err
	}

	return userActivationToken, resp, nil
}

// Deactivates a user.  This operation can only be performed on users that do not have a &#x60;DEPROVISIONED&#x60; status.  Deactivation of a user is an asynchronous operation.  The user will have the &#x60;transitioningToStatus&#x60; property with a value of &#x60;DEPROVISIONED&#x60; during deactivation to indicate that the user hasn&#x27;t completed the asynchronous operation.  The user will have a status of &#x60;DEPROVISIONED&#x60; when the deactivation process is complete.
func (m *UserResource) DeactivateUser(userId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/lifecycle/deactivate", userId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// This operation transitions the user to the status of &#x60;PASSWORD_EXPIRED&#x60; so that the user is required to change their password at their next login.
func (m *UserResource) ExpirePassword(userId string) (*User, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/lifecycle/expire_password?tempPassword=false", userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var user *User

	resp, err := m.client.requestExecutor.Do(req, &user)
	if err != nil {
		return nil, resp, err
	}

	return user, resp, nil
}

// This operation transitions the user to the status of &#x60;PASSWORD_EXPIRED&#x60; and the user&#x27;s password is reset to a temporary password that is returned.
func (m *UserResource) ExpirePasswordAndGetTemporaryPassword(userId string) (*TempPassword, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/lifecycle/expire_password?tempPassword=true", userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var tempPassword *TempPassword

	resp, err := m.client.requestExecutor.Do(req, &tempPassword)
	if err != nil {
		return nil, resp, err
	}

	return tempPassword, resp, nil
}

// Reactivates a user.  This operation can only be performed on users with a &#x60;PROVISIONED&#x60; status.  This operation restarts the activation workflow if for some reason the user activation was not completed when using the activationToken from [Activate User](#activate-user).
func (m *UserResource) ReactivateUser(userId string, qp *query.Params) (*UserActivationToken, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/lifecycle/reactivate", userId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var userActivationToken *UserActivationToken

	resp, err := m.client.requestExecutor.Do(req, &userActivationToken)
	if err != nil {
		return nil, resp, err
	}

	return userActivationToken, resp, nil
}

// This operation resets all factors for the specified user. All MFA factor enrollments returned to the unenrolled state. The user&#x27;s status remains ACTIVE. This link is present only if the user is currently enrolled in one or more MFA factors.
func (m *UserResource) ResetFactors(userId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/lifecycle/reset_factors", userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Generates a one-time token (OTT) that can be used to reset a user&#x27;s password.  The OTT link can be automatically emailed to the user or returned to the API caller and distributed using a custom flow.
func (m *UserResource) ResetPassword(userId string, qp *query.Params) (*ResetPasswordToken, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/lifecycle/reset_password", userId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var resetPasswordToken *ResetPasswordToken

	resp, err := m.client.requestExecutor.Do(req, &resetPasswordToken)
	if err != nil {
		return nil, resp, err
	}

	return resetPasswordToken, resp, nil
}

// Suspends a user.  This operation can only be performed on users with an &#x60;ACTIVE&#x60; status.  The user will have a status of &#x60;SUSPENDED&#x60; when the process is complete.
func (m *UserResource) SuspendUser(userId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/lifecycle/suspend", userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Unlocks a user with a &#x60;LOCKED_OUT&#x60; status and returns them to &#x60;ACTIVE&#x60; status.  Users will be able to login with their current password.
func (m *UserResource) UnlockUser(userId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/lifecycle/unlock", userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Unsuspends a user and returns them to the &#x60;ACTIVE&#x60; state.  This operation can only be performed on users that have a &#x60;SUSPENDED&#x60; status.
func (m *UserResource) UnsuspendUser(userId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/lifecycle/unsuspend", userId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Delete linked objects for a user, relationshipName can be ONLY a primary relationship name
func (m *UserResource) RemoveLinkedObjectForUser(userId string, relationshipName string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/linkedObjects/%v", userId, relationshipName)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Get linked objects for a user, relationshipName can be a primary or associated relationship name
func (m *UserResource) GetLinkedObjectsForUser(userId string, relationshipName string, qp *query.Params) ([]*ResponseLinks, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/linkedObjects/%v", userId, relationshipName)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var responseLinks []*ResponseLinks

	resp, err := m.client.requestExecutor.Do(req, &responseLinks)
	if err != nil {
		return nil, resp, err
	}

	return responseLinks, resp, nil
}

// Lists all roles assigned to a user.
func (m *UserResource) ListAssignedRolesForUser(userId string, qp *query.Params) ([]*Role, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/roles", userId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var role []*Role

	resp, err := m.client.requestExecutor.Do(req, &role)
	if err != nil {
		return nil, resp, err
	}

	return role, resp, nil
}

// Assigns a role to a user.
func (m *UserResource) AssignRoleToUser(userId string, body AssignRoleRequest, qp *query.Params) (*Role, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/roles", userId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("POST", url, body)
	if err != nil {
		return nil, nil, err
	}

	var role *Role

	resp, err := m.client.requestExecutor.Do(req, &role)
	if err != nil {
		return nil, resp, err
	}

	return role, resp, nil
}

// Unassigns a role from a user.
func (m *UserResource) RemoveRoleFromUser(userId string, roleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/roles/%v", userId, roleId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Lists all App targets for an &#x60;APP_ADMIN&#x60; Role assigned to a User. This methods return list may include full Applications or Instances. The response for an instance will have an &#x60;ID&#x60; value, while Application will not have an ID.
func (m *UserResource) ListApplicationTargetsForApplicationAdministratorRoleForUser(userId string, roleId string, qp *query.Params) ([]App, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/roles/%v/targets/catalog/apps", userId, roleId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var application []Application

	resp, err := m.client.requestExecutor.Do(req, &application)
	if err != nil {
		return nil, resp, err
	}

	apps := make([]App, len(application))
	for i := range application {
		apps[i] = &application[i]
	}
	return apps, resp, nil

}

func (m *UserResource) AddAllAppsAsTargetToRole(userId string, roleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/roles/%v/targets/catalog/apps", userId, roleId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *UserResource) RemoveApplicationTargetFromApplicationAdministratorRoleForUser(userId string, roleId string, appName string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/roles/%v/targets/catalog/apps/%v", userId, roleId, appName)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *UserResource) AddApplicationTargetToAdminRoleForUser(userId string, roleId string, appName string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/roles/%v/targets/catalog/apps/%v", userId, roleId, appName)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Remove App Instance Target to App Administrator Role given to a User
func (m *UserResource) RemoveApplicationTargetFromAdministratorRoleForUser(userId string, roleId string, appName string, applicationId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/roles/%v/targets/catalog/apps/%v/%v", userId, roleId, appName, applicationId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Add App Instance Target to App Administrator Role given to a User
func (m *UserResource) AddApplicationTargetToAppAdminRoleForUser(userId string, roleId string, appName string, applicationId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/roles/%v/targets/catalog/apps/%v/%v", userId, roleId, appName, applicationId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *UserResource) ListGroupTargetsForRole(userId string, roleId string, qp *query.Params) ([]*Group, *Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/roles/%v/targets/groups", userId, roleId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var group []*Group

	resp, err := m.client.requestExecutor.Do(req, &group)
	if err != nil {
		return nil, resp, err
	}

	return group, resp, nil
}

func (m *UserResource) RemoveGroupTargetFromRole(userId string, roleId string, groupId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/roles/%v/targets/groups/%v", userId, roleId, groupId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (m *UserResource) AddGroupTargetToRole(userId string, roleId string, groupId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/roles/%v/targets/groups/%v", userId, roleId, groupId)

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Removes all active identity provider sessions. This forces the user to authenticate on the next operation. Optionally revokes OpenID Connect and OAuth refresh and access tokens issued to the user.
func (m *UserResource) ClearUserSessions(userId string, qp *query.Params) (*Response, error) {
	url := fmt.Sprintf("/api/v1/users/%v/sessions", userId)
	if qp != nil {
		url = url + qp.String()
	}

	req, err := m.client.requestExecutor.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}