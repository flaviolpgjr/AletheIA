import { useState } from "react";
import type { SyntheticEvent } from "react";

import { useAnalyzePromise } from "../hooks/useAnalyzePromise";

export function AnalyzeForm() {
  const [text, setText] = useState("");

  const {
    analyze,
    loading,
    error,
    result,
  } = useAnalyzePromise();

  async function handleSubmit(
    event: SyntheticEvent<HTMLFormElement>
  ) {
    event.preventDefault();

    await analyze({ text });
  }

  return (
    <form onSubmit={handleSubmit}>
      <textarea
        value={text}
        onChange={(event) => setText(event.target.value)}
        placeholder="Cole a promessa aqui..."
      />

      <button type="submit" disabled={loading}>
        {loading ? "Analisando..." : "Analisar"}
      </button>

      {error && <p>{error}</p>}

      {result && (
        <div>
          <h2>Resumo</h2>

          <p>{result.summary}</p>

          <h3>Score</h3>

          <p>{result.score}</p>
        </div>
      )}
    </form>
  );
}