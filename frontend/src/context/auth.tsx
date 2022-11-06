import { ReactNode, useEffect, useState } from "react";
import { useUserStore } from "../state/user";

export type User = {
  name: string;
  email: string;
  jokes: string[];
};

/**
 * Fetches currently logged in user.
 * @returns current user
 */
// TODO: Think what can go wrong, and handle the errors.
const getCurrentUser = async () => {
  const res = await fetch("http://localhost:3000/users/me", {
    credentials: "include",
  });

  return res.ok ? await res.json() : undefined;
};

/**
 * AuthProvider provides auth context to it's children components.
 * @param props
 */
export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const setUser = useUserStore((state) => state.setUser);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    getCurrentUser()
      .then((user) => setUser(user))
      .finally(() => setIsLoading(false));
  }, []);

  return isLoading ? <></> : <>{children}</>;
};
