package models

type Actor struct {
	ActorURI                  string
	Type                      string
	PreferredUsername         string
	Inbox                     string
	Outbox                    string
	Name                      string
	Summary                   string
	URL                       string
	Followers                 []string
	Following                 []string
	ManuallyApprovesFollowers bool
	Discoverable              bool
	Published                 string
	PublicKeyPem              string
	PrivateKeyPem             string
}

//type AsActor struct {
//	AID  string
//	Type string
//	Name string
//}
//
//type AsObject struct {
//	AID  string
//	Type string
//	Name string
//}
//
//type AsLink struct {
//	Type      string
//	Href      string
//	HrefLang  string
//	MediaType string
//	Name      string
//}
//
//type AsActivity struct {
//	AID    string
//	Type   string
//	Name   string
//	Actor  AsActor
//	Object AsObject
//}
//
//type AsIntransitiveActivity struct {
//}
//
//type AsCollection struct {
//}
//
//type AsOrderedCollection struct {
//}
//
//type AsCollectionPage struct {
//}
//
//type AsOrderedCollectionPage struct {
//}
