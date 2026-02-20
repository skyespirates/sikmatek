import { Card, CardContent } from "@/components/ui/card";
import {
  Field,
  FieldGroup,
  FieldLabel,
  FieldError,
} from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { useEffect, useRef, useState } from "react";
import { Button } from "@/components/ui/button";
import { Calendar } from "@/components/ui/calendar";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";

import * as z from "zod";
import { Controller, useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { RotateCcw, Save } from "lucide-react";
import { useMutation, useQuery } from "@tanstack/react-query";
import {
  getProfile,
  updateProfile,
  uploadKtp,
  uploadSelfie,
} from "@/services/profile";
import { toast } from "sonner";

export const formSchema = z.object({
  nik: z
    .string()
    .min(16, "nik harus 16 karakter")
    .max(16, "nik harus 16 karakter"),
  full_name: z
    .string()
    .min(4, "nama lengkap minimal 4 karakter")
    .max(36, "nama lengkap maksimal 36 karakter"),
  legal_name: z
    .string()
    .min(4, "nama sesuai ktp minimal 4 karakter")
    .max(24, "nama sesuai ktp maksimal 24 karakter"),
  tempat_lahir: z
    .string()
    .min(3, "tempat lahir minimal 3 karakter")
    .max(24, "tempat lahir maksimal 24 karakter"),
  tanggal_lahir: z.date(),
  gaji: z.number().min(0, "gaji harus bilangan positif"),
});

const Profile = () => {
  const [fileKtp, setFileKtp] = useState<File | null>(null);
  const [fileSelfie, setFileSelfie] = useState<File | null>(null);
  const [open, setOpen] = useState(false);
  const ktpRef = useRef<HTMLInputElement>(null);
  const selfieRef = useRef<HTMLInputElement>(null);

  const { data } = useQuery({
    queryKey: ["profile"],
    queryFn: getProfile,
  });
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      nik: "",
      full_name: "",
      legal_name: "",
      tempat_lahir: "",
      tanggal_lahir: new Date(),
      gaji: 0,
    },
  });

  const mutation = useMutation({
    mutationFn: updateProfile,
    onSuccess: (data) => {
      toast(data.message, { position: "top-right", closeButton: true });
    },
  });

  const ktpMutation = useMutation({
    mutationFn: uploadKtp,
    onSuccess: (data) => {
      setFileKtp(null);
      if (ktpRef.current) {
        ktpRef.current.value = "";
      }

      toast(data.message, { position: "top-right", closeButton: true });
    },
  });

  const selfieMutation = useMutation({
    mutationFn: uploadSelfie,
    onSuccess: (data) => {
      setFileSelfie(null);
      if (selfieRef.current) {
        selfieRef.current.value = "";
      }
      toast(data.message, { position: "top-right", closeButton: true });
    },
  });

  function onSubmit(data: z.infer<typeof formSchema>) {
    // console.log(data);
    mutation.mutate(data);
  }
  useEffect(() => {
    if (data) {
      form.reset({
        nik: data.nik,
        full_name: data.full_name,
        legal_name: data.legal_name,
        tempat_lahir: data.tempat_lahir,
        // convert string from API to Date object
        tanggal_lahir: new Date(data.tanggal_lahir),
        gaji: data.gaji,
      });
    }
  }, [data, form]);

  function handleUploadKtp() {
    if (!fileKtp) {
      return;
    }
    const ktpForm = new FormData();
    ktpForm.append("image", fileKtp);

    ktpMutation.mutate(ktpForm);
  }

  function handleUploadSelfie() {
    if (!fileSelfie) {
      return;
    }
    const selfieForm = new FormData();
    selfieForm.append("image", fileSelfie);

    selfieMutation.mutate(selfieForm);
  }

  return (
    <div>
      <Card>
        <CardContent>
          <form onSubmit={form.handleSubmit(onSubmit)}>
            <FieldGroup className="mb-6">
              <Controller
                name="nik"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field className="w-72" data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor="nik">NIK</FieldLabel>
                    <Input
                      {...field}
                      id="nik"
                      aria-invalid={fieldState.invalid}
                      placeholder="3273056010900009"
                      autoComplete="off"
                    />
                    {fieldState.invalid && (
                      <FieldError errors={[fieldState.error]} />
                    )}
                  </Field>
                )}
              />
              <Controller
                name="full_name"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field className="w-72" data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor="full-name">Nama Lengkap</FieldLabel>
                    <Input
                      {...field}
                      id="full-name"
                      aria-invalid={fieldState.invalid}
                      placeholder="John Doe"
                      autoComplete="off"
                    />
                    {fieldState.invalid && (
                      <FieldError errors={[fieldState.error]} />
                    )}
                  </Field>
                )}
              />
              <Controller
                name="legal_name"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field className="w-72" data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor="legal-name">
                      Nama Sesuai KTP
                    </FieldLabel>
                    <Input
                      {...field}
                      id="legal-name"
                      aria-invalid={fieldState.invalid}
                      placeholder="John D"
                      autoComplete="off"
                    />
                    {fieldState.invalid && (
                      <FieldError errors={[fieldState.error]} />
                    )}
                  </Field>
                )}
              />
              <Controller
                name="tempat_lahir"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field className="w-72" data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor="tempat-lahir">Tempat Lahir</FieldLabel>
                    <Input
                      {...field}
                      id="tempat-lahir"
                      aria-invalid={fieldState.invalid}
                      placeholder="Jakarta"
                      autoComplete="off"
                    />
                    {fieldState.invalid && (
                      <FieldError errors={[fieldState.error]} />
                    )}
                  </Field>
                )}
              />
              <Controller
                name="tanggal_lahir"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field className="w-72" data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor="tanggal-lahir">
                      Tanggal Lahir
                    </FieldLabel>
                    <Popover open={open} onOpenChange={setOpen}>
                      <PopoverTrigger asChild>
                        <Button
                          variant="outline"
                          id="date"
                          className="justify-start font-normal"
                        >
                          {field.value
                            ? field.value.toLocaleDateString()
                            : "Select date"}
                        </Button>
                      </PopoverTrigger>
                      <PopoverContent
                        className="w-auto overflow-hidden p-0"
                        align="start"
                      >
                        <Calendar
                          mode="single"
                          selected={field.value}
                          defaultMonth={field.value}
                          captionLayout="dropdown"
                          onSelect={(date) => {
                            field.onChange(date); // ✅ update RHF state
                            setOpen(false);
                          }}
                        />
                      </PopoverContent>
                    </Popover>
                    {fieldState.invalid && (
                      <FieldError errors={[fieldState.error]} />
                    )}
                  </Field>
                )}
              />
              <Controller
                name="gaji"
                control={form.control}
                render={({ field, fieldState }) => (
                  <Field className="w-72" data-invalid={fieldState.invalid}>
                    <FieldLabel htmlFor="gaji">Gaji (Rp)</FieldLabel>
                    <Input
                      type="number"
                      {...field}
                      value={field.value ?? ""}
                      onChange={(e) => field.onChange(e.target.valueAsNumber)}
                      id="gaji"
                      aria-invalid={fieldState.invalid}
                      autoComplete="off"
                    />
                    {fieldState.invalid && (
                      <FieldError errors={[fieldState.error]} />
                    )}
                  </Field>
                )}
              />
            </FieldGroup>
            <div className="w-72 flex gap-2 justify-end">
              <Button
                variant="outline"
                className="cursor-pointer"
                onClick={() => form.reset()}
              >
                <RotateCcw /> Reset
              </Button>
              <Button className="cursor-pointer">
                <Save /> Simpan
              </Button>
            </div>
          </form>

          <Field className="w-72 my-6">
            <FieldLabel htmlFor="ktp">Upload KTP</FieldLabel>
            <div className="flex gap-2">
              <Input
                className="cursor-pointer"
                ref={ktpRef}
                id="ktp"
                type="file"
                onChange={(e) =>
                  e.target.files && setFileKtp(e.target.files[0])
                }
              />
              <Button className="cursor-pointer" onClick={handleUploadKtp}>
                <Save />
              </Button>
            </div>
          </Field>
          <Field className="w-72">
            <FieldLabel htmlFor="selfie">Upload Selfie</FieldLabel>
            <div className="flex gap-2">
              <Input
                className="cursor-pointer"
                ref={selfieRef}
                id="selfie"
                type="file"
                onChange={(e) =>
                  e.target.files && setFileSelfie(e.target.files[0])
                }
              />
              <Button className="cursor-pointer" onClick={handleUploadSelfie}>
                <Save />
              </Button>
            </div>
          </Field>
        </CardContent>
      </Card>
    </div>
  );
};

export default Profile;
