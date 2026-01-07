"use client";

import {
  createProduct,
  deleteProduct,
  getAllProducts,
  updateProduct,
} from "@/services/products/api";
import React, { useEffect, useState } from "react";

function Page() {
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  const [openModal, setOpenModal] = useState(false);
  const [editPost, setEditPost] = useState<Product | null>(null);

  useEffect(() => {
    GetAll();
  }, []);

  async function GetAll() {
    try {
      const data = await getAllProducts();
      setProducts(data);
    } catch (error) {
      setError(error instanceof Error ? error.message : "Unknown error");
    } finally {
      setLoading(false);
    }
  }

  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault();
    const formData = new FormData(e.currentTarget);
    const productData = {
      name: formData.get("name") as string,
      desc: formData.get("desc") as string,
      price: parseFloat(formData.get("price") as string),
      stock: parseInt(formData.get("stock") as string),
    };

    try {
      if (editPost) {
        await updateProduct(editPost.id, productData);
      } else {
        await createProduct(productData);
      }
      const allProducts = await getAllProducts();
      setOpenModal(false);
      setEditPost(null);
      setProducts(allProducts);
    } catch (error) {
      setError(error instanceof Error ? error.message : "Unknown error");
    }
  }

  async function handleDelete(id: number) {
    try {
      await deleteProduct(id);
      const products = await getAllProducts();
      setProducts(products);
    } catch (error) {
      setError(error instanceof Error ? error.message : "Unknown error");
    }
  }

  if (loading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div>
      <button
        onClick={() => {
          setEditPost(null);
          setOpenModal(true);
        }}
        className="m-4 px-4 py-2 bg-blue-500 text-white rounded-md"
      >
        Create
      </button>

      {openModal && (
        <div className="fixed inset-0 flex items-center justify-center bg-black/50 z-50">
          <div className="bg-white p-6 rounded-lg shadow-lg w-96">
            <h2 className="text-xl font-bold mb-4">
              {editPost ? "Update Product" : "Create Product"}
            </h2>

            <form onSubmit={handleSubmit} className="space-y-4">
              <div>
                <label className="block text-sm font-medium">Name</label>
                <input
                  type="text"
                  name="name"
                  defaultValue={editPost?.name || ""}
                  className="w-full border rounded-md p-2"
                  required
                />
              </div>

              <div>
                <label className="block text-sm font-medium">Description</label>
                <textarea
                  name="desc"
                  defaultValue={editPost?.desc || ""}
                  className="w-full border rounded-md p-2"
                />
              </div>

              <div>
                <label className="block text-sm font-medium">Price</label>
                <input
                  type="number"
                  name="price"
                  step="0.01"
                  defaultValue={editPost?.price || ""}
                  className="w-full border rounded-md p-2"
                  required
                />
              </div>

              <div>
                <label className="block text-sm font-medium">Stock</label>
                <input
                  type="number"
                  name="stock"
                  defaultValue={editPost?.stock || ""}
                  className="w-full border rounded-md p-2"
                  required
                />
              </div>

              <div className="flex justify-end space-x-2">
                <button
                  type="button"
                  onClick={() => {
                    setOpenModal(false);
                    setEditPost(null);
                  }}
                  className="px-4 py-2 bg-gray-300 rounded-md"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  className="px-4 py-2 bg-blue-600 text-white rounded-md"
                >
                  {editPost ? "Update" : "Create"}
                </button>
              </div>
            </form>
          </div>
        </div>
      )}

      <div>
        <ul className="space-y-4 p-4">
          {products
            ? products.map((p) => (
                <li
                  key={p.id}
                  className="p-4 bg-white shadow-md rounded-lg text-gray-700 flex justify-between items-center"
                >
                  <div>
                    <div className="font-bold">{p.name}</div>
                    <div className="text-sm">
                      <div>Desc: {p.desc}</div>
                      <div>Price: {p.price}</div>
                      <div>Stock: {p.stock}</div>
                    </div>
                  </div>
                  <div className="space-x-2">
                    <button
                      onClick={() => {
                        setEditPost(p);
                        setOpenModal(true);
                      }}
                      className="px-3 py-1 bg-blue-500 text-white rounded-md"
                    >
                      Update
                    </button>
                    <button
                      onClick={() => handleDelete(p.id)}
                      className="px-3 py-1 bg-red-500 text-white rounded-md"
                    >
                      Delete
                    </button>
                  </div>
                </li>
              ))
            : "categories is empty"}
        </ul>
      </div>
    </div>
  );
}

export default Page;
