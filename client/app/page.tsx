import { AuthGuard } from "@/components/auth/AuthGuard";
import { FyreArea } from "@/components/fyre/FyreArea";

export default function Home() {
  return (
    <AuthGuard redirectTo="/landing">
      <FyreArea />
    </AuthGuard>
  );
}
