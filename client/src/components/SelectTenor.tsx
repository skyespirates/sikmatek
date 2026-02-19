import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";

type Item = {
  value: string;
  label: string;
};

const tenorOptions: Item[] = [
  {
    value: "3",
    label: "3",
  },
  {
    value: "4",
    label: "4",
  },
  {
    value: "6",
    label: "6",
  },
];

type Props = {
  value: string;
  onChange: React.Dispatch<React.SetStateAction<string>>;
};

export default function SelectTenor({ value, onChange }: Props) {
  return (
    <Select value={value} onValueChange={onChange}>
      <SelectTrigger className="">
        <SelectValue placeholder="Pilih Tenor (bulan)" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          <SelectLabel>Tenor</SelectLabel>
          {tenorOptions.map((item) => (
            <SelectItem key={item.value} value={item.value}>
              {item.label}
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  );
}
