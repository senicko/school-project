import {
  createContext,
  ReactNode,
  useContext,
  useEffect,
  useMemo,
  useState,
} from "react";

export type TAuthContext = {
  user: unknown;
  setUser: React.Dispatch<React.SetStateAction<undefined>>;
};

export const AuthContext = createContext<TAuthContext>({
  user: undefined,
  setUser: () => {},
});

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [user, setUser] = useState<undefined>();
  const [isLoading, setIsLoading] = useState(true);

  const value = useMemo(
    () => ({
      user,
      setUser,
    }),
    [user]
  );

  useEffect(() => {
    const getUser = async () => {
      const res = await fetch("http://localhost:3000/users/me", {
        credentials: "include",
      });

      if (res.ok) {
        const user = await res.json();
        setUser(user);
      }

      setIsLoading(false);
    };

    getUser();
  }, []);

  return isLoading ? (
    <></>
  ) : (
    <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
  );
};

export const useAuth = () => useContext(AuthContext);
