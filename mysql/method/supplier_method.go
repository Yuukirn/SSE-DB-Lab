package method

import "gorm.io/gen"

type SupplierMethod interface {
	// SELECR * FROM @@table
	// {{ where }}
	//   {{ if id != 0 }}
	//     id = @id
	//   {{ end }}
	//   {{ if name != "" }}
	//     AND name LIKE CONCAT('%', @name, '%')
	//   {{ end }}
	// {{ end }}
	GetByCondition(id int32, name string) ([]*gen.T, error)
}
