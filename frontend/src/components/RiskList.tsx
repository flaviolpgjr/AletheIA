type RiskListProps = {
  risks: string[];
};

export function RiskList({ risks }: RiskListProps) {
  return (
    <div className="rounded-2xl border border-zinc-200 bg-white p-5 transition-colors dark:border-zinc-800 dark:bg-zinc-900">
      <h2 className="text-lg font-semibold text-zinc-900 dark:text-white">
        Riscos Identificados
      </h2>

      <ul className="mt-3 space-y-2">
        {risks.map((risk) => (
          <li
            key={risk}
            className="rounded-lg border border-zinc-200 bg-zinc-50 px-3 py-2 text-zinc-700 dark:border-zinc-800 dark:bg-zinc-950 dark:text-zinc-300"
          >
            {risk}
          </li>
        ))}
      </ul>
    </div>
  );
}
