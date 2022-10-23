import { useNavigate, Link } from "react-router-dom";
import { useAuth } from "../context/auth";
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
  const { setUser } = useAuth();
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterCredentials>({
    resolver: zodResolver(registerCredentialsSchema),
  });

  const onSubmit = handleSubmit(async (data) => {
    const user = await registerUser(data);

    if (user) {
      setUser(user);
      navigate("/");
      return;
    }

    setSubmitError("Could not register right now. Try again later.");
  });

  return (
    <main className="flex flex-col h-screen justify-center items-center gap-16 p-8">
      <h1 className="text-2xl text-gray-800">Create an account.</h1>
      <form onSubmit={onSubmit} className="flex flex-col gap-12" noValidate>
        <div className="flex flex-col gap-4">
          <div className="flex flex-col gap-1">
            <label className="text-gray-600" htmlFor="name">
              Name
            </label>
            <input
              className="border border-gray-200 rounded-lg p-3 w-[400px] text-gray-700 focus:outline-gray-300 placeholder:text-gray-300"
              type="text"
              id="name"
              placeholder="Enter your name"
              {...register("name")}
            />
            {errors.name && (
              <span className="text-red-400 text-sm px-2 py-1 bg-red-50 border border-red-100 rounded-lg">
                {errors.name.message}
              </span>
            )}
          </div>
          <div className="flex flex-col gap-1">
            <label className="text-gray-600" htmlFor="email">
              Email
            </label>
            <input
              className="border border-gray-200 rounded-lg p-3 w-[400px] text-gray-700 focus:outline-gray-300 placeholder:text-gray-300"
              type="email"
              id="email"
              placeholder="Enter your email address"
              {...register("email")}
            />
            {errors.email && (
              <span className="text-red-400 text-sm px-2 py-1 bg-red-50 border border-red-100 rounded-lg">
                {errors.email.message}
              </span>
            )}
          </div>
          <div className="flex flex-col gap-1">
            <label className="text-gray-600" htmlFor="password">
              Password
            </label>
            <input
              className="border border-gray-200 rounded-lg p-3 w-[400px] text-gray-700 focus:outline-gray-300 placeholder:text-gray-300"
              type="password"
              id="password"
              placeholder="Choose a password"
              {...register("password")}
            />
            {errors.password && (
              <span className="text-red-400 text-sm px-2 py-1 bg-red-50 border border-red-100 rounded-lg">
                {errors.password.message}
              </span>
            )}
          </div>
          <div className="flex justify-between items-center">
            <div className="flex gap-2 py-2">
              <input id="remember" type="checkbox" />
              <label className="text-gray-600 text-sm" htmlFor="remember">
                Remember me
              </label>
            </div>
          </div>
        </div>
        <div className="flex flex-col gap-2">
          <button
            type="submit"
            className="bg-blue-500 hover:bg-blue-600 transition p-3 text-white rounded-lg focus:outline-blue-600"
          >
            Continue
          </button>
          {submitError && (
            <span className="text-red-400 text-sm px-2 py-1 bg-red-50 border border-red-100 rounded-lg">
              {submitError}
            </span>
          )}
        </div>
      </form>
      <p className="text-sm text-gray-600">
        Already have an account?{" "}
        <Link className="text-blue-500" to="/login">
          Login here.
        </Link>
      </p>
    </main>
  );
};
