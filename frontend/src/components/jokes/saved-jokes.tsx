import { Joke, useUserStore } from "../../state/user";
import { useMemo, useState } from "react";

/**
 * Checks if joke includes string
 * @param joke jokes
 * @param filter strings which must be present in a joke
 * @returns true, meaning joke includes the string, or false, meaning it does not.
 */
const filterJoke = (joke: Joke, filter: string) =>
  joke.content.toLowerCase().includes(filter.toLowerCase());

export const SavedJokes = () => {
  const user = useUserStore((state) => state.user);
  const [filter, setFilter] = useState("");

  return (
    <section className="flex flex-col gap-4">
      <h2 className="text-lg font-medium">Saved jokes</h2>
      <div>
        <label htmlFor="joke-filter">Filter jokes: </label>
        <input
          id="joke-filter"
          type="text"
          className="border border-gray-500"
          onChange={(e) => setFilter(e.target.value)}
        />
      </div>
      <ul>
        {user &&
          user.jokes
            .filter((joke) => filterJoke(joke, filter))
            .map((joke) => (
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
    </section>
  );
};
