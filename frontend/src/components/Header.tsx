import { BookOpen, Info, Moon, Sun } from "lucide-react";
import { useNavigate } from "react-router-dom";

import logoIcon from "../assets/logo-icon.png";
import { useAppContext } from "../hooks/useAppContext";

export function Header() {
  const { theme, toggleTheme } = useAppContext();
  const navigate = useNavigate();

  return (
    <header className="sticky top-0 z-50 border-b border-zinc-200 bg-zinc-50/80 backdrop-blur transition-colors dark:border-zinc-800 dark:bg-zinc-950/80">
      <div className="mx-auto flex max-w-7xl items-center justify-between px-6 py-5">
        <div className="flex items-center gap-2">
          <img src={logoIcon} alt="AletheIA" className="h-12 w-12" />

          <span className="text-xl font-semibold tracking-tight text-zinc-950 dark:text-white">
            Alethe
            <span className="bg-gradient-to-r from-cyan-400 to-violet-500 bg-clip-text text-transparent">
              IA
            </span>
          </span>
        </div>

        <div className="flex items-center gap-3">
          <button
            type="button"
            onClick={() => navigate("/sobre")}
            className="rounded-xl border border-zinc-300 p-3 text-zinc-700 transition hover:bg-zinc-100 dark:border-zinc-700 dark:text-zinc-300 dark:hover:bg-zinc-800"
            aria-label="Sobre o projeto"
            title="Sobre o projeto"
          >
            <Info size={20} />
          </button>

          <button
            type="button"
            onClick={() => navigate("/metodologia")}
            className="rounded-xl border border-zinc-300 p-3 text-zinc-700 transition hover:bg-zinc-100 dark:border-zinc-700 dark:text-zinc-300 dark:hover:bg-zinc-800"
            aria-label="Como calculamos o score"
            title="Como calculamos o score"
          >
            <BookOpen size={20} />
          </button>

          <button
            type="button"
            onClick={toggleTheme}
            className="rounded-xl border border-zinc-300 p-3 text-zinc-700 transition hover:bg-zinc-100 dark:border-zinc-700 dark:text-zinc-300 dark:hover:bg-zinc-800"
            aria-label="Alternar tema"
            title="Alternar tema"
          >
            {theme === "dark" ? <Sun size={20} /> : <Moon size={20} />}
          </button>
        </div>
      </div>
    </header>
  );
}
