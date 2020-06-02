package models

import (
	"database/sql"
	"encoding/json"
	"log"
)

type Session struct {
	SessionId               string   `protobuf:"bytes,1,opt,name=sessionId,proto3" json:"sessionId,omitempty"`
	UserId                  string   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	TenantId                string   `protobuf:"bytes,3,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	UserFirstName           string   `protobuf:"bytes,4,opt,name=userFirstName,proto3" json:"userFirstName,omitempty"`
	UserLastName            string   `protobuf:"bytes,5,opt,name=userLastName,proto3" json:"userLastName,omitempty"`
	UserEmail               string   `protobuf:"bytes,6,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	UserRole                string   `protobuf:"bytes,7,opt,name=userRole,proto3" json:"userRole,omitempty"`
	Data                    string   `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
	Auth0IdToken            string   `protobuf:"bytes,9,opt,name=auth0IdToken,proto3" json:"auth0IdToken,omitempty"`
	AwsCredsAccessKeyId     string   `protobuf:"bytes,10,opt,name=awsCredsAccessKeyId,proto3" json:"awsCredsAccessKeyId,omitempty"`
	AwsCredsSecretAccessKey string   `protobuf:"bytes,11,opt,name=awsCredsSecretAccessKey,proto3" json:"awsCredsSecretAccessKey,omitempty"`
	AwsCredsSessionToken    string   `protobuf:"bytes,12,opt,name=awsCredsSessionToken,proto3" json:"awsCredsSessionToken,omitempty"`
	AwsCredsExpiration      string   `protobuf:"bytes,13,opt,name=awsCredsExpiration,proto3" json:"awsCredsExpiration,omitempty"`
}
