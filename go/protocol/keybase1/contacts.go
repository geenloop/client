// Auto-generated by avdl-compiler v1.3.29 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/contacts.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type ContactComponent struct {
	Label       string          `codec:"label" json:"label"`
	PhoneNumber *RawPhoneNumber `codec:"phoneNumber,omitempty" json:"phoneNumber,omitempty"`
	Email       *EmailAddress   `codec:"email,omitempty" json:"email,omitempty"`
}

func (o ContactComponent) DeepCopy() ContactComponent {
	return ContactComponent{
		Label: o.Label,
		PhoneNumber: (func(x *RawPhoneNumber) *RawPhoneNumber {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.PhoneNumber),
		Email: (func(x *EmailAddress) *EmailAddress {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Email),
	}
}

type Contact struct {
	Name       string             `codec:"name" json:"name"`
	Components []ContactComponent `codec:"components" json:"components"`
}

func (o Contact) DeepCopy() Contact {
	return Contact{
		Name: o.Name,
		Components: (func(x []ContactComponent) []ContactComponent {
			if x == nil {
				return nil
			}
			ret := make([]ContactComponent, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.Components),
	}
}

type ResolvedContact struct {
	DisplayName  string           `codec:"displayName" json:"displayName"`
	ContactIndex int              `codec:"contactIndex" json:"contactIndex"`
	Component    ContactComponent `codec:"component" json:"component"`
	Err          *string          `codec:"err,omitempty" json:"err,omitempty"`
	Resolved     bool             `codec:"resolved" json:"resolved"`
	Uid          UID              `codec:"uid" json:"uid"`
	Username     string           `codec:"username" json:"username"`
	FullName     string           `codec:"fullName" json:"fullName"`
}

func (o ResolvedContact) DeepCopy() ResolvedContact {
	return ResolvedContact{
		DisplayName:  o.DisplayName,
		ContactIndex: o.ContactIndex,
		Component:    o.Component.DeepCopy(),
		Err: (func(x *string) *string {
			if x == nil {
				return nil
			}
			tmp := (*x)
			return &tmp
		})(o.Err),
		Resolved: o.Resolved,
		Uid:      o.Uid.DeepCopy(),
		Username: o.Username,
		FullName: o.FullName,
	}
}

type LookupContactListArg struct {
	SessionID      int        `codec:"sessionID" json:"sessionID"`
	Contacts       []Contact  `codec:"contacts" json:"contacts"`
	UserRegionCode RegionCode `codec:"userRegionCode" json:"userRegionCode"`
}

type ContactsInterface interface {
	LookupContactList(context.Context, LookupContactListArg) ([]ResolvedContact, error)
}

func ContactsProtocol(i ContactsInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.contacts",
		Methods: map[string]rpc.ServeHandlerDescription{
			"lookupContactList": {
				MakeArg: func() interface{} {
					var ret [1]LookupContactListArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]LookupContactListArg)
					if !ok {
						err = rpc.NewTypeError((*[1]LookupContactListArg)(nil), args)
						return
					}
					ret, err = i.LookupContactList(ctx, typedArgs[0])
					return
				},
			},
		},
	}
}

type ContactsClient struct {
	Cli rpc.GenericClient
}

func (c ContactsClient) LookupContactList(ctx context.Context, __arg LookupContactListArg) (res []ResolvedContact, err error) {
	err = c.Cli.Call(ctx, "keybase.1.contacts.lookupContactList", []interface{}{__arg}, &res)
	return
}
