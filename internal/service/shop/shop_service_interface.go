package shop

import "context"

type ShopService interface {
	PurchaseItem(ctx context.Context, userId int, itemId int) (statusCode int, err error)
}
