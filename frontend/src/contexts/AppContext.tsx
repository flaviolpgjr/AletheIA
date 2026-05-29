import { createContext } from "react";

export type Theme = "dark" | "light";

export type AppContextValue = {
  theme: Theme;
  toggleTheme: () => void;
};

export const AppContext = createContext<AppContextValue | null>(null);
