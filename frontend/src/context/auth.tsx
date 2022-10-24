import {
  createContext,
  ReactNode,
  useContext,
  useEffect,
  useMemo,
  useState,
} from "react";

export type User = {
  name: string;
  email: string;
};

export type TAuthContext = {
  user?: User;
  setUser: React.Dispatch<React.SetStateAction<User | undefined>>;
};

export const AuthContext = createContext<TAuthContext | undefined>(undefined);

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
  // TODO: Does it make sense to use react query here?
  const [user, setUser] = useState<User>();
  const [isLoading, setIsLoading] = useState(true);
  const value = useMemo(() => ({ user, setUser }), [user]);

  useEffect(() => {
    getCurrentUser()
      .then((user) => setUser(user))
      .finally(() => setIsLoading(false));
  }, []);

  if (isLoading) return <></>;

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
};

/**
 * useAuth hook gives easy access to AuthContext.
 */
export const useAuth = () => useContext(AuthContext)!;
