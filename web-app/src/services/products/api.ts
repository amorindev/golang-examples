const API_BASE = "http://localhost:8082/v1/products";
const token = "123";

// Get All
export async function getAllProducts(): Promise<Product[]> {
  const resp = await fetch(API_BASE, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
  });

  if (!resp.ok) {
    const errorData = await resp.json().catch(() => null);
    throw new Error(errorData?.message || "Failed to fetch products");
  }
  return resp.json();
}

// Create product
export async function createProduct(
  product: Omit<Product, "id">
): Promise<Product> {
  const resp = await fetch(API_BASE, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(product),
  });

  if (!resp.ok) {
    const errorData = await resp.json().catch(() => null);
    throw new Error(errorData?.message || "Failed to create product");
  }
  return await resp.json();
}

// Update product
export async function updateProduct(
  id: number,
  product: Omit<Product, "id">
): Promise<Product> {
  const resp = await fetch(`${API_BASE}/${id}`, {
    method: "PUT",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(product),
  });

  if (!resp.ok) {
    const errorData = await resp.json().catch(() => null);
    throw new Error(errorData?.message || "Failed to update product");
  }
  return resp.json();
}

// Delete product
export async function deleteProduct(id: number): Promise<void> {
  const resp = await fetch(`${API_BASE}/${id}`, { method: "DELETE" });
  if (!resp.ok) {
    const errorData = await resp.json().catch(() => null);
    throw new Error(errorData?.message || "Failed to delete product");
  }
}
