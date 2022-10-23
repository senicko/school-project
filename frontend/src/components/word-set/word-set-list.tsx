import { useQuery } from "@tanstack/react-query";
import { json } from "react-router-dom";
import { WordSet } from "../../types";
import { WordSetPreview } from "./word-set-preview";

/**
 * getWordSets retrieves all word sets created by the current user.
 * @returns array of word sets
 */
const getWordSets = async (): Promise<WordSet[]> => {
  const res = await fetch("http://localhost:3000/word-set", {
    credentials: "include",
  });

  return await res.json();
};

/**
 * WordSetList is a component that renders all word sets created by the currently logged in user.
 */
export const WordSetList = () => {
  // Use react query to manage the wordSets query
  const wordSetsQuery = useQuery(["wordSets"], getWordSets);

  // If word sets are loading
  if (wordSetsQuery.isLoading) return <section>Loading ...</section>;

  // If query failed
  if (wordSetsQuery.isError) return <section>An error has occured.</section>;

  return (
    <section>
      {wordSetsQuery.data &&
        wordSetsQuery.data.map((wordSet) => (
          <WordSetPreview wordSet={wordSet} />
        ))}
    </section>
  );
};
