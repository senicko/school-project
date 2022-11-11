import { useUserStore } from "../../state/user";

export const SavedJokes = () => {
  const user = useUserStore((state) => state.user);

  return (
    <div className="flex flex-col gap-4">
      <h2 className="text-lg font-medium">Saved jokes</h2>
      <ul>
        {user &&
          user.jokes.map((joke) => (
            <li className="mb-2" key={joke.content}>
              <div>
                <p>{joke.content}</p>
                <span className="text-sm text-gray-600">
                  {Intl.DateTimeFormat("en-US", { dateStyle: "full" }).format(
                    new Date(joke.savedAt)
                  )}
                </span>
              </div>
            </li>
          ))}
      </ul>
    </div>
  );
};
