// Code generated by hertz generator.

package main

import (
	"db_lab_library/biz/handler/book"
	"db_lab_library/biz/handler/delivery"
	"db_lab_library/biz/handler/order"
	"db_lab_library/biz/handler/out_of_stock_record"
	"db_lab_library/biz/handler/purchase_order"
	"db_lab_library/biz/handler/supplier"
	"db_lab_library/biz/handler/user"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	bookController := r.Group("book")
	{
		bookController.POST("/insert", book.Insert)
		bookController.POST("/get", book.Get)
	}

	userController := r.Group("user")
	{
		userController.POST("", user.Insert)
		userController.POST("/get", user.Get)
		userController.POST("/update_basic_info", user.UpdateBasicInfo)
		userController.POST("/update_credit_rating", user.UpdateCreditRating)
	}

	supplierController := r.Group("supplier")
	{
		supplierController.POST("/get_supply_info", supplier.GetSupplyInfo)
		supplierController.POST("/get_supplier_info", supplier.GetSupplierInfo)
	}

	outOfStockRecordController := r.Group("out_of_stock_record")
	{
		outOfStockRecordController.POST("/insert", out_of_stock_record.Insert)
	}

	orderController := r.Group("order")
	{
		orderController.POST("/insert", order.Insert)
	}

	purchaseOrderController := r.Group("purchase_order")
	{
		purchaseOrderController.POST("/insert", purchase_order.Insert)
		purchaseOrderController.POST("/update_status", purchase_order.UpdateStatus)
	}

	deliveryController := r.Group("delivery")
	{
		deliveryController.POST("/insert", delivery.Insert)
	}
}
