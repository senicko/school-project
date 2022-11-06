import create from "zustand";
import { persist, devtools } from "zustand/middleware";

export type User = {
  name: string;
  email: string;
  jokes: string[];
};

export type UserState = {
  user?: User;
  setUser: (user: User) => void;
  addJoke: (joke: string) => void;
};

export const useUserStore = create<UserState>()(
  devtools(
    persist(
      (set) => ({
        user: undefined,
        setUser: (user: User) => set(() => ({ user })),
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
