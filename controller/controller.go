package controller

import (
	"net/http"

	"github.com/pilinux/postmark/model"

	"github.com/gin-gonic/gin"

	"github.com/pilinux/gorest/database"
	"github.com/pilinux/gorest/lib/renderer"

	log "github.com/sirupsen/logrus"
)

// Outbound - postmark webhooks for outbound
func Outbound(c *gin.Context) {
	db := database.GetDB()
	incomingPayload := model.PostmarkOutbound{}
	incomingPayloadFinal := model.PostmarkOutbound{}

	// bind JSON
	if err := c.ShouldBindJSON(&incomingPayload); err != nil {
		renderer.Render(c, gin.H{"msg": "bad request"}, http.StatusBadRequest)
		return
	}

	// check events based on record types
	recordType := incomingPayload.RecordType
	if recordType == "Delivery" {
		incomingPayloadFinal.EventAt = incomingPayload.DeliveredAt
		incomingPayloadFinal.To = incomingPayload.Recipient
	}
	if recordType == "Bounce" {
		incomingPayloadFinal.EventAt = incomingPayload.BouncedAt
		incomingPayloadFinal.To = incomingPayload.Email
	}
	if recordType == "SpamComplaint" {
		incomingPayloadFinal.EventAt = incomingPayload.BouncedAt
		incomingPayloadFinal.To = incomingPayload.Email
	}
	if recordType == "Open" {
		incomingPayloadFinal.EventAt = incomingPayload.ReceivedAt
		incomingPayloadFinal.To = incomingPayload.Recipient
	}
	if recordType == "Click" {
		incomingPayloadFinal.EventAt = incomingPayload.ReceivedAt
		incomingPayloadFinal.To = incomingPayload.Recipient
	}
	if recordType == "SubscriptionChange" {
		incomingPayloadFinal.EventAt = incomingPayload.ChangedAt
		incomingPayloadFinal.To = incomingPayload.Recipient
	}

	// user must not be able to manipulate all fields like PK
	incomingPayloadFinal.RecordType = incomingPayload.RecordType
	incomingPayloadFinal.Type = incomingPayload.Type
	incomingPayloadFinal.TypeCode = incomingPayload.TypeCode
	incomingPayloadFinal.MessageID = incomingPayload.MessageID
	incomingPayloadFinal.Tag = incomingPayload.Tag
	incomingPayloadFinal.From = incomingPayload.From
	incomingPayloadFinal.ServerID = incomingPayload.ServerID

	tx := db.Begin()
	if err := tx.Create(&incomingPayloadFinal).Error; err != nil {
		tx.Rollback()
		log.WithError(err).Error("custom error code: 10101")
		renderer.Render(c, gin.H{"msg": "internal server error"}, http.StatusInternalServerError)
	} else {
		tx.Commit()
		renderer.Render(c, gin.H{"msg": "OK"}, http.StatusOK)
	}
}
