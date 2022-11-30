package ccloud

import (
	"context"
	"errors"
	"net/http"

	"github.com/dghubble/sling"
)

const (
	// USERS base endpoint for user operations
	USERS = "/api/users"
	// USER specific user operation endpoint
	USER            = USERS + "/%s"
	SERVICEACCOUNTS = "/api/service_accounts"
	SERVICEACCOUNT  = SERVICEACCOUNTS + "/%d"
	// deprecated: legacy endpoint for email invite to join org. org id in path is ignored, only the org id in the auth context matters
	INVITEURL = "/api/organizations/0/invites"
	// new invitation flow
	INVITATIONURL = "/api/invitations"
	// USER_PROFILES base endpoint for use profile operations
	USER_PROFILES = "/api/user_profiles"
	// customer-initiated org deletion flow
	SUSPEND_ORG_URL_TEMPLATE = "/api/organizations/%d/suspend"
)

// UserService provides methods for managing users on Confluent Control Plane.
type LoginRealmService struct {
	client *http.Client
	sling  *sling.Sling
}

var _ LoginRealm = (*LoginRealmService)(nil)

// NewUserService returns a new UserService.
func NewLoginRealmService(client *Client) *LoginRealmService {
	return &LoginRealmService{
		client: client.HttpClient,
		sling:  client.sling,
	}
}

// LoginRealm gets a user's login realm information. can be used to determine if the user is an SSO user
func (s *LoginRealmService) LoginRealm(_ context.Context, req *GetLoginRealmRequest) (*GetLoginRealmReply, error) {
	reply := new(GetLoginRealmReply)

	if req == nil || req.Email == "" {
		return nil, errors.New("non-nil user object must be passed and have Email set in order to call LoginRealm")
	}

	_, err := s.sling.New().Get("/api/login/realm").QueryStruct(req).Receive(reply, reply)
	if err := ReplyErr(reply, err); err != nil {
		return nil, WrapErr(err, "error checking login realm")
	}

	return reply, nil
}
