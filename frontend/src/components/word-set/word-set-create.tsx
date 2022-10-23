import { WordSetEntry } from "../../types";
import { useState } from "react";
import { useForm } from "react-hook-form";

type WordInputProps = {
  onSubmit: (wordSetEntry: WordSetEntry) => void;
};

export const WordInput = ({ onSubmit }: WordInputProps) => {
  const { handleSubmit, register } = useForm<WordSetEntry>();

  return (
    <div className="flex gap-2">
      <form onSubmit={handleSubmit((data) => onSubmit(data))}>
        <input
          className="border border-gray-200 rounded-lg p-3 w-[400px] text-gray-700 focus:outline-gray-300 placeholder:text-gray-300"
          type="text"
          autoComplete="off"
          {...register("word")}
        />
        <input
          className="border border-gray-200 rounded-lg p-3 w-[400px] text-gray-700 focus:outline-gray-300 placeholder:text-gray-300"
          type="text"
          autoComplete="off"
          {...register("meaning")}
        />
        <button>Add</button>
      </form>
    </div>
  );
};

export const WordSetCreate = () => {
  const [wordSetEntries, setWordSetEntries] = useState<WordSetEntry[]>([]);

  return (
    <section>
      {wordSetEntries.map(({ word, meaning }) => (
        <div key={word}>
          {word} = {meaning}
        </div>
      ))}
      <WordInput
        onSubmit={(wordSetEntry) =>
          setWordSetEntries([...wordSetEntries, wordSetEntry])
        }
      />
    </section>
  );
};
