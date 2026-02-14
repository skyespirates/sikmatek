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
import Dialog from "./Dialog";

type Limit = {
  id: number;
  requested_limit: number;
  used_limit: number;
  remaining_limit: number;
  status: string;
  consumer_id: number;
};

export function LimitTable({ limits }: { limits: Limit[] }) {
  return (
    <Card>
      <CardContent>
        <Dialog />
        <Table>
          <TableCaption>A list of your recent limits.</TableCaption>
          <TableHeader>
            <TableRow>
              <TableHead className="w-25">ID</TableHead>
              <TableHead>Requested</TableHead>
              <TableHead>Used</TableHead>
              <TableHead>Remaining</TableHead>
              <TableHead>Status</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {limits.map((limit) => (
              <TableRow key={limit.id}>
                <TableCell className="font-medium">{limit.id}</TableCell>
                <TableCell>{formatNominal(limit.requested_limit)}</TableCell>
                <TableCell>{formatNominal(limit.used_limit)}</TableCell>
                <TableCell>{formatNominal(limit.remaining_limit)}</TableCell>
                <TableCell>
                  <Badge className={`${status[limit.status as statusKey]}`}>
                    {limit.status.toLowerCase()}
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

type statusKey = "PENDING" | "REJECTED" | "APPROVED";

const status = {
  PENDING:
    "bg-yellow-50 text-yellow-700 dark:bg-yellow-950 dark:text-yellow-300",
  REJECTED: "bg-red-50 text-red-700 dark:bg-red-950 dark:text-red-300",
  APPROVED: "bg-green-50 text-green-700 dark:bg-green-950 dark:text-green-300",
} as const;
