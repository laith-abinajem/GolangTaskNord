package keys

import (
	"fmt"
)

const (
	KEY_CTX_Session = "session"
)

func KEY_Transaction(transactionId uint) string {
	return fmt.Sprintf("transaction_%d", transactionId)
}

func KEY_TotalSales(tenantID, productID uint) string {
	return fmt.Sprintf("total_sales_tenant_%d_product_%d", tenantID, productID)
}

func KEY_TopProducts() string {
	return "top_selling_products"
}
