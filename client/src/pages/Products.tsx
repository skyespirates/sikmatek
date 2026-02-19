import CreateContractForm from "@/components/CreateContractForm";
import ProductCard from "@/components/ProductCard";
import { getProductList } from "@/services/product";
import { useQuery } from "@tanstack/react-query";
import { useState } from "react";

type Product = {
  id: number;
  nama_produk: string;
  kategori: string;
  harga: number;
};

const Products = () => {
  const [open, setOpen] = useState(false);
  const [selected, setSelected] = useState<Product>({} as Product);
  const { data } = useQuery({
    queryKey: ["products"],
    queryFn: getProductList,
  });
  return (
    <div>
      Products
      <div className="grid grid-cols-3 gap-4">
        {data &&
          Array.isArray(data) &&
          data.map((p) => (
            <ProductCard
              key={p.id}
              product={p}
              setOpen={setOpen}
              setSelected={setSelected}
            />
          ))}
        <CreateContractForm open={open} setOpen={setOpen} product={selected} />
      </div>
    </div>
  );
};

export default Products;
