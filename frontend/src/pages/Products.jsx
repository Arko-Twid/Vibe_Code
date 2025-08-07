import { useEffect, useState } from "react";
import axios from "axios";

export default function Products() {
  const [products, setProducts] = useState([]);
  useEffect(() => {
    axios.get("http://localhost:8080/products")
      .then(res => setProducts(res.data))
      .catch(() => setProducts([]));
  }, []);

  return (
    <div className="p-8">
      <h2 className="text-2xl font-bold mb-4">Products</h2>
      <ul>
        {products.map(p => (
          <li key={p.id} className="mb-2">
            <span className="font-semibold">{p.name}</span> - ₹{p.price}
          </li>
        ))}
      </ul>
    </div>
  );
}
