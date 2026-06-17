import { Button } from "@/components/ui/button";

export default function App() {
  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <Button variant="default">Click me</Button>
      <Button variant="secondary">Click me</Button>
      <Button variant="ghost">Click me</Button>
      <Button variant="destructive">Click me</Button>
      <Button variant="link">Click me</Button>
    </div>
  );
}
