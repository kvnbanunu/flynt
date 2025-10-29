import { AuthGuard } from "@/components/auth/AuthGuard";
import { FriendsComponent } from "@/components/friends/FriendsComponent";
import { FriendsProvider } from "@/contexts/FriendsContext";

export default function Home() {
  return (
    <AuthGuard redirectTo="/login">
      <FriendsProvider>
        <FriendsComponent />
      </FriendsProvider>
    </AuthGuard>
  );
}
