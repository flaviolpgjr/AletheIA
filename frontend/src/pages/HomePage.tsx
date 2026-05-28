import { AnalyzeForm } from "../components/AnalyzeForm";

export function HomePage() {
  return (
    <main>
      <h1>AletheIA</h1>

      <p>
        Analise promessas e identifique riscos automaticamente.
      </p>

      <AnalyzeForm />
    </main>
  );
}