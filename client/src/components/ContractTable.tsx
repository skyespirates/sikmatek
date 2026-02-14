import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Badge } from "./ui/badge";
import { formatNominal } from "@/lib/utils";
import { Card, CardContent } from "./ui/card";

type Contract = {
  nomor_kontrak: string;
  otr: number;
  admin_fee: number;
  jumlah_bunga: number;
  tenor: number;
  total_pembiayaan: number;
  status: string;
  consumer_id: number;
  product_id: number;
  limit_id: number;
};

export function ContractTable({ contracts }: { contracts: Contract[] }) {
  return (
    <Card>
      <CardContent>
        <Table>
          <TableCaption>A list of your recent contracts.</TableCaption>
          <TableHeader>
            <TableRow>
              <TableHead className="w-25">Nomor Kontrak</TableHead>
              <TableHead>OTR (Rp)</TableHead>
              <TableHead>Admin Fee (Rp)</TableHead>
              <TableHead>Jumlah Bunga (Rp)</TableHead>
              <TableHead>Tenor (Bulan)</TableHead>
              <TableHead>Total Pembiayaan (Rp)</TableHead>
              <TableHead>Status</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {contracts.map((contract) => (
              <TableRow key={contract.nomor_kontrak}>
                <TableCell className="font-medium">
                  {contract.nomor_kontrak}
                </TableCell>
                <TableCell>{formatNominal(contract.otr)}</TableCell>
                <TableCell>{formatNominal(contract.admin_fee)}</TableCell>
                <TableCell>{formatNominal(contract.jumlah_bunga)}</TableCell>
                <TableCell>{contract.tenor}</TableCell>
                <TableCell>
                  {formatNominal(contract.total_pembiayaan)}
                </TableCell>
                <TableCell>
                  <Badge className={`${status[contract.status as statusKey]}`}>
                    {contract.status.toLowerCase()}
                  </Badge>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </CardContent>
    </Card>
  );
}

type statusKey =
  | "DRAFT"
  | "QUOTED"
  | "CANCELLED"
  | "CONFIRMED"
  | "ACTIVE"
  | "COMPLETED"
  | "DEFAULT";

const status = {
  DRAFT: "bg-grey-50 text-grey-700 dark:bg-grey-950 dark:text-grey-300",
  QUOTED:
    "bg-yellow-50 text-yellow-700 dark:bg-yellow-950 dark:text-yellow-300",
  CANCELLED: "bg-red-50 text-red-700 dark:bg-red-950 dark:text-red-300",
  CONFIRMED: "bg-blue-50 text-blue-700 dark:bg-blue-950 dark:text-blue-300",
  ACTIVE: "bg-green-50 text-green-700 dark:bg-green-950 dark:text-green-300",
  COMPLETED: "bg-green-50 text-green-700 dark:bg-green-950 dark:text-green-300",
  DEFAULT: "bg-grey-50 text-grey-700 dark:bg-grey-950 dark:text-grey-300",
} as const;
