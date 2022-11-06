import { ReactNode } from "react";
import { Navigate } from "react-router-dom";
import { useUserStore } from "../state/user";

/**
 * Unauthorized renders it's children only if current user is not logged in.
 * Otherwise it redirects to the home page.
 */
export const Unauthorized = ({ children }: { children: ReactNode }) => {
  const user = useUserStore((state) => state.user);
  return user ? <Navigate to="/" /> : <>{children}</>;
};
