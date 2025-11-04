import { AuthGuard } from "@/components/auth/AuthGuard";
import { Feed } from "@/components/feed/Feed";

export default function Home() {
  return (
    <AuthGuard redirectTo="/landing">
      <Feed />
    </AuthGuard>
  );
}
