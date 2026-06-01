import { useLocation, useNavigate } from "react-router-dom";

import { AboutModal } from "../components/AboutModal";
import { AnalyzeForm } from "../components/AnalyzeForm";
import { Header } from "../components/Header";
import { HeroSection } from "../components/HeroSection";

export function HomePage() {
  const location = useLocation();
  const navigate = useNavigate();

  const isAboutOpen = location.pathname === "/sobre";

  return (
    <main className="min-h-screen bg-zinc-50 text-zinc-950 transition-colors dark:bg-zinc-950 dark:text-white">
      <Header />

      <div className="mx-auto flex max-w-4xl flex-col items-center px-6 py-14">
        <HeroSection />

        <div className="w-full">
          <AnalyzeForm />
        </div>
      </div>

      {isAboutOpen && <AboutModal onClose={() => navigate("/")} />}
    </main>
  );
}
