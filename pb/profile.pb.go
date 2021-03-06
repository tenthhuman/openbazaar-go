// Code generated by protoc-gen-go.
// source: profile.proto
// DO NOT EDIT!

/*
Package profile is a generated protocol buffer package.

It is generated from these files:
	profile.proto

It has these top-level messages:
	Profile
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Profile struct {
	Handle           string                   `protobuf:"bytes,1,opt,name=handle" json:"handle,omitempty"`
	Name             string                   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Location         string                   `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
	About            string                   `protobuf:"bytes,4,opt,name=about" json:"about,omitempty"`
	ShortDescription string                   `protobuf:"bytes,5,opt,name=shortDescription" json:"shortDescription,omitempty"`
	Website          string                   `protobuf:"bytes,6,opt,name=website" json:"website,omitempty"`
	Email            string                   `protobuf:"bytes,7,opt,name=email" json:"email,omitempty"`
	PhoneNumber      string                   `protobuf:"bytes,8,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
	Social           []*Profile_SocialAccount `protobuf:"bytes,9,rep,name=social" json:"social,omitempty"`
	AvgRating        uint32                   `protobuf:"varint,10,opt,name=avgRating" json:"avgRating,omitempty"`
	NumRatings       uint32                   `protobuf:"varint,11,opt,name=numRatings" json:"numRatings,omitempty"`
	Nsfw             bool                     `protobuf:"varint,12,opt,name=nsfw" json:"nsfw,omitempty"`
	Vendor           bool                     `protobuf:"varint,13,opt,name=vendor" json:"vendor,omitempty"`
	Moderator        bool                     `protobuf:"varint,14,opt,name=moderator" json:"moderator,omitempty"`
	PrimaryColor     string                   `protobuf:"bytes,15,opt,name=primaryColor" json:"primaryColor,omitempty"`
	SecondaryColor   string                   `protobuf:"bytes,16,opt,name=secondaryColor" json:"secondaryColor,omitempty"`
	TextColor        string                   `protobuf:"bytes,17,opt,name=textColor" json:"textColor,omitempty"`
	AvatarHashes     *Profile_Image           `protobuf:"bytes,18,opt,name=avatarHashes" json:"avatarHashes,omitempty"`
	HeaderHashes     *Profile_Image           `protobuf:"bytes,19,opt,name=headerHashes" json:"headerHashes,omitempty"`
	FollowerCount    uint32                   `protobuf:"varint,20,opt,name=followerCount" json:"followerCount,omitempty"`
	FollowingCount   uint32                   `protobuf:"varint,21,opt,name=followingCount" json:"followingCount,omitempty"`
	ListingCount     uint32                   `protobuf:"varint,22,opt,name=listingCount" json:"listingCount,omitempty"`
	LastModified     uint32                   `protobuf:"varint,23,opt,name=lastModified" json:"lastModified,omitempty"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Profile) GetSocial() []*Profile_SocialAccount {
	if m != nil {
		return m.Social
	}
	return nil
}

func (m *Profile) GetAvatarHashes() *Profile_Image {
	if m != nil {
		return m.AvatarHashes
	}
	return nil
}

func (m *Profile) GetHeaderHashes() *Profile_Image {
	if m != nil {
		return m.HeaderHashes
	}
	return nil
}

type Profile_SocialAccount struct {
	Type     string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Proof    string `protobuf:"bytes,3,opt,name=proof" json:"proof,omitempty"`
}

func (m *Profile_SocialAccount) Reset()                    { *m = Profile_SocialAccount{} }
func (m *Profile_SocialAccount) String() string            { return proto.CompactTextString(m) }
func (*Profile_SocialAccount) ProtoMessage()               {}
func (*Profile_SocialAccount) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

type Profile_Image struct {
	Tiny     string `protobuf:"bytes,1,opt,name=tiny" json:"tiny,omitempty"`
	Small    string `protobuf:"bytes,2,opt,name=small" json:"small,omitempty"`
	Medium   string `protobuf:"bytes,3,opt,name=medium" json:"medium,omitempty"`
	Large    string `protobuf:"bytes,4,opt,name=large" json:"large,omitempty"`
	Original string `protobuf:"bytes,5,opt,name=original" json:"original,omitempty"`
}

func (m *Profile_Image) Reset()                    { *m = Profile_Image{} }
func (m *Profile_Image) String() string            { return proto.CompactTextString(m) }
func (*Profile_Image) ProtoMessage()               {}
func (*Profile_Image) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 1} }

func init() {
	proto.RegisterType((*Profile)(nil), "Profile")
	proto.RegisterType((*Profile_SocialAccount)(nil), "Profile.SocialAccount")
	proto.RegisterType((*Profile_Image)(nil), "Profile.Image")
}

var fileDescriptor3 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x93, 0xdd, 0x6e, 0x13, 0x3d,
	0x10, 0x86, 0xd5, 0xaf, 0xcd, 0xdf, 0xe4, 0xe7, 0x2b, 0xa6, 0x04, 0x2b, 0x42, 0x28, 0xaa, 0x10,
	0xaa, 0x38, 0xc8, 0x41, 0xb8, 0x02, 0x54, 0x0e, 0xe0, 0x00, 0x84, 0x16, 0x71, 0x01, 0xce, 0xee,
	0x24, 0xb1, 0xe4, 0xb5, 0x23, 0xdb, 0x9b, 0x10, 0x71, 0xa5, 0xdc, 0x0d, 0xf6, 0xac, 0xb3, 0xc9,
	0x16, 0x8e, 0xe2, 0xf7, 0x9d, 0x67, 0x66, 0xed, 0xc9, 0x0c, 0x8c, 0x77, 0xd6, 0xac, 0xa5, 0xc2,
	0x45, 0xf8, 0xf5, 0xe6, 0xfe, 0x77, 0x0f, 0x7a, 0xdf, 0x6a, 0x87, 0x4d, 0xa1, 0xbb, 0x15, 0xba,
	0x50, 0xc8, 0xaf, 0xe6, 0x57, 0x0f, 0x83, 0x2c, 0x29, 0xc6, 0xe0, 0x46, 0x8b, 0x12, 0xf9, 0x7f,
	0xe4, 0xd2, 0x99, 0xcd, 0xa0, 0xaf, 0x4c, 0x2e, 0xbc, 0x34, 0x9a, 0x5f, 0x93, 0xdf, 0x68, 0x76,
	0x07, 0x1d, 0xb1, 0x32, 0x95, 0xe7, 0x37, 0x14, 0xa8, 0x05, 0x7b, 0x07, 0xb7, 0x6e, 0x6b, 0xac,
	0xff, 0x88, 0x2e, 0xb7, 0x72, 0x47, 0x99, 0x1d, 0x02, 0xfe, 0xf2, 0x19, 0x87, 0xde, 0x01, 0x57,
	0x4e, 0x7a, 0xe4, 0x5d, 0x42, 0x4e, 0x32, 0xd6, 0xc6, 0x52, 0x48, 0xc5, 0x7b, 0x75, 0x6d, 0x12,
	0x6c, 0x0e, 0xc3, 0xdd, 0xd6, 0x68, 0xfc, 0x5a, 0x95, 0x2b, 0xb4, 0xbc, 0x4f, 0xb1, 0x4b, 0x8b,
	0x2d, 0xa0, 0xeb, 0x4c, 0x2e, 0x85, 0xe2, 0x83, 0xf9, 0xf5, 0xc3, 0x70, 0x39, 0x5d, 0xa4, 0x57,
	0x2f, 0xbe, 0x93, 0xfd, 0x21, 0xcf, 0x4d, 0xa5, 0x7d, 0x96, 0x28, 0xf6, 0x0a, 0x06, 0x62, 0xbf,
	0xc9, 0xc2, 0x83, 0xf4, 0x86, 0x43, 0xa8, 0x37, 0xce, 0xce, 0x06, 0x7b, 0x0d, 0xa0, 0xab, 0xb2,
	0x16, 0x8e, 0x0f, 0x29, 0x7c, 0xe1, 0x50, 0xc7, 0xdc, 0xfa, 0xc0, 0x47, 0x21, 0xd2, 0xcf, 0xe8,
	0x1c, 0xbb, 0xbb, 0x47, 0x5d, 0x18, 0xcb, 0xc7, 0xe4, 0x26, 0x15, 0xbf, 0x54, 0x9a, 0x02, 0xad,
	0xf0, 0x21, 0x34, 0xa1, 0xd0, 0xd9, 0x60, 0xf7, 0x30, 0xda, 0x59, 0x59, 0x0a, 0x7b, 0x7c, 0x34,
	0x2a, 0x00, 0xff, 0xd3, 0xd3, 0x5a, 0x1e, 0x7b, 0x0b, 0x13, 0x87, 0xb9, 0xd1, 0x45, 0x43, 0xdd,
	0x12, 0xf5, 0xc4, 0x8d, 0x5f, 0xf2, 0xf8, 0xd3, 0xd7, 0xc8, 0x33, 0x42, 0xce, 0x06, 0x5b, 0xc2,
	0x48, 0xec, 0x85, 0x17, 0xf6, 0x93, 0x70, 0x5b, 0x74, 0x9c, 0x05, 0x60, 0xb8, 0x9c, 0x34, 0x7d,
	0xfa, 0x5c, 0x8a, 0x0d, 0x66, 0x2d, 0x26, 0xe6, 0x6c, 0x51, 0x84, 0xbb, 0xa6, 0x9c, 0xe7, 0xff,
	0xce, 0xb9, 0x64, 0xd8, 0x1b, 0x18, 0xaf, 0x8d, 0x52, 0xe6, 0x80, 0xf6, 0x31, 0xb6, 0x9c, 0xdf,
	0x51, 0xfb, 0xda, 0x66, 0x7c, 0x53, 0x6d, 0x84, 0x7e, 0xd6, 0xd8, 0x0b, 0xc2, 0x9e, 0xb8, 0xb1,
	0x3f, 0x4a, 0x3a, 0xdf, 0x50, 0x53, 0xa2, 0x5a, 0x1e, 0x31, 0xc2, 0xf9, 0x2f, 0xa6, 0x90, 0x6b,
	0x89, 0x05, 0x7f, 0x99, 0x98, 0x0b, 0x6f, 0xf6, 0x03, 0xc6, 0xad, 0x41, 0x88, 0x7f, 0xa1, 0x3f,
	0xee, 0x4e, 0xab, 0x40, 0xe7, 0x38, 0xf4, 0x95, 0x43, 0x7b, 0xb1, 0x0c, 0x8d, 0x8e, 0x83, 0x19,
	0x36, 0xca, 0xac, 0xd3, 0x36, 0xd4, 0x62, 0xf6, 0x0b, 0x3a, 0xd4, 0x03, 0x2a, 0x27, 0xf5, 0xb1,
	0x29, 0x17, 0xce, 0x31, 0xc5, 0x95, 0x42, 0xa9, 0x54, 0xab, 0x16, 0x71, 0x4e, 0x4a, 0x2c, 0x64,
	0x55, 0xa6, 0x4a, 0x49, 0x45, 0x5a, 0x09, 0xbb, 0xc1, 0xd3, 0x56, 0x91, 0x88, 0x57, 0x32, 0x56,
	0x6e, 0xa4, 0x0e, 0x93, 0x5d, 0x6f, 0x53, 0xa3, 0x57, 0x5d, 0x5a, 0xf1, 0xf7, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x1b, 0x26, 0x1d, 0x4e, 0xf3, 0x03, 0x00, 0x00,
}
