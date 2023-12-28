// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameBook = "book"

// Book mapped from table <book>
type Book struct {
	ID           int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name         string    `gorm:"column:name;not null" json:"name"`
	Authors      string    `gorm:"column:authors;not null" json:"authors"`
	Publisher    string    `gorm:"column:publisher;not null" json:"publisher"`
	Price        float64   `gorm:"column:price;not null" json:"price"`
	Keywords     string    `gorm:"column:keywords;not null" json:"keywords"`
	Catalogue    string    `gorm:"column:catalogue" json:"catalogue"`
	Cover        string    `gorm:"column:cover" json:"cover"`
	Inventory    int32     `gorm:"column:inventory;not null" json:"inventory"`
	PublisherIds string    `gorm:"column:publisher_ids" json:"publisher_ids"`
	Location     string    `gorm:"column:location" json:"location"`
	SubbookIds   string    `gorm:"column:subbook_ids" json:"subbook_ids"`
	CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName Book's table name
func (*Book) TableName() string {
	return TableNameBook
}
