import { createRootRoute, Outlet, redirect, useMatchRoute } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";

import { Navbar } from "@/components/ui/navbar";
import { TooltipProvider } from "@/components/ui/tooltip";
import { fetchSetupStatus } from "@/queries/setup";

const RootLayout = () => {
  const matchRoute = useMatchRoute();
  const isOnboardingRoute = Boolean(matchRoute({ to: "/onboarding" }));

  return (
    <TooltipProvider>
      {!isOnboardingRoute && <Navbar />}
      <Outlet />
      <ReactQueryDevtools buttonPosition="bottom-right" />
      <TanStackRouterDevtools position="bottom-left" />
    </TooltipProvider>
  );
};

export const Route = createRootRoute({
  component: RootLayout,

  beforeLoad: async ({ location }) => {
    const setupStatus = await fetchSetupStatus();
    const isOnboardingRoute = location.pathname.startsWith("/onboarding");

    if (!setupStatus?.completed && !isOnboardingRoute) {
      throw redirect({ to: "/onboarding" });
    }
  },
});
