package api

import (
	uuid "github.com/satori/go.uuid"
)

// SubscriptionDocuments represents subscription documents.
// pkg/database/cosmosdb requires its definition.
type SubscriptionDocuments struct {
	Count                 int                     `json:"_count,omitempty"`
	ResourceID            string                  `json:"_rid,omitempty"`
	SubscriptionDocuments []*SubscriptionDocument `json:"Documents,omitempty"`
}

// SubscriptionDocument represents a subscription document.
// pkg/database/cosmosdb requires its definition.
type SubscriptionDocument struct {
	MissingFields

	ID          string `json:"id,omitempty"`
	ResourceID  string `json:"_rid,omitempty"`
	Timestamp   int    `json:"_ts,omitempty"`
	Self        string `json:"_self,omitempty"`
	ETag        string `json:"_etag,omitempty"`
	Attachments string `json:"_attachments,omitempty"`

	Key Key `json:"key,omitempty"` // also the partition key

	LeaseOwner   *uuid.UUID `json:"leaseOwner,omitempty"`
	LeaseExpires int        `json:"leaseExpires,omitempty"`
	Dequeues     int        `json:"dequeues,omitempty"`

	Deleting bool `json:"deleting,omitempty"`

	Subscription *Subscription `json:"subscription,omitempty"`
}
