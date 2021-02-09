package schema

import (
	"errors"
	"regexp"
	"strings"
	"strconv"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"

)
// Bill holds the schema definition for the Bill entity.
type Bill struct {
	ent.Schema
}
// Fields of the Bill.
func (Bill) Fields() []ent.Field {
	return []ent.Field{
		field.String("Amount").NotEmpty().
		Validate(func(s string) error {
			_, err := strconv.ParseFloat(s, 64)
				if (err != nil) {
				return errors.New("Amount must be Number")
			}
			return nil
		}),
		field.String("Payer").NotEmpty().
		Validate(func(s string) error {
			if (strings.ContainsAny(s,"1234567890")) {
				return errors.New("Payer name can not have number init")
			}
			return nil
		}),
		field.String("Payercontact").NotEmpty().MinLen(10).MaxLen(10).
		Validate(func(s string) error {
			match, _ := regexp.MatchString("[0]\\d",s)  
			if !match {
				return errors.New("Phone number must be Number and start with 0")
			}
			return nil
		}),
		
		field.Time("Date"),
	}
}

// Edges of the Bill.
func (Bill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("EdgesOfPaytype", Paytype.Type).Ref("EdgesOfBills").Unique(),
		edge.From("EdgesOfOfficer", Financier.Type).Ref("EdgesOfBills").Unique(),
		edge.From("EdgesOfTreatment", Unpaybill.Type).Ref("EdgesOfBills").Unique(),
	}
}
