package order

import (
	"context"
	"route256/loms/internal/model"

	"github.com/jackc/pgx/v5"
)

// OrderPay - транзакция на покупку
func (o *OrderRepository) OrderPay(ctx context.Context, orderId model.OrderId, order *model.Order) error {
	// Получаем стоки
	skus := make([]uint32, len(order.Items))
	for i := range order.Items {
		skus[i] = order.Items[i].Sku
	}

	const queryGetItem = `
		SELECT sku, total_count, reserved
		FROM stocks
		WHERE sku = ANY($1)
	`
	currentItems := make(map[uint32]model.StockItem)

	rows, err := o.Conn.Query(ctx, queryGetItem, skus)
	if err != nil {
		return err
	}

	for rows.Next() {
		item := model.StockItem{}
		rows.Scan(&item.Sku, &item.TotalCount, &item.Reserved)
		currentItems[item.Sku] = item
	}

	// Делаем транзакцию на покупку
	tx, err := o.Conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Меняем статус
	const queryStatus = `
		UPDATE orders
		SET status = @status
		WHERE id = @id
	`
	argsStatus := pgx.NamedArgs{
		"status": model.StatusPayed,
		"id":     int(orderId),
	}
	_, err = tx.Exec(ctx, queryStatus, argsStatus)
	if err != nil {
		return err
	}

	// Меняем стоки
	batch := &pgx.Batch{}
	for _, itm := range order.Items {
		remainsReserved := currentItems[itm.Sku].Reserved - uint64(itm.Count)

		const queryReserveRemove = `
			UPDATE stocks
			SET reserved = @reserved
			WHERE sku = @sku
		`
		args := pgx.NamedArgs{
			"reserved": remainsReserved,
			"sku":      itm.Sku,
		}

		batch.Queue(queryReserveRemove, args)

		remainsTotalCount := currentItems[itm.Sku].TotalCount - uint64(itm.Count)

		const queryTotalCountRemove = `
		UPDATE stocks
		SET total_count = @total_count
		WHERE sku = @sku
	`
		argsTotalCount := pgx.NamedArgs{
			"total_count": remainsTotalCount,
			"sku":         itm.Sku,
		}

		batch.Queue(queryTotalCountRemove, argsTotalCount)
	}

	res := tx.SendBatch(ctx, batch)
	for i := 0; i < batch.Len(); i++ {
		_, err := res.Exec()
		if err != nil {
			return err
		}
	}
	res.Close() // если поставить defer, то появляется ошибка

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
