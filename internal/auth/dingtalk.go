package auth

import (
	"context"
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/dingtalk/contact_1_0"
	"github.com/alibabacloud-go/dingtalk/oauth2_1_0"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/tang95/x-seek/internal/model"
	"gorm.io/gorm"
	"net/url"
)

const (
	DINGTALK = "dingtalk"
)

type dingtalkAuth struct {
	*Auth
	clientID      string
	AuthURL       string
	clientSecret  string
	dingtalkOauth *oauth2_1_0.Client
	contactClient *contact_1_0.Client
}

func newDingtalk(clientID string, clientSecret string, auth *Auth) OAuth {
	dingtalkOauth, _ := oauth2_1_0.NewClient(&openapi.Config{
		Protocol: tea.String("https"),
		RegionId: tea.String("central"),
	})
	contactClient, _ := contact_1_0.NewClient(&openapi.Config{
		Protocol: tea.String("https"),
		RegionId: tea.String("central"),
	})
	return &dingtalkAuth{
		clientID:      clientID,
		clientSecret:  clientSecret,
		AuthURL:       "https://login.dingtalk.com/oauth2/auth",
		Auth:          auth,
		dingtalkOauth: dingtalkOauth,
		contactClient: contactClient,
	}
}

func (d *dingtalkAuth) LoginByCode(ctx context.Context, code string, autoRegister bool) (*User, error) {
	dUser, err := d.getDingtalkUserByCode(ctx, code)
	if err != nil {
		return nil, err
	}
	user, err := d.service.GetUserByDingtalkID(ctx, *dUser.UnionId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if user == nil {
		if autoRegister {
			user, err = d.service.CreateUser(ctx, &model.User{
				Name:       *dUser.Nick,
				Avatar:     *dUser.AvatarUrl,
				Role:       model.Member,
				DingtalkID: *dUser.UnionId,
			})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, errors.New("user not found")
		}
	}
	return &User{
		ID:          user.ID,
		Name:        user.Name,
		Avatar:      user.Avatar,
		Description: user.Description,
		Role:        user.Role,
	}, nil
}

func (d *dingtalkAuth) AuthorizeUrl(ctx context.Context) string {
	baseUrl, _ := url.Parse(d.AuthURL)
	query := baseUrl.Query()
	query.Add("client_id", d.clientID)
	query.Add("redirect_uri", d.config.Domain+"/login/dingtalk")
	query.Add("response_type", "code")
	query.Add("scope", "openid")
	query.Add("prompt", "consent")
	baseUrl.RawQuery = query.Encode()
	return baseUrl.String()
}

func (d *dingtalkAuth) getDingtalkUserByCode(ctx context.Context, code string) (*contact_1_0.GetUserResponseBody, error) {
	token, err := d.dingtalkOauth.GetUserToken(&oauth2_1_0.GetUserTokenRequest{
		ClientId:     tea.String(d.clientID),
		ClientSecret: tea.String(d.clientSecret),
		Code:         tea.String(code),
		GrantType:    tea.String("authorization_code"),
	})
	if err != nil {
		return nil, err
	}
	response, err := d.contactClient.GetUserWithOptions(tea.String("me"), &contact_1_0.GetUserHeaders{
		XAcsDingtalkAccessToken: token.Body.AccessToken,
	}, &util.RuntimeOptions{})
	if err != nil {
		return nil, err
	}
	return response.Body, nil
}
