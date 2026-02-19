import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog";
import { Field } from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useState } from "react";
import { getApprovedLimits } from "@/services/limit";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { formatNominal } from "@/lib/utils";
import SelectLimit from "./SelectLimit";
import SelectTenor from "./SelectTenor";
import { buatKontrak } from "@/services/contract";

type Product = {
  id: number;
  nama_produk: string;
  kategori: string;
  harga: number;
};

type Props = {
  product: Product;
  open: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
};

type KontrakData = {
  product_id: number;
  limit_id: number;
  tenor: number;
};

const CreateContractForm = ({ product, open, setOpen }: Props) => {
  const [tenor, setTenor] = useState("");
  const [limit, setLimit] = useState("");
  const { data, isSuccess } = useQuery({
    queryKey: ["approved"],
    queryFn: getApprovedLimits,
  });
  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: buatKontrak,
    onSuccess: (data) => {
      queryClient.invalidateQueries({ queryKey: ["kontrak"] });
      console.log(data);
      setOpen(false);
    },
  });
  return (
    <Dialog open={open}>
      <form>
        <DialogContent className="sm:max-w-106.25">
          <DialogHeader>
            <DialogTitle>Buat Kontrak</DialogTitle>
            <DialogDescription>
              Lengkapi informasi yang dibutuhkan untuk membuat kontrak.
            </DialogDescription>
          </DialogHeader>
          <div className="grid gap-4">
            <Field data-disabled>
              <Label>Nama Produk</Label>
              <Input
                placeholder="Nama Produk"
                value={product.nama_produk}
                disabled
              />
            </Field>

            <Field data-disabled>
              <Label>Harga (Rp)</Label>
              <Input
                placeholder="Harga"
                value={formatNominal(product.harga)}
                disabled
              />
            </Field>
            <Field data-disabled>
              <Label>Tenor</Label>
              <SelectTenor value={tenor} onChange={setTenor} />
            </Field>
            <Field data-disabled>
              <Label>Limit</Label>
              {isSuccess && Array.isArray(data) && (
                <SelectLimit items={data} value={limit} onChange={setLimit} />
              )}
            </Field>
          </div>
          <DialogFooter>
            <DialogClose asChild>
              <Button onClick={() => setOpen(false)} variant="outline">
                Batal
              </Button>
            </DialogClose>
            <Button
              type="submit"
              onClick={() => {
                const payload: KontrakData = {
                  product_id: product.id,
                  limit_id: Number(limit),
                  tenor: Number(tenor),
                };
                mutation.mutate(payload);
              }}
            >
              Buat
            </Button>
          </DialogFooter>
        </DialogContent>
      </form>
    </Dialog>
  );
};

export default CreateContractForm;
