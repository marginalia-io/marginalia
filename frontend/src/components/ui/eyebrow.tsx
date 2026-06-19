import classNames from "classnames";

interface EyebrowProps {
  children: React.ReactNode;
  className?: string;
  variant?: "default" | "primary" | "secondary";
}

export function Eyebrow({ children, className, variant = "default" }: EyebrowProps) {
  return (
    <span
      className={classNames(
        "text-xxs font-medium uppercase tracking-wider text-muted-foreground",
        variant === "primary" && "text-primary",
        variant === "secondary" && "text-secondary",
        className,
      )}
    >
      {children}
    </span>
  );
}
