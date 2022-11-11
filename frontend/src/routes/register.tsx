import { useNavigate, Link } from "react-router-dom";
import { useUserStore } from "../state/user";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { zodResolver } from "@hookform/resolvers/zod";
import { useState } from "react";

// Create validation schema for register credentials.
const registerCredentialsSchema = z.object({
  name: z
    .string()
    .min(1, "Name must be at least 1 character long")
    .max(64, "Name can be at most 64 characters long."),
  email: z
    .string({ required_error: "Email is required." })
    .email({ message: "Email must be a correct email address." }),
  password: z
    .string({ required_error: "Password is required." })
    .min(8, "Password must be at least 8 characters long.")
    .max(32, "password can be at most 32 characters long."),
});

// Infer type of register credentials.
type RegisterCredentials = z.infer<typeof registerCredentialsSchema>;

/**
 * Makes a request to the api and registers the user.
 * @param data user credentials
 * @returns the user
 */
// TODO: Think what can go wrong and handle the errors.
const registerUser = async (data: RegisterCredentials) => {
  const res = await fetch("http://localhost:3000/users/register", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
    body: JSON.stringify(data),
  });

  return res.ok ? await res.json() : undefined;
};

export const Register = () => {
  const navigate = useNavigate();
  const [submitError, setSubmitError] = useState<string>();
  const setUser = useUserStore((state) => state.setUser);
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterCredentials>({
    resolver: zodResolver(registerCredentialsSchema),
  });

  /**
   * Register user handler.
   */
  const onSubmit = handleSubmit(async (data) => {
    const user = await registerUser(data);

    if (!user) {
      setSubmitError("Could not register right now. Try again later.");
      return;
    }

    setUser(user);
    navigate("/");
  });

  return (
    <main className="flex h-screen flex-col items-center justify-center gap-16 p-8">
      <h1 className="text-2xl text-gray-800">Create an account.</h1>
      <form onSubmit={onSubmit} className="flex flex-col gap-12" noValidate>
        <div className="flex flex-col gap-4">
          <div className="flex flex-col gap-1">
            <label className="text-gray-600" htmlFor="name">
              Name
            </label>
            <input
              className="w-[400px] rounded-lg border border-gray-200 p-3 text-gray-700 placeholder:text-gray-300 focus:outline-gray-300"
              type="text"
              id="name"
              placeholder="Enter your name"
              autoComplete="off"
              {...register("name")}
            />
            {errors.name && (
              <span className="rounded-lg border border-red-100 bg-red-50 px-2 py-1 text-sm text-red-400">
                {errors.name.message}
              </span>
            )}
          </div>
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
          <div className="flex flex-col gap-1">
            <label className="text-gray-600" htmlFor="password">
              Password
            </label>
            <input
              className="w-[400px] rounded-lg border border-gray-200 p-3 text-gray-700 placeholder:text-gray-300 focus:outline-gray-300"
              type="password"
              id="password"
              placeholder="Choose a password"
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
          </div>
        </div>
        <div className="flex flex-col gap-2">
          <button
            type="submit"
            className="rounded-lg bg-blue-500 p-3 text-white transition hover:bg-blue-600 focus:outline-blue-600"
          >
            Continue
          </button>
          {submitError && (
            <span className="rounded-lg border border-red-100 bg-red-50 px-2 py-1 text-sm text-red-400">
              {submitError}
            </span>
          )}
        </div>
      </form>
      <footer>
        <p className="text-sm text-gray-600">
          Already have an account?{" "}
          <Link className="text-blue-500" to="/login">
            Login here.
          </Link>
        </p>
      </footer>
    </main>
  );
};
