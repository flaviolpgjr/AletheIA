import { Modal } from "./Modal";

type MethodologyModalProps = {
  onClose: () => void;
};

export function MethodologyModal({ onClose }: MethodologyModalProps) {
  return (
    <Modal title="Como calculamos o score" onClose={onClose}>
      <div className="max-h-[70vh] space-y-6 overflow-y-auto pr-2">
        <p>
          O score do AletheIA varia de <strong>0 a 100</strong> e representa uma
          avaliação da qualidade, verificabilidade e plausibilidade de uma
          promessa pública com base em critérios objetivos e evidências
          públicas.
        </p>

        <p>
          A metodologia é inspirada em critérios internacionais de avaliação de
          políticas públicas e dados públicos verificáveis.
        </p>

        <section>
          <h3 className="mb-3 font-semibold">Critérios avaliados</h3>

          <div className="overflow-hidden rounded-xl border border-zinc-200 dark:border-zinc-800">
            <table className="w-full text-left text-sm">
              <thead className="bg-zinc-100 dark:bg-zinc-800">
                <tr>
                  <th className="px-4 py-3">Critério</th>
                  <th className="px-4 py-3 text-right">Peso</th>
                </tr>
              </thead>

              <tbody>
                <MethodologyRow
                  name="Clareza"
                  weight={10}
                  description="A promessa é objetiva e compreensível?"
                />

                <MethodologyRow
                  name="Mensurabilidade"
                  weight={15}
                  description="Existe um indicador que permita medir o resultado?"
                />

                <MethodologyRow
                  name="Prazo definido"
                  weight={10}
                  description="Há um prazo ou período claramente definido?"
                />

                <MethodologyRow
                  name="Dados públicos disponíveis"
                  weight={15}
                  description="Existem fontes públicas para acompanhar a promessa?"
                />

                <MethodologyRow
                  name="Histórico comparável"
                  weight={15}
                  description="Há dados históricos que permitam comparação?"
                />

                <MethodologyRow
                  name="Plausibilidade baseada em evidências"
                  weight={20}
                  description="A meta parece compatível com os dados públicos identificados?"
                />

                <MethodologyRow
                  name="Riscos e dependências"
                  weight={15}
                  description="A promessa depende de fatores externos relevantes?"
                />
              </tbody>
            </table>
          </div>
        </section>

        <section>
          <h3 className="mb-2 font-semibold">Como funciona a plausibilidade</h3>

          <p>
            Quando uma promessa possui uma meta quantitativa, o AletheIA tenta
            identificar uma linha de base pública para comparação.
          </p>

          <p className="mt-2">Exemplo:</p>

          <div className="mt-3 rounded-xl border border-zinc-200 bg-zinc-50 p-4 dark:border-zinc-800 dark:bg-zinc-900">
            <p>
              Meta proposta: <strong>100 hospitais</strong>
            </p>

            <p>
              Linha de base identificada: <strong>5.115 hospitais</strong>
            </p>

            <p>
              Relação: <strong>1,96%</strong>
            </p>
          </div>

          <p className="mt-3">
            Essa comparação ajuda a avaliar se a meta proposta parece compatível
            com a realidade observada nos dados públicos.
          </p>
        </section>

        <section>
          <h3 className="mb-2 font-semibold">Limitações importantes</h3>

          <ul className="list-disc space-y-1 pl-5">
            <li>O score não representa uma previsão de sucesso.</li>
            <li>O score não representa uma verdade absoluta.</li>
            <li>
              A qualidade da análise depende da disponibilidade de dados
              públicos.
            </li>
            <li>
              Algumas promessas exigem interpretação humana e contexto político.
            </li>
          </ul>
        </section>
      </div>
    </Modal>
  );
}

type MethodologyRowProps = {
  name: string;
  weight: number;
  description: string;
};

function MethodologyRow({ name, weight, description }: MethodologyRowProps) {
  return (
    <tr className="border-t border-zinc-200 dark:border-zinc-800">
      <td className="px-4 py-3">
        <div>
          <p className="font-medium">{name}</p>

          <p className="mt-1 text-xs text-zinc-500 dark:text-zinc-400">
            {description}
          </p>
        </div>
      </td>

      <td className="px-4 py-3 text-right font-semibold">{weight}</td>
    </tr>
  );
}
