import { queryOptions } from "@tanstack/react-query";

export const fetchSetupStatus = async () => {
  const response = await fetch("/api/setup");
  return response.json();
};

export const setupStatusQuery = queryOptions({
  queryKey: ["setup-status"],
  queryFn: fetchSetupStatus,
});
