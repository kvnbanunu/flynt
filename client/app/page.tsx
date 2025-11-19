import { AuthGuard } from "@/components/auth/AuthGuard";
import { FyreTotal } from "@/components/flame/FyreTotal";
import { FyreArea } from "@/components/fyre/FyreArea";

export default function Home() {
  return (
    <AuthGuard redirectTo="/landing">
      <div className="flex flex-col w-full gap-4">
        <FyreTotal />
        <FyreArea />
      </div>
    </AuthGuard>
  );
}
