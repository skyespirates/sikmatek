import { NumericFormat } from "react-number-format";
import { Input } from "@/components/ui/input"; // adjust import path to your shadcn setup

export function PriceInput({ onChange }: { onChange: any }) {
  return (
    <NumericFormat
      customInput={Input}
      thousandSeparator="."
      decimalSeparator=","
      allowNegative={false}
      onValueChange={(values) => {
        onChange(values.value);
      }}
    />
  );
}
