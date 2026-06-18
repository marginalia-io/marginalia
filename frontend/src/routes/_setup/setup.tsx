import { createFileRoute } from "@tanstack/react-router";

import { Card, CardTitle, CardHeader, CardDescription, CardContent } from "@/components/ui/card";

export const Route = createFileRoute("/_setup/setup")({
  component: Setup,
});

function Setup() {
  return (
    <div className="p-2">
      <Card>
        <CardHeader>
          <CardTitle>Setup</CardTitle>
          <CardDescription>Welcome to the Marginalia onboarding process.</CardDescription>
        </CardHeader>
        <CardContent>
          <p>Please follow the instructions to setup your Marginalia account.</p>
        </CardContent>
      </Card>
    </div>
  );
}
