import { Check, Database, HardDrive, Loader2, Shield, X } from "lucide-react";
import { createFileRoute } from "@tanstack/react-router";
import { useQuery } from "@tanstack/react-query";

import { Button } from "@/components/ui/button";
import { Card, CardContent, CardFooter, CardHeader, CardTitle } from "@/components/ui/card";
import { Eyebrow } from "@/components/ui/eyebrow";
import { Tooltip, TooltipContent, TooltipTrigger } from "@/components/ui/tooltip";
import { formatBytes, storageQuery } from "@/queries/storage";

export const Route = createFileRoute("/onboarding/")({
  component: Index,
});

function Index() {
  const storage = useQuery(storageQuery);

  return (
    <Card className="w-full max-w-lg gap-y-4 pb-0">
      <CardHeader>
        <CardTitle className="flex flex-col items-start gap-1">
          <Eyebrow variant="primary">First-time setup</Eyebrow>
          <h1 className="text-3xl font-medium font-serif">Welcome to your new library</h1>
        </CardTitle>
      </CardHeader>

      <CardContent className="flex flex-col gap-y-4">
        <p className="text-muted-foreground">
          Marginalia will help you build a library of your favorite books. Let's get you set up.
        </p>

        <ul className="border border-border divide-y divide-border rounded-md">
          <li className="flex items-center gap-x-3 px-4 py-3 text-xs text-muted-foreground">
            <Shield className="size-4 text-muted-foreground/75" />
            <Eyebrow variant="default">System check</Eyebrow>
          </li>

          <li className="flex items-center gap-x-3 px-4 py-3 text-xs text-muted-foreground">
            <Database className="size-4 text-muted-foreground/75" />

            <div className="flex flex-col items-start gap-y-0.5">
              <span className="text-sm font-semibold text-foreground">Database</span>
              <span className="text-xxs text-muted-foreground/75 font-mono">
                SQLite &middot; connected
              </span>
            </div>

            <div className="ml-auto">
              <span className="flex items-center gap-x-1.5 bg-chart-2/10 rounded-full">
                <Check className="size-4 text-chart-2 m-1" />
              </span>
            </div>
          </li>

          <li className="flex items-center gap-x-3 px-4 py-3 text-xs text-muted-foreground">
            <HardDrive className="size-4 text-muted-foreground/75" />

            <div className="flex flex-col items-start gap-y-0.5 min-w-0">
              <span className="text-sm font-semibold text-foreground">Storage</span>
              {storage.isPending && (
                <span className="text-xxs text-muted-foreground/75 font-mono">
                  Checking storage&hellip;
                </span>
              )}
              {storage.isError && (
                <span className="text-xxs text-destructive font-mono">Unable to read storage</span>
              )}
              {storage.isSuccess && (
                <span className="text-xxs text-muted-foreground/75 font-mono flex items-center min-w-0 w-full max-w-full">
                  <Tooltip>
                    <TooltipTrigger asChild>
                      <span className="min-w-0 flex-1 truncate">{storage.data.path}</span>
                    </TooltipTrigger>
                    <TooltipContent side="top" className="font-mono max-w-sm break-all">
                      {storage.data.path}
                    </TooltipContent>
                  </Tooltip>
                  <span className="shrink-0">
                    &nbsp;&middot; {formatBytes(storage.data.available_bytes)} available
                  </span>
                </span>
              )}
            </div>

            <div className="ml-auto shrink-0">
              {storage.isPending && (
                <span className="flex items-center gap-x-1.5 bg-muted rounded-full">
                  <Loader2 className="size-4 text-muted-foreground m-1 animate-spin" />
                </span>
              )}
              {storage.isError && (
                <span className="flex items-center gap-x-1.5 bg-destructive/10 rounded-full">
                  <X className="size-4 text-destructive m-1" />
                </span>
              )}
              {storage.isSuccess && (
                <span className="flex items-center gap-x-1.5 bg-chart-2/10 rounded-full">
                  <Check className="size-4 text-chart-2 m-1" />
                </span>
              )}
            </div>
          </li>
        </ul>
      </CardContent>

      <CardFooter className="flex justify-end bg-muted border-t border-border py-3! mt-2">
        <Button>Continue</Button>
      </CardFooter>
    </Card>
  );
}
