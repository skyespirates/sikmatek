import {
  Table as T,
  TableCaption,
  TableHeader,
  TableRow,
  TableHead,
  TableBody,
  TableCell,
} from "@/components/ui/table";

type TableProps = {
  data: any;
};

const Table = (props: TableProps) => {
  return (
    <T>
      <TableCaption>A list of your recent invoices.</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead className="w-[100px]">ID</TableHead>
          <TableHead>Durasi (bulan)</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {props.data?.map((dt: any) => (
          <TableRow key={dt.id}>
            <TableCell className="font-medium">{dt.id}</TableCell>
            <TableCell>{dt.durasi_bulan}</TableCell>
          </TableRow>
        ))}
      </TableBody>
    </T>
  );
};

export default Table;
