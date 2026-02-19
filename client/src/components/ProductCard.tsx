import { Badge } from "@/components/ui/badge";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardAction,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { formatNominal } from "@/lib/utils";

type Product = {
  id: number;
  nama_produk: string;
  kategori: string;
  harga: number;
};

type Props = {
  product: Product;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  setSelected: React.Dispatch<React.SetStateAction<Product>>;
};

const ProductCard = ({ product, setOpen, setSelected }: Props) => {
  function handleClick() {
    setSelected(product);
    setOpen(true);
  }

  return (
    <Card className="relative mx-auto w-full max-w-sm pt-0">
      <div className="absolute inset-0 z-30 aspect-video bg-black/35 " />
      <img
        src="https://www.bhinneka.com/blog/wp-content/uploads/2022/03/lemari-es-2-pintu-terbaik-aqua-222x500.jpg"
        alt="Event cover"
        className="relative z-20 aspect-video w-full object-contain brightness-60 grayscale dark:brightness-40"
      />
      <CardHeader>
        <CardAction>
          <Badge variant="secondary">Rp. {formatNominal(product.harga)}</Badge>
        </CardAction>
        <CardTitle>{product.nama_produk}</CardTitle>
        <CardDescription>{product.kategori}</CardDescription>
      </CardHeader>
      <CardFooter>
        <Button onClick={handleClick} className="w-full cursor-pointer">
          Buat Kontrak
        </Button>
      </CardFooter>
    </Card>
  );
};

export default ProductCard;
