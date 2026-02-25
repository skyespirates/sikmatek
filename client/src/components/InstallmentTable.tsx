import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { formatNominal } from "@/lib/utils";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "@/components/ui/button";
import { MoreHorizontalIcon } from "lucide-react";
import { format } from "date-fns";

type Installment = {
  id: number;
  nomor_kontrak: string;
  bulan_ke: number;
  jumlah_tagihan: number;
  due_date: string;
  status: string;
  paid_at?: string;
};

const data: Installment[] = [
  {
    id: 1,
    nomor_kontrak: "WG-20260212-8af7fcad",
    bulan_ke: 1,
    jumlah_tagihan: 25000000,
    due_date: "2026-02-24",
    status: "UNPAID",
    paid_at: "2026-02-24",
  },
];

type Props = {
  installments: Installment[];
};

const InstallmentTable = ({ installments }: Props) => {
  return (
    <Table>
      <TableCaption>A list of your installments.</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead className="w-25">Nomor Kontrak</TableHead>
          <TableHead className="text-center">Bulan Ke</TableHead>
          <TableHead>Jumlah Tagihan</TableHead>
          <TableHead>Due Date</TableHead>
          <TableHead>Status</TableHead>
          <TableHead className="text-center">Paid At</TableHead>
          <TableHead className="text-right">Actions</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {installments.map((i) => (
          <TableRow key={i.id}>
            <TableCell className="font-medium">{i.nomor_kontrak}</TableCell>
            <TableCell className="text-center">{i.bulan_ke}</TableCell>
            <TableCell>{formatNominal(i.jumlah_tagihan)}</TableCell>
            <TableCell>{format(i.due_date, "dd-MMM-yyyy")}</TableCell>
            <TableCell>{i.status}</TableCell>
            <TableCell className="text-center">{i.paid_at || "-"}</TableCell>
            <TableCell className="text-right">
              <Button className="cursor-pointer">Bayar</Button>
            </TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
};

export default InstallmentTable;
