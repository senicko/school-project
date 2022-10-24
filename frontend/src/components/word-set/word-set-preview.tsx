import { WordSet } from "../../types";

export type WordSetProps = {
  wordSet: WordSet;
};

export const WordSetPreview = ({ wordSet }: WordSetProps) => {
  return (
    <div className="flex flex-col gap-4 p-4">
      <p>{wordSet.title}</p>
      <table className="table-fixed border-collapse border-gray-200">
        <thead>
          <tr>
            <th className="border border-gray-200 bg-gray-100">Word</th>
            <th className="border border-gray-200 bg-gray-100">Meaning</th>
          </tr>
        </thead>
        <tbody>
          {wordSet.words.map(({ word, meaning }) => (
            <tr>
              <td className="p-4 border border-gray-200">
                <span className="text-gray-700">{word}</span>{" "}
              </td>
              <td className="p-4 border border-gray-200">
                <span className="text-black">{meaning}</span>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};
