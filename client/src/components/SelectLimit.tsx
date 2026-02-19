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

type Props = {
  items: Item[];
  value: string;
  onChange: React.Dispatch<React.SetStateAction<string>>;
};

export default function SelectLimit({ items, value, onChange }: Props) {
  return (
    <Select value={value} onValueChange={onChange}>
      <SelectTrigger className="">
        <SelectValue placeholder="Pilih limit" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          <SelectLabel>Limit aktif</SelectLabel>
          {items.map((item) => (
            <SelectItem key={item.value} value={item.value}>
              {item.label}
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  );
}
