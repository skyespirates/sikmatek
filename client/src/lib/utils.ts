import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function formatNominal(nominal: number): string {
  return new Intl.NumberFormat("id-ID").format(nominal);
}
