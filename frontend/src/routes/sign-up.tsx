import { useContext } from "react";
import { useNavigate } from "react-router-dom";
import { AuthContext } from "../context/auth";
import { useForm } from "react-hook-form";

type FormData = {
  name: string;
  email: string;
  password: string;
};

export const SignUp = () => {
  const navigate = useNavigate();
  const { setUser } = useContext(AuthContext);
  const { register, handleSubmit } = useForm<FormData>();

  const signUp = handleSubmit(async (data) => {
    const res = await fetch("http://localhost:3000/users", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify(data),
    });

    const user = await res.json();
    setUser(user);
    navigate("/");
  });

  return (
    <main className="flex flex-col h-screen items-center gap-8 p-8">
      <h1 className="text-lg text-gray-800">Sign Up</h1>
      <form onSubmit={signUp} className="flex flex-col gap-4">
        <div className="flex flex-col gap-1">
          <label className="text-gray-600" htmlFor="name">
            Name
          </label>
          <input
            className="border border-gray-300 rounded-md p-2 w-[350px] text-gray-800 focus:outline-blue-500"
            type="text"
            id="name"
            placeholder="Enter your name"
            autoComplete="off"
            {...register("name")}
          />
        </div>

        <div className="flex flex-col gap-1">
          <label className="text-gray-600" htmlFor="email">
            Email
          </label>
          <input
            className="border border-gray-300 rounded-md p-2 w-[350px] text-gray-800 focus:outline-blue-500"
            type="email"
            id="email"
            placeholder="Enter your email address"
            {...register("email")}
          />
        </div>

        <div className="flex flex-col gap-1">
          <label className="text-gray-600" htmlFor="password">
            Password
          </label>
          <input
            className="border border-gray-300 rounded-md p-2 w-[350px] text-gray-800 focus:outline-blue-500"
            type="password"
            id="password"
            placeholder="Choose a password"
            {...register("password")}
          />
        </div>
        <button
          type="submit"
          className="w-[350px] mt-4 bg-blue-600 hover:bg-blue-700 transition p-2 text-white rounded-lg focus:outline-blue-900"
        >
          Login
        </button>
      </form>
    </main>
  );
};
