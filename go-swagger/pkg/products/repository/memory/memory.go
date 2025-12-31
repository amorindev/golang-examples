package memory

import "example.com/pkg/products/domain"

// products stores products in memory.
// In a real-world environment, this data is usually persisted in a database.
var products = []*domain.Product{}