import { WordSet } from "../../types";

export type WordSetProps = {
  wordSet: WordSet;
};

export const WordSetPreview = ({ wordSet }: WordSetProps) => {
  return (
    <div>
      <p>{wordSet.title}</p>
    </div>
  );
};
