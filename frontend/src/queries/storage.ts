import { queryOptions } from "@tanstack/react-query";

export type StorageInfo = {
  path: string;
  available_bytes: number;
  total_bytes: number;
};

export const fetchStorage = async (): Promise<StorageInfo> => {
  const response = await fetch("/api/storage");
  if (!response.ok) {
    throw new Error("Failed to fetch storage info");
  }
  return response.json();
};

export const storageQuery = queryOptions({
  queryKey: ["storage"],
  queryFn: fetchStorage,
});

export function formatBytes(bytes: number): string {
  const units = ["B", "KB", "MB", "GB", "TB"];
  let value = bytes;
  let unitIndex = 0;

  while (value >= 1024 && unitIndex < units.length - 1) {
    value /= 1024;
    unitIndex++;
  }

  const decimals = unitIndex === 0 ? 0 : 1;
  return `${value.toFixed(decimals)} ${units[unitIndex]}`;
}
