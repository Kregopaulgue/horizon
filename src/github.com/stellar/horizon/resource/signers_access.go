package resource

import (
	"github.com/stellar/horizon/db2/core"
	"github.com/stellar/horizon/httpx"
	"github.com/stellar/horizon/render/hal"
	"golang.org/x/net/context"
)

func (this *SignersAccess) Populate(ctx context.Context, row core.SignersAccess) {
	this.AccessGiverID = row.Accessgiverid
	this.AccessTakerID = row.Accesstakerid

	lb := hal.LinkBuilder{httpx.BaseURL(ctx)}
	this.Links.Self = lb.Linkf("/signersaccesses/%d", row.Accesstakerid, row.Accessgiverid)
	this.Links.SignersAccessMaker = lb.Linkf("/accounts/%s", row.Accessgiverid)
	return
}

func (this SignersAccess) PagingToken() string {
	return this.PT
}
