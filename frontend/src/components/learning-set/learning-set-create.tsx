import { Definition, LearningSetCreateData } from "../../types";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";

type WordInputProps = {
  onSubmit: (definition: Definition) => void;
};

/**
 * createLearningSet sends a new word set to the api.
 */
const createLearningSet = async (wordSet: LearningSetCreateData) => {
  await fetch("http://localhost:3000/word-set", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify(wordSet),
  });
};

const WordInput = ({ onSubmit }: WordInputProps) => {
  const { handleSubmit, register } = useForm<Definition>();

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

export const LearningSetCreate = () => {
  const navigate = useNavigate();
  const [definitions, setDefinitions] = useState<Definition[]>([]);

  /**
   * addLearningSet creates a new word set and navigates the user to the home page.
   */
  const addLearningSet = async () => {
    await createLearningSet({
      title: "Test Word Set",
      words: definitions,
    });

    navigate("/");
  };

  return (
    <section>
      <button
        className="p-4 bg-blue-500 rounded-lg text-white"
        onClick={addLearningSet}
      >
        Create
      </button>
      <WordInput
        onSubmit={(definition) => setDefinitions([...definitions, definition])}
      />
      {definitions.map(({ word, meaning }) => (
        <div key={word}>
          {word} = {meaning}
        </div>
      ))}
    </section>
  );
};
