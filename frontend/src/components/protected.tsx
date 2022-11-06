import { ReactNode, useEffect } from "react";
import { Navigate } from "react-router-dom";
import { useUserStore } from "../state/user";

/**
 * ProtectedRoute renders it's children only if user is logged in.
 * Otherwise it navigates to register page.
 */
export const ProtectedRoute = ({ children }: { children: ReactNode }) => {
  const user = useUserStore((state) => state.user);
  return user ? <>{children}</> : <Navigate to="/register" />;
};
