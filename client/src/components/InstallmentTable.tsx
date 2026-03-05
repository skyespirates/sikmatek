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
import { Button } from "@/components/ui/button";
import { format } from "date-fns";
import { useMutation } from "@tanstack/react-query";
import { payInstallment } from "@/services/installment";
import { queryClient } from "@/main";
import { toast } from "sonner";
import { Badge } from "./ui/badge";

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
  installments: Installment[];
};

const InstallmentTable = ({ installments }: Props) => {
  const mutation = useMutation({
    mutationFn: payInstallment,
    onSuccess: (data) => {
      toast(data.message, {
        position: "top-center",
        closeButton: true,
      });
      queryClient.invalidateQueries({
        queryKey: ["installments", installments[0].nomor_kontrak],
      });
    },
  });

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
        {!installments ? (
          <TableRow>
            <TableCell
              colSpan={7}
              className="h-24 text-center text-muted-foreground"
            >
              Belum ada cicilan yang dibuat.
            </TableCell>
          </TableRow>
        ) : (
          installments.map((i) => (
            <TableRow key={i.id}>
              <TableCell className="font-medium">{i.nomor_kontrak}</TableCell>
              <TableCell className="text-center">{i.bulan_ke}</TableCell>
              <TableCell>{formatNominal(i.jumlah_tagihan)}</TableCell>
              <TableCell>{format(i.due_date, "dd-MMM-yyyy")}</TableCell>
              <TableCell>
                <Badge
                  className={
                    i.status === "UNPAID"
                      ? "bg-yellow-50 text-yellow-700 dark:bg-yellow-950 dark:text-yellow-300"
                      : "bg-green-50 text-green-700 dark:bg-green-950 dark:text-green-300"
                  }
                >
                  {i.status.toLocaleLowerCase()}
                </Badge>
              </TableCell>
              <TableCell className="text-center">
                {(i.paid_at &&
                  format(new Date(i.paid_at), "dd-MMM-yyyy 'at' HH:mm")) ||
                  "-"}
              </TableCell>
              <TableCell className="text-right">
                <Button
                  disabled={i.paid_at ? true : false}
                  className="cursor-pointer"
                  onClick={() => mutation.mutate(i.id)}
                >
                  Bayar
                </Button>
              </TableCell>
            </TableRow>
          ))
        )}
      </TableBody>
    </Table>
  );
};

export default InstallmentTable;
