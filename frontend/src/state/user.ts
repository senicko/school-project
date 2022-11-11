import create from "zustand";
import { persist, devtools } from "zustand/middleware";

export type Joke = {
  content: string;
  savedAt: number;
};

export type User = {
  name: string;
  email: string;
  jokes: Joke[];
};

export type UserState = {
  user?: User;
  setUser: (user: User | undefined) => void;
  addJoke: (joke: Joke) => void;
};

export const useUserStore = create<UserState>()(
  devtools(
    persist(
      (set) => ({
        user: undefined,
        setUser: (user: User | undefined) => set(() => ({ user })),
        addJoke: (joke) => {
          set((state) =>
            state.user
              ? { user: { ...state.user, jokes: [...state.user.jokes, joke] } }
              : state
          );
        },
      }),
      {
        name: "user-store",
      }
    )
  )
);
