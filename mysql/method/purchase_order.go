package method

import "gorm.io/gen"

type PurchaseOrder interface {
	// SELECT * FROM @@table
	// {{ where }}
	//   {{ if id != 0 }}
	//     id = @id
	//   {{ end }}
	//   {{ if outOfStockRecordID != 0 }}
	//     AND out_of_stock_record_id = @outOfStockRecordID
	//   {{ end }}
	// {{ end }}
	GetByCondition(id, outOfStockRecordID int32) ([]*gen.T, error)
}
