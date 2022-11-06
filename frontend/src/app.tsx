import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { ErrorPage } from "./error-page";
import { Register } from "./routes/register";
import { Login } from "./routes/login";
import { Root } from "./routes/root";
import { AuthProvider } from "./context/auth";
import { ProtectedRoute } from "./components/protected";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Unauthorized } from "./components/unauthorized";

const router = createBrowserRouter([
  {
    path: "/",
    element: (
      <ProtectedRoute>
        <Root />
      </ProtectedRoute>
    ),
    errorElement: <ErrorPage />,
  },
  {
    path: "/register",
    element: (
      <Unauthorized>
        <Register />
      </Unauthorized>
    ),
  },
  {
    path: "/login",
    element: (
      <Unauthorized>
        <Login />
      </Unauthorized>
    ),
  },
]);

const queryClient = new QueryClient();

export const App = () => {
  return (
    <QueryClientProvider client={queryClient}>
      <AuthProvider>
        <RouterProvider router={router} />
      </AuthProvider>
    </QueryClientProvider>
  );
};
