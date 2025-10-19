import { AuthGuard } from "@/components/auth/AuthGuard";
import { MainContainer } from "@/components/MainContainer";

export default function Home() {
  return (
    <AuthGuard redirectTo="/login">
      <MainContainer>
        Home
      </MainContainer>
    </AuthGuard>
  );
}
