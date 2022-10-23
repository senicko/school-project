import { useAuth } from "../context/auth";
import { ReactNode } from "react";
import { Navigate } from "react-router-dom";

/**
 * Unauthorized renders it's children only if current user is not logged in.
 * Otherwise it redirects to the home page.
 */
export const Unauthorized = ({ children }: { children: ReactNode }) => {
  const { user } = useAuth();
  return user ? <Navigate to="/" /> : <>{children}</>;
};
