import { useEffect, useState } from "react";
import { BookmarkIcon as BookmarkIconOutline } from "@heroicons/react/24/outline";
import { BookmarkIcon as BookmarkIconSolid } from "@heroicons/react/24/solid";
import { useUserStore } from "../../state/user";

/**
 * Represents object returned from api.yomomma.info api.
 */
type Joke = {
  joke: string;
};

/**
 * fetchJoke fetches a joke using an api procy to 'api.yomomma.info' api.
 * @returns fetched joke about yomomma.
 */
const fetchJoke = async (): Promise<Joke> => {
  const res = await fetch("http://localhost:3000/joke");
  return res.json();
};

/**
 * Saves the joke in user's joke collection.
 * @param joke joke that will be saved.
 * @returns timestamp when the joke was saved.
 */
const saveJoke = async (joke: string): Promise<number> => {
  const savedAt = Date.now();

  await fetch("http://localhost:3000/users/me/jokes", {
    method: "POST",
    credentials: "include",
    body: JSON.stringify({
      content: joke,
      savedAt,
    }),
  });

  return savedAt;
};

/**
 * JokesGenerator component fetches a joke every 10 seconds. It allows to save the joke if user is logged in.
 */
export const JokesGenerator = () => {
  // Get current user
  const user = useUserStore((state) => state.user);
  // Get function for adding saved jokes to his list
  const addJoke = useUserStore((state) => state.addJoke);

  // Joke generator state
  const [joke, setJoke] = useState("");
  const [timer, setTimer] = useState(10);

  // jokeInterval
  useEffect(() => {
    fetchJoke().then(({ joke }) => {
      setJoke(joke);
      setTimer(10);
    });

    let timerInterval = setInterval(
      () => setTimer((timer) => (timer > 1 ? timer - 1 : timer)),
      1000
    );

    const jokeInterval = setInterval(() => {
      fetchJoke().then(({ joke }) => {
        setJoke(joke);
        setTimer(10);

        clearInterval(timerInterval);
        timerInterval = setInterval(
          () => setTimer((timer) => (timer > 1 ? timer - 1 : timer)),
          1000
        );
      });
    }, 10 * 1000);

    return () => {
      clearInterval(timerInterval);
      clearInterval(jokeInterval);
    };
  }, []);

  /**
   * bookmark saves the joke in user's collection.
   */
  const bookmark = async () => {
    const savedAt = await saveJoke(joke);

    addJoke({
      content: joke,
      savedAt,
    });
  };

  return (
    <div className="flex flex-col items-center justify-center gap-8">
      <div className="flex w-full justify-end px-8">
        {user && user.jokes.some(({ content }) => content === joke) ? (
          <BookmarkIconSolid className="h-8 w-8 " />
        ) : (
          <BookmarkIconOutline
            className="h-8 w-8 cursor-pointer"
            onClick={bookmark}
          />
        )}
      </div>
      <span className="text-center text-5xl font-medium">{joke}</span>
      <span className="text-3xl italic">next joke in {timer}</span>
    </div>
  );
};
