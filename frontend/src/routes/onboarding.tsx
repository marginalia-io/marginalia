import { createFileRoute, Outlet } from "@tanstack/react-router";

import { Logo } from "@/components/ui/logo";

export const Route = createFileRoute("/onboarding")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <div className="flex flex-col items-center justify-start gap-10 mt-20 h-screen">
      <Logo mode="light" />

      <Outlet />
    </div>
  );
}
