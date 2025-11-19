import { AuthGuard } from "@/components/auth/AuthGuard";
import { Feed } from "@/components/feed/Feed";

export default function FeedPage() {
  return (
    <AuthGuard redirectTo="/landing">
      <Feed />
    </AuthGuard>
  );
}
