package method

import (
	"gorm.io/gen"
)

type BookMethod interface {
	// SELECR * FROM @@table
	// {{ where }}
	//   {{ if id != 0 }}
	//     id = @id
	//   {{ end }}
	//   {{ if name != "" }}
	//     AND name LIKE CONCAT('%', @name, '%')
	//   {{ end }}
	//  {{ if publisher != "" }}
	//     AND publisher LIKE CONCAT('%', @publisher, '%')
	//   {{ end }}
	//   {{ if keyword != "" }}
	//     AND @keyword MEMBER OF(JSON_EXTRACT(keywords, '$[*]'))
	//   {{ end }}
	//   {{ for i, author := range authors }}
	//     {{ if author != "" }}
	//       AND JSON_CONTAINS(authors, @author, CONCAT('$[', @i, ']')
	// 	   {{ end }}
	//   {{ end }}
	// {{ end }}
	GetByCondition(id int32, name, publisher, keyword string, authors []string) ([]*gen.T, error)

	// SELECT * FROM @@table
	// WHERE @supplierID MEMBER OF(JSON_EXTRACT(supplier_ids, '$[*]'))
	GetBySupplierID(supplierID int32) ([]*gen.T, error)
}
