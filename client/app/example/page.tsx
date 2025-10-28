"use client";
import { AuthGuard } from "@/components/auth/AuthGuard";
import FriendsList from "@/components/examples/friendlist/FriendsList";
import { FyreList } from "@/components/examples/fyrelist";
import { useAuth } from "@/contexts/AuthContext";

export default function Examples() {
  const { user, isAuthenticated } = useAuth();

  if (user && isAuthenticated) {
    return (
      <AuthGuard redirectTo="/login">
        <div className="flex">
          <FriendsList />
          <FyreList user={user} />
        </div>
      </AuthGuard>
    );
  }
}
