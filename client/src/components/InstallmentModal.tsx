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
import InstallmentTable from "./InstallmentTable";

type Installment = {
  id: number;
  nomor_kontrak: string;
  bulan_ke: number;
  jumlah_tagihan: number;
  due_date: string;
  status: string;
  paid_at?: string;
};

type Props = {
  open: boolean;
  setOpen: React.Dispatch<React.SetStateAction<boolean>>;
  installments: Installment[];
};

const InstallmentModal = ({ open, setOpen, installments }: Props) => {
  return (
    <Dialog open={open} onOpenChange={setOpen}>
      <DialogContent className="min-w-6xl">
        <DialogHeader>
          <DialogTitle>Daftar Cicilan</DialogTitle>
          <DialogDescription>
            Harap bayar cicilan tepat waktu 👌.
          </DialogDescription>
        </DialogHeader>
        <InstallmentTable installments={installments} />
        <DialogFooter>
          <DialogClose asChild>
            <Button variant="outline">Close</Button>
          </DialogClose>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
};

export default InstallmentModal;
