/*
 * Copyright © 2018-2020 Musing Studio LLC.
 *
 * This file is part of WriteFreely.
 *
 * WriteFreely is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, included
 * in the LICENSE file in this source code package.
 */

package writefreely

import (
	"net/http"

	"github.com/writeas/impart"
)

// Commonly returned HTTP errors
var (
	ErrBadFormData    = impart.HTTPError{Status: http.StatusBadRequest, Message: "Expected valid form data."}
	ErrBadJSON        = impart.HTTPError{Status: http.StatusBadRequest, Message: "Expected valid JSON object."}
	ErrBadJSONArray   = impart.HTTPError{Status: http.StatusBadRequest, Message: "Expected valid JSON array."}
	ErrBadAccessToken = impart.HTTPError{Status: http.StatusUnauthorized, Message: "Invalid access token."}
	ErrNoAccessToken  = impart.HTTPError{Status: http.StatusBadRequest, Message: "Authorization token required."}
	ErrNotLoggedIn    = impart.HTTPError{Status: http.StatusUnauthorized, Message: "Not logged in."}

	ErrForbiddenCollection        = impart.HTTPError{Status: http.StatusForbidden, Message: "You don't have permission to add to this collection."}
	ErrForbiddenCollectionAccess  = impart.HTTPError{Status: http.StatusForbidden, Message: "You don't have permission to access this collection."}
	ErrForbiddenEditPost          = impart.HTTPError{Status: http.StatusForbidden, Message: "You don't have permission to update this post."}
	ErrUnauthorizedEditPost       = impart.HTTPError{Status: http.StatusUnauthorized, Message: "Invalid editing credentials."}
	ErrUnauthorizedGeneral        = impart.HTTPError{Status: http.StatusUnauthorized, Message: "You don't have permission to do that."}
	ErrBadRequestedType           = impart.HTTPError{Status: http.StatusNotAcceptable, Message: "Bad requested Content-Type."}
	ErrCollectionUnauthorizedRead = impart.HTTPError{Status: http.StatusUnauthorized, Message: "You don't have permission to access this collection."}

	ErrNoPublishableContent = impart.HTTPError{Status: http.StatusBadRequest, Message: "Supply something to publish."}

	ErrInternalGeneral       = impart.HTTPError{Status: http.StatusInternalServerError, Message: "The humans messed something up. They've been notified."}
	ErrInternalCookieSession = impart.HTTPError{Status: http.StatusInternalServerError, Message: "Could not get cookie session."}

	ErrUnavailable = impart.HTTPError{Status: http.StatusServiceUnavailable, Message: "Service temporarily unavailable due to high load."}

	ErrCollectionNotFound     = impart.HTTPError{Status: http.StatusNotFound, Message: "Collection doesn't exist."}
	ErrCollectionGone         = impart.HTTPError{Status: http.StatusGone, Message: "This blog was unpublished."}
	ErrCollectionPageNotFound = impart.HTTPError{Status: http.StatusNotFound, Message: "Collection page doesn't exist."}
	ErrPostNotFound           = impart.HTTPError{Status: http.StatusNotFound, Message: "Post not found."}
	ErrPostBanned             = impart.HTTPError{Status: http.StatusGone, Message: "Post removed."}
	ErrPostUnpublished        = impart.HTTPError{Status: http.StatusGone, Message: "Post unpublished by author."}
	ErrPostFetchError         = impart.HTTPError{Status: http.StatusInternalServerError, Message: "We encountered an error getting the post. The humans have been alerted."}

	ErrUserNotFound       = impart.HTTPError{Status: http.StatusNotFound, Message: "User doesn't exist."}
	ErrRemoteUserNotFound = impart.HTTPError{Status: http.StatusNotFound, Message: "Remote user not found."}
	ErrUserNotFoundEmail  = impart.HTTPError{Status: http.StatusNotFound, Message: "Please enter your username instead of your email address."}

	ErrUserSilenced = impart.HTTPError{Status: http.StatusForbidden, Message: "Account is silenced."}

	ErrDisabledPasswordAuth = impart.HTTPError{Status: http.StatusForbidden, Message: "Password authentication is disabled."}
)

// Post operation errors
var (
	ErrPostNoUpdatableVals = impart.HTTPError{Status: http.StatusBadRequest, Message: "Supply some properties to update."}
)
