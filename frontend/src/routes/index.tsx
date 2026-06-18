import { createFileRoute } from "@tanstack/react-router";

import { Button } from "@/components/ui/button";
import { Card, CardTitle, CardHeader, CardDescription, CardContent } from "@/components/ui/card";

export const Route = createFileRoute("/")({
  component: Index,
});

function Index() {
  return (
    <div className="p-2">
      <Card>
        <CardHeader>
          <CardTitle>Hello, I'm Marginalia</CardTitle>
          <CardDescription>Welcome to the Marginalia home page.</CardDescription>
        </CardHeader>
        <CardContent>
          <p>I'm a chatbot that can help you with your questions.</p>
        </CardContent>
      </Card>

      <Button>Start a new conversation</Button>
      <Button variant="secondary">View my history</Button>
      <Button variant="ghost">View my settings</Button>
      <Button variant="destructive">View my profile</Button>
      <Button variant="link">View my notifications</Button>
    </div>
  );
}
