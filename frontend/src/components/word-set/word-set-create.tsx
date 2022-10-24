import { WordSetEntry, WordSet, WordSetCreateData } from "../../types";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { useNavigate, useNavigation } from "react-router-dom";

type WordInputProps = {
  onSubmit: (wordSetEntry: WordSetEntry) => void;
};

/**
 * createWordSet sends a new word set to the api.
 */
const createWordSet = async (wordSet: WordSetCreateData) => {
  const res = await fetch("http://localhost:3000/word-set", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify(wordSet),
  });
};

const WordInput = ({ onSubmit }: WordInputProps) => {
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
  const navigate = useNavigate();
  const [wordSetEntries, setWordSetEntries] = useState<WordSetEntry[]>([]);
  const [title, setTitel] = useState("");

  /**
   * addWordSet creates a new word set and navigates the user to the home page.
   */
  const addWordSet = async () => {
    await createWordSet({
      title: "Test Word Set",
      words: wordSetEntries,
    });

    navigate("/");
  };

  return (
    <section>
      <button className="" onClick={addWordSet}>
        Create Word Set
      </button>
      <WordInput
        onSubmit={(wordSetEntry) =>
          setWordSetEntries([...wordSetEntries, wordSetEntry])
        }
      />
      {wordSetEntries.map(({ word, meaning }) => (
        <div key={word}>
          {word} = {meaning}
        </div>
      ))}
    </section>
  );
};
