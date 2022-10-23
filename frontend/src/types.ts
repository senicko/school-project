export type WordSetEntry = {
  word: string;
  meaning: string;
};

export type WordSet = {
  id: number;
  title: string;
  words: WordSetEntry[];
};
