import classNames from "classnames";

interface LogoProps {
  mode: "light" | "dark";
}

export function Logo({ mode }: LogoProps) {
  return (
    <h1
      className={classNames(
        "flex items-center gap-2 text-2xl font-serif italic",
        mode === "light" ? "text-ink" : "text-white",
      )}
    >
      <span className="text-primary not-italic">&para;</span>
      Marginalia
    </h1>
  );
}
