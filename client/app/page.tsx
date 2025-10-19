import { AuthGuard } from "@/components/auth/AuthGuard";
import { FyreArea } from "@/components/fyre/FyreArea";
import { MainContainer } from "@/components/MainContainer";

export default function Home() {
  return (
    <AuthGuard redirectTo="/login">
      <MainContainer>
        <FyreArea />
      </MainContainer>
    </AuthGuard>
  );
}
