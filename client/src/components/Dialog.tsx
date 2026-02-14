import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Label } from "@/components/ui/label";
import { useState } from "react";
import { PriceInput } from "./PriceInput";
import { createLimit } from "@/services/limit";
import { useMutation, useQueryClient } from "@tanstack/react-query";

function Modals() {
  const [rawValue, setRawValue] = useState("");
  const [open, setOpen] = useState(false);
  const queryClient = useQueryClient();
  const mutation = useMutation({
    mutationFn: createLimit,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["limits"] });
      setOpen(false);
    },
  });
  return (
    <Dialog open={open}>
      <form>
        <DialogTrigger asChild>
          <Button
            variant="outline"
            className="cursor-pointer"
            onClick={() => setOpen(true)}
          >
            Ajukan Limit
          </Button>
        </DialogTrigger>
        <DialogContent className="sm:max-w-106.25">
          <DialogHeader>
            <DialogTitle>Pengajuan Limit</DialogTitle>
            <DialogDescription>
              Lengkapi informasi yang dibutuhkan untuk pengajuan limit.
            </DialogDescription>
          </DialogHeader>
          <div className="grid gap-4">
            <div className="grid gap-3">
              <Label htmlFor="nominal">Nominal Limit (Rp)</Label>
              <PriceInput onChange={setRawValue} />
            </div>
          </div>
          <DialogFooter>
            <DialogClose asChild>
              <Button onClick={() => setOpen(false)} variant="outline">
                Batal
              </Button>
            </DialogClose>
            <Button
              type="submit"
              onClick={() => mutation.mutate(Number(rawValue))}
            >
              Ajukan
            </Button>
          </DialogFooter>
        </DialogContent>
      </form>
    </Dialog>
  );
}

export default Modals;
