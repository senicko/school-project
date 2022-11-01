export type Definition = {
  word: string;
  meaning: string;
};

export type LearningSet = {
  id: number;
  title: string;
  words: Definition[];
};

export type LearningSetCreateData = Pick<LearningSet, "title" | "words">;
