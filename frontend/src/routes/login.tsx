import { useForm } from "react-hook-form";
import { Link, useNavigate } from "react-router-dom";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useUserStore } from "../state/user";

// Create validation schema for register credentials.
const loginCredentialsSchema = z.object({
  email: z
    .string({ required_error: "Email is required." })
    .email({ message: "Email must be a correct email address." }),
  password: z
    .string({ required_error: "Password is required." })
    .min(1, "Password is required"),
});

// Infer type of login credentials
type LoginCredentials = z.infer<typeof loginCredentialsSchema>;

/**
 * Makes a request to the api and logs in the user.
 * @param data user credentials
 * @returns the user
 */
// TODO: Think what can go wrong and handle the errors.
const loginUser = async (data: LoginCredentials) => {
  const res = await fetch("http://localhost:3000/users/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify(data),
  });

  return res.ok ? await res.json() : undefined;
};

export const Login = () => {
  const navigate = useNavigate();

  const setUser = useUserStore((state) => state.setUser);
  const {
    handleSubmit,
    register,
    formState: { errors },
  } = useForm<LoginCredentials>({
    resolver: zodResolver(loginCredentialsSchema),
  });

  /**
   * Login user handler.
   */
  const onSubmit = handleSubmit(async (data) => {
    const user = await loginUser(data);
    if (!user) return;

    setUser(user);
    navigate("/");
  });

  return (
    <main className="flex h-screen flex-col items-center justify-center gap-16 p-8">
      <h1 className="text-2xl text-gray-800">Welcome back.</h1>
      <form onSubmit={onSubmit} className="flex flex-col gap-12" noValidate>
        <div className="flex flex-col gap-4">
          <div className="flex flex-col gap-1">
            <label className="text-gray-600" htmlFor="email">
              Email
            </label>
            <input
              className="w-[400px] rounded-lg border border-gray-200 p-3 text-gray-700 placeholder:text-gray-300 focus:outline-gray-300"
              type="email"
              id="email"
              placeholder="Enter your email address"
              {...register("email")}
            />
            {errors.email && (
              <span className="rounded-lg border border-red-100 bg-red-50 px-2 py-1 text-sm text-red-400">
                {errors.email.message}
              </span>
            )}
          </div>
          <div className="flex flex-col gap-2">
            <label className="text-gray-600" htmlFor="password">
              Password
            </label>
            <input
              className="w-[400px] rounded-lg border border-gray-200 p-3 text-gray-700 placeholder:text-gray-300 focus:outline-gray-300"
              type="password"
              id="password"
              placeholder="Enter your password"
              {...register("password")}
            />
            {errors.password && (
              <span className="rounded-lg border border-red-100 bg-red-50 px-2 py-1 text-sm text-red-400">
                {errors.password.message}
              </span>
            )}
          </div>
          <div className="flex items-center justify-between">
            <div className="flex gap-2 py-2">
              <input id="remember" type="checkbox" />
              <label className="text-sm text-gray-600" htmlFor="remember">
                Remember me
              </label>
            </div>
            <Link className="text-sm text-blue-500" to="/">
              Forgot Password
            </Link>
          </div>
        </div>
        <button
          type="submit"
          className="rounded-lg bg-blue-500 p-3 text-white transition hover:bg-blue-600 focus:outline-blue-600 disabled:bg-blue-200"
        >
          Log in
        </button>
      </form>
      <footer>
        <p className="text-sm text-gray-600">
          Don't have an account yet?{" "}
          <Link className="text-blue-500" to="/register">
            Register here.
          </Link>
        </p>
      </footer>
    </main>
  );
};
