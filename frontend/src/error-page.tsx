import { useRouteError } from "react-router-dom";

export const ErrorPage = () => {
  const error = useRouteError() as Response;
  console.error(error);

  return (
    <main className="flex h-screen w-full flex-col items-center justify-center gap-4">
      <h1 className="font-bold">Oops!</h1>
      <p>Sorry, an unexpected error has occured.</p>
      <p>
        <i className="text-gray-500">{error.statusText}</i>
      </p>
    </main>
  );
};
