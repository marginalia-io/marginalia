import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/_setup")({
  component: RouteComponent,
});

function RouteComponent() {
  // return <Outlet />;
  return <div>Hello "/_setup"!</div>;
}
