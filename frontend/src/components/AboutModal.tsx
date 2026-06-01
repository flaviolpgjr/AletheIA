import { Modal } from "./Modal";

type AboutModalProps = {
  onClose: () => void;
};

export function AboutModal({ onClose }: AboutModalProps) {
  return (
    <Modal title="Sobre o AletheIA" onClose={onClose}>
      <div className="space-y-4">
        <p>
          O <strong>AletheIA</strong> é um projeto que utiliza Inteligência
          Artificial e dados públicos para analisar a viabilidade de promessas
          públicas de forma transparente e explicável.
        </p>

        <p>
          O objetivo não é prever o futuro ou determinar verdades absolutas, mas
          apresentar análises fundamentadas em dados, indicadores e informações
          verificáveis.
        </p>

        <p>
          Os resultados são acompanhados por um nível de confiança, fontes
          utilizadas e limitações identificadas durante a análise.
        </p>

        <div>
          <h3 className="mb-2 font-semibold">Princípios do projeto</h3>

          <ul className="list-disc space-y-1 pl-5">
            <li>Não inventar dados.</li>
            <li>Mostrar as fontes utilizadas.</li>
            <li>Explicar as conclusões apresentadas.</li>
            <li>Indicar limitações e dados ausentes.</li>
            <li>Ser transparente sobre o nível de confiança da análise.</li>
          </ul>
        </div>
      </div>
    </Modal>
  );
}
