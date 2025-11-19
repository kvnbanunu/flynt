import { AuthGuard } from "@/components/auth/AuthGuard";
import { Profile } from "@/components/profile/Profile";

export default function ProfilePage() {
  return (
    <AuthGuard redirectTo="/landing">
      <Profile />
    </AuthGuard>
  );
}
