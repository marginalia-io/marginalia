import { createRootRoute, Outlet, redirect, useMatchRoute } from "@tanstack/react-router";
import { TanStackRouterDevtools } from "@tanstack/react-router-devtools";
import { ReactQueryDevtools } from "@tanstack/react-query-devtools";

import { Navbar } from "@/components/ui/navbar";
import { fetchSetupStatus } from "@/queries/setup";

const RootLayout = () => {
  const matchRoute = useMatchRoute();
  const isSetupRoute = matchRoute({ to: "/setup" });

  return (
    <>
      {!isSetupRoute ? (
        <>
          <Navbar />
        </>
      ) : (
        <></>
      )}
      <Outlet />
      <ReactQueryDevtools buttonPosition="bottom-right" />
      <TanStackRouterDevtools position="bottom-left" />
    </>
  );
};

export const Route = createRootRoute({
  component: RootLayout,

  beforeLoad: async () => {
    const setupStatus = await fetchSetupStatus();

    if (!setupStatus?.completed && !window.location.pathname.startsWith("/setup")) {
      throw redirect({ to: "/setup" });
    }
  },
});
