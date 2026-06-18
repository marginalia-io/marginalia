import { Button } from "./button";
import { Logo } from "./logo";
import { Link } from "@tanstack/react-router";

export function Navbar() {
  return (
    <nav className="flex items-center justify-between">
      <Logo mode="light" />
      <Button>
        <Link to="/setup" className="[&.active]:font-bold">
          Setup
        </Link>
      </Button>
    </nav>
  );
}
