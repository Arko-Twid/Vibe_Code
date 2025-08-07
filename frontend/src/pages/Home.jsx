import { Link } from "react-router-dom";

export default function Home() {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen">
      <h1 className="text-3xl font-bold mb-4">Welcome to Vibe-Code Stationary Marketplace</h1>
      <Link to="/products" className="text-blue-500 underline">Browse Products</Link>
    </div>
  );
}
